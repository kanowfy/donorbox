package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/kanowfy/donorbox/internal/convert"
	"github.com/kanowfy/donorbox/internal/db"
	"github.com/kanowfy/donorbox/internal/dto"
	"github.com/kanowfy/donorbox/internal/model"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type User interface {
	GetUserByID(ctx context.Context, userID int64) (*model.User, error)
	UpdateAccount(ctx context.Context, user *model.User, req dto.UpdateAccountRequest) error
	ChangePassword(ctx context.Context, userID int64, req dto.ChangePasswordRequest) error
	UploadDocument(ctx context.Context, userID int64, docLink string) error
	GetPendingVerificationUsers(ctx context.Context) ([]dto.PendingUserVerificationResponse, error)
}

type user struct {
	repository db.Querier
	auditSvc   AuditTrail
}

func NewUser(repository db.Querier, auditSvc AuditTrail) User {
	return &user{
		repository,
		auditSvc,
	}
}

func (u *user) GetUserByID(ctx context.Context, userID int64) (*model.User, error) {
	user, err := u.repository.GetUserByID(ctx, userID)
	if err != nil {
		return nil, ErrUserNotFound
	}

	return &model.User{
		ID:                 user.ID,
		Email:              user.Email,
		FirstName:          user.FirstName,
		LastName:           user.LastName,
		ProfilePicture:     user.ProfilePicture,
		Activated:          user.Activated,
		VerificationStatus: model.VerificationStatus(user.VerificationStatus),
		UserType:           model.REGULAR,
		CreatedAt:          convert.MustPgTimestampToTime(user.CreatedAt),
	}, nil
}

func (u *user) UpdateAccount(ctx context.Context, user *model.User, req dto.UpdateAccountRequest) error {
	queries := u.repository.(*db.Queries)
	q, tx, err := queries.BeginTX(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	var updateParams db.UpdateUserByIDParams
	updateParams.ID = user.ID

	trailParams := LogActionParams{
		UserID:        &user.ID,
		EntityType:    "user",
		EntityID:      &user.ID,
		OperationType: "UPDATE",
	}

	if req.Email != nil {
		trailParams.FieldName = "email"
		trailParams.OldValue = updateParams.Email
		trailParams.OldValue = *req.Email

		updateParams.Email = *req.Email
	} else {
		updateParams.Email = user.Email
	}

	if req.FirstName != nil {
		trailParams.FieldName = "first_name"
		trailParams.OldValue = updateParams.FirstName
		trailParams.OldValue = *req.FirstName

		updateParams.FirstName = *req.FirstName
	} else {
		updateParams.FirstName = user.FirstName
	}

	if req.LastName != nil {
		trailParams.FieldName = "last_name"
		trailParams.OldValue = updateParams.LastName
		trailParams.OldValue = *req.LastName

		updateParams.LastName = *req.LastName
	} else {
		updateParams.LastName = user.LastName
	}

	if req.ProfilePicture != nil {
		trailParams.FieldName = "profile_picture"
		trailParams.OldValue = updateParams.ProfilePicture
		trailParams.OldValue = *req.ProfilePicture

		updateParams.ProfilePicture = req.ProfilePicture
	} else {
		updateParams.ProfilePicture = user.ProfilePicture
	}

	if err := q.UpdateUserByID(ctx, updateParams); err != nil {
		return fmt.Errorf("update user: %w", err)
	}

	if err := u.auditSvc.LogAction(ctx, trailParams); err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (u *user) ChangePassword(ctx context.Context, userID int64, req dto.ChangePasswordRequest) error {
	queries := u.repository.(*db.Queries)
	q, tx, err := queries.BeginTX(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	user, err := q.GetUserByID(ctx, userID)
	if err != nil {
		return err
	}

	// check password
	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(req.OldPassword))
	if err != nil {
		return ErrWrongPassword
	}

	newHashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)

	args := db.UpdateUserPasswordParams{
		ID:             user.ID,
		HashedPassword: string(newHashedPassword),
	}

	if err = q.UpdateUserPassword(ctx, args); err != nil {
		return err
	}

	if err = u.auditSvc.LogAction(ctx, LogActionParams{
		UserID:        &user.ID,
		EntityType:    "user",
		EntityID:      &user.ID,
		OperationType: "UPDATE",
		FieldName:     "password",
	}); err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (u *user) UploadDocument(ctx context.Context, userID int64, docLink string) error {
	queries := u.repository.(*db.Queries)
	q, tx, err := queries.BeginTX(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	if err := q.UpdateVerificationStatus(ctx, db.UpdateVerificationStatusParams{
		ID:                      userID,
		VerificationStatus:      db.VerificationStatusPending,
		VerificationDocumentUrl: &docLink,
	}); err != nil {
		return err
	}

	if err := u.auditSvc.LogAction(ctx, LogActionParams{
		UserID:        &userID,
		EntityType:    "user",
		EntityID:      &userID,
		OperationType: "UPDATE",
		FieldName:     "verification_document_url", // how about status?
		NewValue:      docLink,
	}); err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (u *user) GetPendingVerificationUsers(ctx context.Context) ([]dto.PendingUserVerificationResponse, error) {
	dbUsers, err := u.repository.GetPendingVerificationUsers(ctx)
	if err != nil {
		return nil, err
	}

	var users []dto.PendingUserVerificationResponse
	for _, u := range dbUsers {
		users = append(users, dto.PendingUserVerificationResponse{
			ID:          u.ID,
			Email:       u.Email,
			FirstName:   u.FirstName,
			LastName:    u.LastName,
			DocumentUrl: *u.VerificationDocumentUrl,
			CreatedAt:   convert.MustPgTimestampToTime(u.CreatedAt),
		})
	}

	return users, nil
}
