package handler

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/kanowfy/donorbox/internal/config"
	"github.com/kanowfy/donorbox/internal/dto"
	"github.com/kanowfy/donorbox/internal/service"
	"github.com/kanowfy/donorbox/pkg/helper"
	"github.com/kanowfy/donorbox/pkg/httperror"
	"github.com/kanowfy/donorbox/pkg/json"
	"github.com/markbates/goth/gothic"
)

type Auth interface {
	Login(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
	RegisterEscrow(w http.ResponseWriter, r *http.Request)
	ActivateUser(w http.ResponseWriter, r *http.Request)
	ForgotPassword(w http.ResponseWriter, r *http.Request)
	ResetPassword(w http.ResponseWriter, r *http.Request)
	GetAuthenticationToken(w http.ResponseWriter, r *http.Request)
	GoogleAuthCallback(w http.ResponseWriter, r *http.Request)
	StartOAuth(w http.ResponseWriter, r *http.Request)
}

type auth struct {
	service   service.Auth
	validator *validator.Validate
	cfg       config.Config
}

func NewAuth(service service.Auth, validator *validator.Validate, cfg config.Config) Auth {
	return &auth{
		service,
		validator,
		cfg,
	}
}

func (a *auth) Login(w http.ResponseWriter, r *http.Request) {
	var req dto.LoginRequest

	err := json.ReadJSON(w, r, &req)
	if err != nil {
		httperror.BadRequestResponse(w, r, err)
		return
	}

	if err = a.validator.Struct(req); err != nil {
		httperror.FailedValidationResponse(w, r, err)
		return
	}

	token, err := a.service.Login(r.Context(), req)
	if err != nil {
		switch err {
		case service.ErrUserNotFound:
			httperror.NotFoundResponse(w, r)
		case service.ErrWrongPassword:
			httperror.InvalidCredentialsResponse(w, r)
		default:
			httperror.ServerErrorResponse(w, r, err)
		}
		return

	}

	if err := json.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"token": token,
	}, nil); err != nil {
		httperror.ServerErrorResponse(w, r, err)
	}
}

func (a *auth) Register(w http.ResponseWriter, r *http.Request) {
	var req dto.UserRegisterRequest

	err := json.ReadJSON(w, r, &req)
	if err != nil {
		httperror.BadRequestResponse(w, r, err)
		return
	}

	if err = a.validator.Struct(req); err != nil {
		httperror.FailedValidationResponse(w, r, err)
		return
	}

	user, err := a.service.Register(r.Context(), req, a.cfg.ClientUrl)
	if err != nil {
		httperror.BadRequestResponse(w, r, err)
		return
	}

	if err = json.WriteJSON(w, http.StatusAccepted, map[string]interface{}{
		"user": user,
	}, nil); err != nil {
		httperror.ServerErrorResponse(w, r, err)
	}
}

// TODO: Add route protection
func (a *auth) RegisterEscrow(w http.ResponseWriter, r *http.Request) {
	var req dto.EscrowRegisterRequest

	err := json.ReadJSON(w, r, &req)
	if err != nil {
		httperror.BadRequestResponse(w, r, err)
		return
	}

	if err = a.validator.Struct(req); err != nil {
		httperror.FailedValidationResponse(w, r, err)
		return
	}

	user, err := a.service.RegisterEscrow(r.Context(), req)
	if err != nil {
		httperror.BadRequestResponse(w, r, err)
		return
	}

	if err = json.WriteJSON(w, http.StatusAccepted, map[string]interface{}{
		"escrow_user": user,
	}, nil); err != nil {
		httperror.ServerErrorResponse(w, r, err)
	}
}

func (a *auth) ActivateUser(w http.ResponseWriter, r *http.Request) {
	tokenString := helper.ReadString(r.URL.Query(), "token", "")
	if tokenString == "" {
		httperror.BadRequestResponse(w, r, errors.New("missing token"))
		return
	}

	err := a.service.ActivateAccount(r.Context(), tokenString)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrUserNotFound):
			httperror.NotFoundResponse(w, r)
		case errors.Is(err, service.ErrUserActivated):
			httperror.BadRequestResponse(w, r, err)
		default:
			httperror.ServerErrorResponse(w, r, err)
		}
		return
	}

	if err = json.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"message": "profile updated successfully",
	}, nil); err != nil {
		httperror.ServerErrorResponse(w, r, err)
	}
}

func (a *auth) ForgotPassword(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email string `json:"email" validate:"required,email"`
	}

	err := json.ReadJSON(w, r, &req)
	if err != nil {
		httperror.BadRequestResponse(w, r, err)
		return
	}

	if err = a.validator.Struct(req); err != nil {
		httperror.FailedValidationResponse(w, r, err)
		return
	}

	// check if email exist
	// if not send 201 Accepted, if yet, generate jwt to put in a link
	err = a.service.SendResetPasswordToken(r.Context(), req.Email, a.cfg.ClientUrl)
	if err != nil {
		if !errors.Is(err, service.ErrEmailNotExists) {
			httperror.ServerErrorResponse(w, r, err)
			return
		}
	}

	if err = json.WriteJSON(w, http.StatusOK, nil, nil); err != nil {
		httperror.ServerErrorResponse(w, r, err)
	}
}

func (a *auth) ResetPassword(w http.ResponseWriter, r *http.Request) {
	var req dto.ResetPasswordRequest

	err := json.ReadJSON(w, r, &req)
	if err != nil {
		httperror.BadRequestResponse(w, r, err)
		return
	}

	if err = a.validator.Struct(req); err != nil {
		httperror.FailedValidationResponse(w, r, err)
		return
	}

	err = a.service.ResetPassword(r.Context(), req)
	if err != nil {
		httperror.BadRequestResponse(w, r, err)
	}

	if err = json.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"message": "password reset successfully",
	}, nil); err != nil {
		httperror.ServerErrorResponse(w, r, err)
	}
}

func (a *auth) GoogleAuthCallback(w http.ResponseWriter, r *http.Request) {
	gothUser, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		httperror.ServerErrorResponse(w, r, err)
		return
	}

	fmt.Printf("%+v\n", gothUser)

	token, err := a.service.LoginOAuth(r.Context(), gothUser)
	if err != nil {
		httperror.ServerErrorResponse(w, r, err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 3 * 24),
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
	})

	http.Redirect(w, r, fmt.Sprintf("%s/login/google", a.cfg.ClientUrl), http.StatusFound)
}

func (a *auth) GetAuthenticationToken(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("auth_token")
	if err != nil {
		http.Error(w, "Token not found", http.StatusUnauthorized)
		return
	}

	if err := json.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"token": cookie.Value,
	}, nil); err != nil {
		httperror.ServerErrorResponse(w, r, err)
	}

}

func (a *auth) StartOAuth(w http.ResponseWriter, r *http.Request) {
	provider := r.PathValue("provider")
	r = r.WithContext(context.WithValue(context.Background(), "provider", provider))

	gothic.BeginAuthHandler(w, r)
}
