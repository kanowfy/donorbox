package service

import (
	"context"
	"fmt"
	"log/slog"
	"math/big"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/kanowfy/donorbox/internal/contract"
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
	ApproveOfProject(ctx context.Context, escrowID int64, req dto.ProjectApprovalRequest) error
	ResolveMilestone(ctx context.Context, escorwID int64, milestoneID int64, req dto.ResolveMilestoneRequest) error
	ApproveUserVerification(ctx context.Context, escrowID int64, req dto.VerificationApprovalRequest) error
	ApproveSpendingProof(ctx context.Context, escrowID int64, req dto.ProofApprovalRequest) error
	ReviewReport(ctx context.Context, escrowID int64, req dto.ReportReviewRequest) error
	GetStatistics(ctx context.Context) (*dto.GetStatisticsResponse, error)
	ResolveDispute(ctx context.Context, escrowID int64, req dto.DisputedProjectActionRequest) error
}

type escrow struct {
	repository db.Querier
	mailer     mail.Mailer
	publisher  publish.Publisher
	auditSvc   AuditTrail
	transactor *contract.BlockchainTransactor
}

func NewEscrow(querier db.Querier, mailer mail.Mailer, publisher publish.Publisher, auditSvc AuditTrail, transactor *contract.BlockchainTransactor) Escrow {
	return &escrow{
		repository: querier,
		mailer:     mailer,
		publisher:  publisher,
		auditSvc:   auditSvc,
		transactor: transactor,
	}
}

const MAX_PROOF_COUNT = 3

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

