package service

import (
	"context"
	"errors"
	"time"

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
	ConfirmResolvedMilestone(ctx context.Context, milestoneID int64) error
}

type user struct {
	repository db.Querier
}

func NewUser(repository db.Querier) User {
	return &user{
		repository,
	}
}

func (u *user) GetUserByID(ctx context.Context, userID int64) (*model.User, error) {
	user, err := u.repository.GetUserByID(ctx, userID)
	if err != nil {
		return nil, ErrUserNotFound
	}

	return &model.User{
		ID:             user.ID,
		Email:          user.Email,
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		ProfilePicture: user.ProfilePicture,
		Activated:      user.Activated,
		UserType:       model.REGULAR,
		CreatedAt:      user.CreatedAt,
	}, nil
}

func (u *user) UpdateAccount(ctx context.Context, user *model.User, req dto.UpdateAccountRequest) error {
	var updateParams db.UpdateUserByIDParams
	updateParams.ID = user.ID

	if req.Email != nil {
		updateParams.Email = *req.Email
	} else {
		updateParams.Email = user.Email
	}

	if req.FirstName != nil {
		updateParams.FirstName = *req.FirstName
	} else {
		updateParams.FirstName = user.FirstName
	}

	if req.LastName != nil {
		updateParams.LastName = *req.LastName
	} else {
		updateParams.LastName = user.LastName
	}

	if req.ProfilePicture != nil {
		updateParams.ProfilePicture = req.ProfilePicture
	} else {
		updateParams.ProfilePicture = user.ProfilePicture
	}

	if err := u.repository.UpdateUserByID(ctx, updateParams); err != nil {
		return err
	}

	return nil
}

func (u *user) ChangePassword(ctx context.Context, userID int64, req dto.ChangePasswordRequest) error {
	user, err := u.repository.GetUserByID(ctx, userID)
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

	if err = u.repository.UpdateUserPassword(ctx, args); err != nil {
		return err
	}

	return nil
}

func (u *user) ConfirmResolvedMilestone(ctx context.Context, milestoneID int64) error {
	if err := u.repository.UpdateVerifyingCertificate(ctx, db.UpdateVerifyingCertificateParams{
		MilestoneID: milestoneID,
		VerifiedAt:  time.Now(),
	}); err != nil {
		return err
	}
	//TODO: Generate blockchain stuff and hash

	return nil
}
