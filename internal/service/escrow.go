package service

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/kanowfy/donorbox/internal/convert"
	"github.com/kanowfy/donorbox/internal/db"
	"github.com/kanowfy/donorbox/internal/dto"
	"github.com/kanowfy/donorbox/internal/mail"
	"github.com/kanowfy/donorbox/internal/model"
	"github.com/kanowfy/donorbox/internal/publish"
	"github.com/kanowfy/donorbox/internal/token"
	"github.com/kanowfy/donorbox/pkg/helper"
	"golang.org/x/crypto/bcrypt"
)

type Escrow interface {
	Login(ctx context.Context, req dto.LoginRequest) (string, error)
	GetEscrowByID(ctx context.Context, id int64) (*model.EscrowUser, error)
	ApproveOfProject(ctx context.Context, req dto.ProjectApprovalRequest) error
	ResolveMilestone(ctx context.Context, escrowID int64, req dto.ResolveMilestoneRequest) error
	ApproveUserVerification(ctx context.Context, req dto.VerificationApprovalRequest) error
	GetStatistics(ctx context.Context) (*dto.GetStatisticsResponse, error)
}

type escrow struct {
	repository db.Querier
	mailer     mail.Mailer
	publisher  publish.Publisher
}

func NewEscrow(querier db.Querier, mailer mail.Mailer, publisher publish.Publisher) Escrow {
	return &escrow{
		repository: querier,
		mailer:     mailer,
		publisher:  publisher,
	}
}

