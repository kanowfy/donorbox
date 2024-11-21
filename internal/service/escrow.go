package service

import (
	"context"
	"strconv"
	"time"

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
	ResolveMilestone(ctx context.Context, escrowID int64, milestoneID int64) error
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
		CreatedAt: escrow.CreatedAt,
	}, nil
}

func (e *escrow) ApproveOfProject(ctx context.Context, req dto.ProjectApprovalRequest) error {
	pid, err := strconv.ParseInt(req.ProjectID, 10, 64)
	if err != nil {
		return err
	}

	var status db.NullProjectStatus
	if req.Approved {
		status.Scan(db.ProjectStatusOngoing)
	} else {
		//TODO: do something with other req fields
		status.Scan(db.ProjectStatusRejected)
	}
	if err := e.repository.UpdateProjectStatus(ctx, db.UpdateProjectStatusParams{
		ID:     pid,
		Status: status,
	}); err != nil {
		return err
	}

	return nil
}

func (e *escrow) ResolveMilestone(ctx context.Context, escrowID int64, milestoneID int64) error {
	//TODO: create cert, send confirmation email,...
	milestone, err := e.repository.GetMilestoneByID(ctx, milestoneID)
	if err != nil {
		return err
	}

	project, err := e.repository.GetProjectByID(ctx, milestone.ProjectID)
	if err != nil {
		return err
	}

	_, err = e.repository.CreateCertificate(ctx, db.CreateCertificateParams{
		EscrowUserID: escrowID,
		UserID:       project.UserID,
		MilestoneID:  milestoneID,
	})
	if err != nil {
		return err
	}

	// Send mail
	return nil
}
