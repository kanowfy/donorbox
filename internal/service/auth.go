package service

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/kanowfy/donorbox/internal/convert"
	"github.com/kanowfy/donorbox/internal/db"
	"github.com/kanowfy/donorbox/internal/dto"
	"github.com/kanowfy/donorbox/internal/mail"
	"github.com/kanowfy/donorbox/internal/model"
	"github.com/kanowfy/donorbox/internal/token"
	"github.com/kanowfy/donorbox/pkg/helper"
	"github.com/markbates/goth"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrWrongPassword  = errors.New("wrong password")
	ErrEmailNotExists = errors.New("email does not exist")
	ErrUserActivated  = errors.New("user already activated")
)

type Auth interface {
	Login(ctx context.Context, request dto.LoginRequest) (string, error)
	Register(ctx context.Context, request dto.UserRegisterRequest, hostPath string) (*model.User, error)
	RegisterEscrow(ctx context.Context, req dto.EscrowRegisterRequest) (*model.EscrowUser, error)
	ActivateAccount(ctx context.Context, activationToken string) error
	SendResetPasswordToken(ctx context.Context, email string, hostPath string) error
	ResetPassword(ctx context.Context, request dto.ResetPasswordRequest) error
	LoginOAuth(ctx context.Context, oauthUser goth.User) (string, error)
}

type auth struct {
	repository db.Querier
	mailer     mail.Mailer
}

func NewAuth(querier db.Querier, mailer mail.Mailer) Auth {
	return &auth{
		repository: querier,
		mailer:     mailer,
	}
}

func (a *auth) Login(ctx context.Context, req dto.LoginRequest) (string, error) {
	user, err := a.repository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return "", ErrUserNotFound
	}

	// validate password
	if err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(req.Password)); err != nil {
		return "", ErrWrongPassword
	}

	token, err := token.GenerateToken(user.ID, time.Hour*3*24)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (a *auth) Register(ctx context.Context, req dto.UserRegisterRequest, hostPath string) (*model.User, error) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	args := db.CreateUserParams{
		Email:          req.Email,
		HashedPassword: string(hashedPassword),
		FirstName:      req.FirstName,
		LastName:       req.LastName,
	}

	user, err := a.repository.CreateUser(ctx, args)
	if err != nil {
		return nil, err
	}

	token, err := token.GenerateToken(user.ID, time.Hour*3*24)
	if err != nil {
		return nil, err
	}

	helper.Background(func() {
		payload := map[string]interface{}{
			"activationUrl": fmt.Sprintf("%s/verify?token=%s", hostPath, token), // adjust url as needed
		}

		if err := a.mailer.Send(req.Email, "registration.tmpl", payload); err != nil {
			slog.Error(err.Error())
		}
	})

	return &model.User{
		ID:             user.ID,
		Email:          user.Email,
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		ProfilePicture: user.ProfilePicture,
		Activated:      user.Activated,
		UserType:       model.REGULAR,
		CreatedAt:      convert.MustPgTimestampToTime(user.CreatedAt),
	}, nil
}

func (a *auth) RegisterEscrow(ctx context.Context, req dto.EscrowRegisterRequest) (*model.EscrowUser, error) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	args := db.CreateEscrowUserParams{
		Email:          req.Email,
		HashedPassword: string(hashedPassword),
	}

	escrow, err := a.repository.CreateEscrowUser(ctx, args)
	if err != nil {
		return nil, err
	}

	return &model.EscrowUser{
		ID:        escrow.ID,
		Email:     escrow.Email,
		UserType:  model.ESCROW,
		CreatedAt: convert.MustPgTimestampToTime(escrow.CreatedAt),
	}, nil
}

func (a *auth) ActivateAccount(ctx context.Context, activationToken string) error {
	userID, err := token.VerifyToken(activationToken)
	if err != nil {
		return err
	}

	user, err := a.repository.GetUserByID(ctx, userID)
	if err != nil {
		return ErrUserNotFound
	}

	if user.Activated {
		return ErrUserActivated
	}

	if err = a.repository.ActivateUser(ctx, user.ID); err != nil {
		return err
	}

	return nil
}

func (a *auth) SendResetPasswordToken(ctx context.Context, email string, hostPath string) error {
	user, err := a.repository.GetUserByEmail(ctx, email)
	if err != nil {
		return ErrEmailNotExists
	}

	token, err := token.GenerateToken(user.ID, time.Minute*15)
	if err != nil {
		return err
	}

	helper.Background(func() {
		payload := map[string]interface{}{
			"firstName":        user.FirstName,
			"resetPasswordUrl": fmt.Sprintf("%s/password/reset?token=%s", hostPath, token), // adjust url as needed
		}

		if err := a.mailer.Send(user.Email, "resetpassword.tmpl", payload); err != nil {
			slog.Error(err.Error())
		}
	})

	return nil
}

func (a *auth) ResetPassword(ctx context.Context, req dto.ResetPasswordRequest) error {
	id, err := token.VerifyToken(req.ResetToken)
	if err != nil {
		return err
	}

	user, err := a.repository.GetUserByID(ctx, id)
	if err != nil {
		return ErrUserNotFound
	}

	newHashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)

	args := db.UpdateUserPasswordParams{
		ID:             user.ID,
		HashedPassword: string(newHashedPassword),
	}

	// consider returning a sentinel error
	if err = a.repository.UpdateUserPassword(ctx, args); err != nil {
		return err
	}

	return nil
}

func (a *auth) LoginOAuth(ctx context.Context, oauthUser goth.User) (string, error) {
	var user db.User

	user, err := a.repository.GetUserByEmail(ctx, oauthUser.Email)
	if err != nil {
		params := db.CreateSocialLoginUserParams{
			Email:          oauthUser.Email,
			FirstName:      oauthUser.FirstName,
			LastName:       oauthUser.LastName,
			ProfilePicture: &oauthUser.AvatarURL,
		}

		if params.FirstName == "" {
			params.FirstName = "Anonymous"
		}

		user, err = a.repository.CreateSocialLoginUser(ctx, params)

		if err != nil {
			return "", err
		}
	}

	token, err := token.GenerateToken(user.ID, time.Hour*3*24)
	if err != nil {
		return "", err
	}

	return token, nil
}