func (e *escrow) Login(ctx context.Context, req dto.LoginRequest) (string, error) {
	escrow, err := e.repository.GetEscrowUserByEmail(ctx, req.Email)
	if err != nil {
		return "", ErrUserNotFound
	}

	// validate password
	if err = bcrypt.CompareHashAndPassword([]byte(escrow.HashedPassword), []byte(req.Password)); err != nil {
		return "", ErrWrongPassword
	}

	token, err := token.GenerateToken(escrow.ID, time.Hour*3*24)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (e *escrow) GetEscrowByID(ctx context.Context, id int64) (*model.EscrowUser, error) {
	escrow, err := e.repository.GetEscrowUserByID(ctx, id)
	if err != nil {
		return nil, ErrUserNotFound
	}

	return &model.EscrowUser{
		ID:        escrow.ID,
		Email:     escrow.Email,
		UserType:  model.ESCROW,
		CreatedAt: convert.MustPgTimestampToTime(escrow.CreatedAt),
	}, nil
}

func (e *escrow) ApproveOfProject(ctx context.Context, req dto.ProjectApprovalRequest) error {
	project, err := e.repository.GetProjectByID(ctx, req.ProjectID)
	if err != nil {
		return err
	}

	user, err := e.repository.GetUserByID(ctx, project.UserID)
	if err != nil {
		return err
	}

	params := db.UpdateProjectStatusParams{
		ID: project.ID,
	}

	if req.Approved != nil {
		params.Status = db.ProjectStatusOngoing
		notif, err := e.repository.CreateNotification(ctx, db.CreateNotificationParams{
			UserID:           user.ID,
			NotificationType: db.NotificationTypeApprovedProject,
			Message:          fmt.Sprintf("Congratulations! Your project \"%s\" has been approved and ready to receive funds.", project.Title),
			ProjectID:        &project.ID,
		})
		if err != nil {
			return err
		}
		helper.Background(func() {
			event := model.Notification{
				ID:               notif.ID,
				UserID:           notif.UserID,
				NotificationType: model.NotificationType(notif.NotificationType),
				Message:          notif.Message,
				ProjectID:        notif.ProjectID,
				IsRead:           notif.IsRead,
				CreatedAt:        convert.MustPgTimestampToTime(notif.CreatedAt),
			}
			e.publisher.Publish(event)
		})
	} else {
		params.Status = db.ProjectStatusRejected
		notif, err := e.repository.CreateNotification(ctx, db.CreateNotificationParams{
			UserID:           user.ID,
			NotificationType: db.NotificationTypeRejectedProject,
			Message:          fmt.Sprintf("We are sorry! Your project \"%s\" has insufficient requirements and can not be approved of funding.", project.Title),
			ProjectID:        &project.ID,
		})
		if err != nil {
			return err
		}

		helper.Background(func() {
			event := model.Notification{
				ID:               notif.ID,
				UserID:           notif.UserID,
				NotificationType: model.NotificationType(notif.NotificationType),
				Message:          notif.Message,
				ProjectID:        notif.ProjectID,
				IsRead:           notif.IsRead,
				CreatedAt:        convert.MustPgTimestampToTime(notif.CreatedAt),
			}
			e.publisher.Publish(event)
		})

		helper.Background(func() {
			payload := map[string]interface{}{
				"rejectReason": *req.RejectReason,
			}

			if err := e.mailer.Send(user.Email, "reject_project_application.tmpl", payload); err != nil {
				slog.Error(err.Error())
			}
		})
	}
	if err := e.repository.UpdateProjectStatus(ctx, params); err != nil {
		return err
	}

	return nil
}

func (e *escrow) ResolveMilestone(ctx context.Context, milestoneID int64, req dto.ResolveMilestoneRequest) error {
	queries := e.repository.(*db.Queries)
	q, tx, err := queries.BeginTX(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	//TODO: create milestone completion, send confirmation email,...
	milestone, err := q.GetMilestoneByID(ctx, milestoneID)
	if err != nil {
		return err
	}

	if err := q.UpdateMilestoneStatus(ctx, milestone.ID); err != nil {
		return err
	}

	params := db.CreateMilestoneCompletionParams{
		MilestoneID:    milestone.ID,
		TransferAmount: req.Amount,
	}

	if req.Description != nil {
		params.TransferNote = req.Description
	}

	if req.Image != nil {
		params.TransferImage = req.Image
	}

	_, err = q.CreateMilestoneCompletion(ctx, params)
	if err != nil {
		return err
	}

	project, err := q.GetProjectByID(ctx, milestone.ProjectID)
	if err != nil {
		return err
	}

	// Check if the project has been finished
	milestones, err := q.GetMilestoneForProject(ctx, project.ID)
	if err != nil {
		return err
	}

	var incomplete int
	for _, m := range milestones {
		if !m.Completed {
			incomplete++
		}
	}

	if incomplete == 0 {
		if err = q.UpdateProjectStatus(ctx, db.UpdateProjectStatusParams{
			ID:     project.ID,
			Status: db.ProjectStatusFinished,
		}); err != nil {
			return err
		}
	}

	user, err := q.GetUserByID(ctx, project.UserID)
	if err != nil {
		return err
	}

	notif, err := q.CreateNotification(ctx, db.CreateNotificationParams{
		UserID:           user.ID,
		NotificationType: db.NotificationTypeMilestoneCompletion,
		Message:          fmt.Sprintf("Congratulations! Milestone \"%s\" in your campaign \"%s\" has been resolved.", milestone.Title, project.Title),
		ProjectID:        &project.ID,
	})
	if err != nil {
		return err
	}

	helper.Background(func() {
		e.publisher.Publish(model.Notification{
			ID:               notif.ID,
			UserID:           notif.UserID,
			NotificationType: model.NotificationType(notif.NotificationType),
			Message:          notif.Message,
			ProjectID:        notif.ProjectID,
			IsRead:           notif.IsRead,
			CreatedAt:        convert.MustPgTimestampToTime(notif.CreatedAt),
		})
	})

	// Send mail
	return tx.Commit(ctx)
}

func (e *escrow) ApproveUserVerification(ctx context.Context, req dto.VerificationApprovalRequest) error {
	user, err := e.repository.GetUserByID(ctx, req.UserID)
	if err != nil {
		return err
	}

	params := db.UpdateVerificationStatusParams{
		ID: user.ID,
	}

	if req.Approved != nil {
		params.VerificationStatus = db.VerificationStatusVerified
		params.VerificationDocumentUrl = user.VerificationDocumentUrl

		helper.Background(func() {
			payload := map[string]interface{}{
				"firstName": user.FirstName,
			}

			if err := e.mailer.Send(user.Email, "approve_verification.tmpl", payload); err != nil {
				slog.Error(err.Error())
			}
		})
	} else {
		params.VerificationStatus = db.VerificationStatusUnverified
		params.VerificationDocumentUrl = nil

		helper.Background(func() {
			payload := map[string]interface{}{
				"firstName":    user.FirstName,
				"rejectReason": *req.RejectReason, // adjust url as needed
			}

			if err := e.mailer.Send(user.Email, "reject_verification.tmpl", payload); err != nil {
				slog.Error(err.Error())
			}
		})
	}

	if err := e.repository.UpdateVerificationStatus(ctx, params); err != nil {
		return err
	}

	return nil
}

func (e *escrow) GetStatistics(ctx context.Context) (*dto.GetStatisticsResponse, error) {
	stats, err := e.repository.GetStatsAggregation(ctx)
	if err != nil {
		return nil, err
	}

	categoriesCount, err := e.repository.GetCategoriesCount(ctx)
	if err != nil {
		return nil, err
	}

	var cc []dto.CategoryCount

	for _, dbCount := range categoriesCount {
		cc = append(cc, dto.CategoryCount{
			ID:    int(dbCount.ID),
			Name:  dbCount.Name,
			Count: dbCount.Count,
		})
	}

	return &dto.GetStatisticsResponse{
		TotalFund:     stats.TotalFund,
		DonationCount: stats.BackingCount,
		ProjectCount: dto.ProjectCount{
			Pending:  stats.ProjectsPending,
			Ongoing:  stats.ProjectsOngoing,
			Finished: stats.ProjectsFinished,
			Rejected: stats.ProjectsRejected,
		},
		PendingVerificationCount: stats.VerificationCount,
		CategoriesCount:          cc,
	}, nil
}
