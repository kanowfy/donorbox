package service

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/kanowfy/donorbox/internal/convert"
	"github.com/kanowfy/donorbox/internal/db"
	"github.com/kanowfy/donorbox/internal/dto"
	"github.com/kanowfy/donorbox/internal/model"
	"github.com/kanowfy/donorbox/internal/token"
	"golang.org/x/crypto/bcrypt"
)

type Escrow interface {
	Login(ctx context.Context, req dto.LoginRequest) (string, error)
	GetEscrowByID(ctx context.Context, id int64) (*model.EscrowUser, error)
	ApproveOfProject(ctx context.Context, req dto.ProjectApprovalRequest) error
	ResolveMilestone(ctx context.Context, escrowID int64, req dto.ResolveMilestoneRequest) error
	ApproveUserVerification(ctx context.Context, req dto.VerificationApprovalRequest) error
}

type escrow struct {
	repository db.Querier
}

func NewEscrow(querier db.Querier) Escrow {
	return &escrow{
		repository: querier,
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
	params := db.UpdateProjectStatusParams{
		ID: req.ProjectID,
	}

	if req.Approved != nil {
		params.Status = db.ProjectStatusOngoing
	} else {
		//TODO: do something with other req fields
		params.Status = db.ProjectStatusRejected
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
	} else {
		params.VerificationStatus = db.VerificationStatusUnverified
		params.VerificationDocumentUrl = nil
		// Send email on rejection
	}

	if err := e.repository.UpdateVerificationStatus(ctx, params); err != nil {
		return err
	}

	return nil
}