func (e *escrow) ApproveOfProject(ctx context.Context, escrowID int64, req dto.ProjectApprovalRequest) error {
	queries := e.repository.(*db.Queries)
	q, tx, err := queries.BeginTX(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	project, err := q.GetProjectByID(ctx, req.ProjectID)
	if err != nil {
		return err
	}

	user, err := q.GetUserByID(ctx, project.UserID)
	if err != nil {
		return err
	}

	params := db.UpdateProjectStatusParams{
		ID: project.ID,
	}

	trailParams := LogActionParams{
		EscrowID:      &escrowID,
		EntityType:    "project",
		EntityID:      &project.ID,
		OperationType: "UPDATE",
		FieldName:     "status",
		OldValue:      db.ProjectStatusPending,
	}

	if req.Approved != nil {
		params.Status = db.ProjectStatusOngoing
		trailParams.NewValue = db.ProjectStatusOngoing
		notif, err := q.CreateNotification(ctx, db.CreateNotificationParams{
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
		trailParams.NewValue = db.ProjectStatusRejected
		notif, err := q.CreateNotification(ctx, db.CreateNotificationParams{
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

	if err := q.UpdateProjectStatus(ctx, params); err != nil {
		return err
	}

	if err := e.auditSvc.LogAction(ctx, trailParams); err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (e *escrow) ResolveMilestone(ctx context.Context, escrowID int64, milestoneID int64, req dto.ResolveMilestoneRequest) error {
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

	completion, err := q.CreateMilestoneCompletion(ctx, params)
	if err != nil {
		return err
	}

	if err := e.auditSvc.LogAction(ctx, LogActionParams{
		EscrowID:      &escrowID,
		EntityType:    "escrow_milestone_completion",
		EntityID:      &completion.ID,
		OperationType: "CREATE",
		NewValue:      completion,
	}); err != nil {
		return err
	}

	if err := q.UpdateMilestoneStatus(ctx, db.UpdateMilestoneStatusParams{
		ID:     milestoneID,
		Status: db.MilestoneStatusFundReleased,
	}); err != nil {
		return err
	}

	if err := e.auditSvc.LogAction(ctx, LogActionParams{
		EscrowID:      &escrowID,
		EntityType:    "milestone",
		EntityID:      &milestoneID,
		OperationType: "UPDATE",
		FieldName:     "status",
		OldValue:      db.MilestoneStatusPending,
		NewValue:      db.MilestoneStatusFundReleased,
	}); err != nil {
		return err
	}

	project, err := q.GetProjectByID(ctx, milestone.ProjectID)
	if err != nil {
		return err
	}

	//TODO: Move this over when escrow confirm user proof
	// Check if the project has been finished
	/*
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

			if err := e.auditSvc.LogAction(ctx, LogActionParams{
				EscrowID:      &escrowID,
				EntityType:    "project",
				EntityID:      &project.ID,
				OperationType: "UPDATE",
				FieldName:     "status",
				OldValue:      db.ProjectStatusOngoing,
				NewValue:      db.ProjectStatusFinished,
			}); err != nil {
				return err
			}
		}
	*/

	user, err := q.GetUserByID(ctx, project.UserID)
	if err != nil {
		return err
	}

	notif, err := q.CreateNotification(ctx, db.CreateNotificationParams{
		UserID:           user.ID,
		NotificationType: db.NotificationTypeReleasedFundMilestone,
		Message:          fmt.Sprintf("Funding for milestone \"%s\" of your project \"%s\" has been released", milestone.Title, project.Title),
		ProjectID:        &project.ID,
		MilestoneID:      &milestone.ID,
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

	var tfNote string
	if completion.TransferNote != nil {
		tfNote = *completion.TransferNote
	}

	// create blockchain tx
	txn, err := e.transactor.Contract.StoreFundRelease(e.transactor.AuthData, big.NewInt(completion.ID), uint64(project.ID), uint64(milestone.ID), *completion.TransferImage, tfNote, completion.CreatedAt.Time.String())
	if err != nil {
		return err
	}
	slog.Info("new transaction processed", "txn_id", txn.Hash().String())

	// Send mail
	return tx.Commit(ctx)
}

func (e *escrow) ApproveUserVerification(ctx context.Context, escrowID int64, req dto.VerificationApprovalRequest) error {
	queries := e.repository.(*db.Queries)
	q, tx, err := queries.BeginTX(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	user, err := q.GetUserByID(ctx, req.UserID)
	if err != nil {
		return err
	}

	params := db.UpdateVerificationStatusParams{
		ID: user.ID,
	}

	trailParams := LogActionParams{
		EscrowID:      &escrowID,
		EntityType:    "user",
		EntityID:      &user.ID,
		OperationType: "UPDATE",
		FieldName:     "verification_status",
		OldValue:      db.VerificationStatusPending,
	}

	if req.Approved != nil {
		params.VerificationStatus = db.VerificationStatusVerified
		params.VerificationDocumentUrl = user.VerificationDocumentUrl
		trailParams.NewValue = db.VerificationStatusVerified

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
		trailParams.NewValue = db.VerificationStatusVerified

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

	if err := q.UpdateVerificationStatus(ctx, params); err != nil {
		return err
	}

	if err := e.auditSvc.LogAction(ctx, trailParams); err != nil {
		return err
	}

	return tx.Commit(ctx)
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

func (e *escrow) ApproveSpendingProof(ctx context.Context, escrowID int64, req dto.ProofApprovalRequest) error {
	queries := e.repository.(*db.Queries)
	q, tx, err := queries.BeginTX(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	proof, err := q.GetSpendingProofByID(ctx, req.ProofID)
	if err != nil {
		return err
	}

	milestone, err := q.GetMilestoneByID(ctx, proof.MilestoneID)
	if err != nil {
		return err
	}

	project, err := q.GetProjectByID(ctx, milestone.ProjectID)
	if err != nil {
		return err
	}

	// when this method is executed, we know user still has chance to upload
	if req.Approved != nil {
		if err := q.UpdateSpendingProofStatus(ctx, db.UpdateSpendingProofStatusParams{
			ID:     proof.ID,
			Status: db.ProofStatusApproved,
		}); err != nil {
			return err
		}

		if err := q.UpdateMilestoneStatus(ctx, db.UpdateMilestoneStatusParams{
			ID:     milestone.ID,
			Status: db.MilestoneStatusCompleted,
		}); err != nil {
			return err
		}

		milestones, err := q.GetMilestoneForProject(ctx, project.ID)
		if err != nil {
			return err
		}

		var incomplete int
		for _, m := range milestones {
			if m.Status != db.MilestoneStatusCompleted {
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

		notif, err := q.CreateNotification(ctx, db.CreateNotificationParams{
			UserID:           project.UserID,
			NotificationType: db.NotificationTypeApprovedProof,
			Message: fmt.Sprintf("Proof of expenditure for milestone \"%s\" has been approved!",
				milestone.Title),
			MilestoneID: &milestone.ID,
			ProjectID:   &milestone.ProjectID,
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

		// create blockchain tx
		txn, err := e.transactor.Contract.StoreVerifiedProof(e.transactor.AuthData, big.NewInt(proof.ID), uint64(project.ID), uint64(milestone.ID), proof.TransferImage, proof.ProofMediaUrl, proof.CreatedAt.Time.String())
		if err != nil {
			return err
		}
		slog.Info("new transaction processed", "txn_id", txn.Hash().String())
	} else {
		if err := q.UpdateSpendingProofStatus(ctx, db.UpdateSpendingProofStatusParams{
			ID:            proof.ID,
			Status:        db.ProofStatusRejected,
			RejectedCause: req.RejectReason,
		}); err != nil {
			return err
		}

		notif, err := q.CreateNotification(ctx, db.CreateNotificationParams{
			UserID:           project.UserID,
			NotificationType: db.NotificationTypeRejectedProof,
			Message: fmt.Sprintf("We are sorry, proof of expenditure for milestone \"%s\" is invalid: %s",
				milestone.Title, *req.RejectReason),
			MilestoneID: &milestone.ID,
			ProjectID:   &project.ID,
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

		// check if the count of rejected proof now is max
		proofs, err := q.GetSpendingProofsForMilestone(ctx, milestone.ID)
		if err != nil {
			return err
		}

		var count int
		for _, p := range proofs {
			if p.Status == db.ProofStatusRejected {
				count++
			}
		}

		if count == MAX_PROOF_COUNT {
			if err := q.UpdateMilestoneStatus(ctx, db.UpdateMilestoneStatusParams{
				ID:     proof.MilestoneID,
				Status: db.MilestoneStatusRefuted,
			}); err != nil {
				return err
			}

			notif, err := q.CreateNotification(ctx, db.CreateNotificationParams{
				UserID:           project.UserID,
				NotificationType: db.NotificationTypeRefutedMilestone,
				Message: fmt.Sprintf("Milestone \"%s\" has been refuted due to the lack of valid proof of expenditure and will be further investigated",
					milestone.Title),
				MilestoneID: &milestone.ID,
				ProjectID:   &milestone.ProjectID,
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
		}
	}

	return tx.Commit(ctx)
}

func (e *escrow) ReviewReport(ctx context.Context, escrowID int64, req dto.ReportReviewRequest) error {
	queries := e.repository.(*db.Queries)
	q, tx, err := queries.BeginTX(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	report, err := q.GetProjectReportByID(ctx, req.ReportID)
	if err != nil {
		return fmt.Errorf("review report: %w", err)
	}

	if req.MarkDispute != nil {
		if err := q.UpdateProjectStatus(ctx, db.UpdateProjectStatusParams{
			ID:     report.ProjectID,
			Status: db.ProjectStatusDisputed,
		}); err != nil {
			return fmt.Errorf("review report: %w", err)
		}

		if err := q.UpdateProjectReportStatus(ctx, db.UpdateProjectReportStatusParams{
			ID:     report.ID,
			Status: db.ReportStatusResolved,
		}); err != nil {
			return fmt.Errorf("review report: %w", err)
		}
	} else {
		if err := q.UpdateProjectReportStatus(ctx, db.UpdateProjectReportStatusParams{
			ID:     report.ID,
			Status: db.ReportStatusDismissed,
		}); err != nil {
			return fmt.Errorf("review report: %w", err)
		}
	}

	return tx.Commit(ctx)
}

func (e *escrow) ResolveDispute(ctx context.Context, escrowID int64, req dto.DisputedProjectActionRequest) error {
	queries := e.repository.(*db.Queries)
	q, tx, err := queries.BeginTX(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	project, err := q.GetProjectByID(ctx, req.ProjectID)
	if err != nil {
		return fmt.Errorf("resolve dispute: %w", err)
	}

	if req.MarkStopped != nil {
		if err = q.UpdateProjectStatus(ctx, db.UpdateProjectStatusParams{
			ID:     req.ProjectID,
			Status: db.ProjectStatusStopped,
		}); err != nil {
			return fmt.Errorf("resolve dispute: %w", err)
		}
	} else if req.MarkReconciled != nil {
		if project.TotalFund >= project.FundGoal {
			if err = q.UpdateProjectStatus(ctx, db.UpdateProjectStatusParams{
				ID:     req.ProjectID,
				Status: db.ProjectStatusFinished,
			}); err != nil {
				return fmt.Errorf("resolve dispute: %w", err)
			}
		} else {
			if err = q.UpdateProjectStatus(ctx, db.UpdateProjectStatusParams{
				ID:     req.ProjectID,
				Status: db.ProjectStatusOngoing,
			}); err != nil {
				return fmt.Errorf("resolve dispute: %w", err)
			}
		}
	}

	return tx.Commit(ctx)
}
