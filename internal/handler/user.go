package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/kanowfy/donorbox/internal/dto"
	"github.com/kanowfy/donorbox/internal/rcontext"
	"github.com/kanowfy/donorbox/internal/service"
	"github.com/kanowfy/donorbox/pkg/httperror"
	"github.com/kanowfy/donorbox/pkg/json"
)

type User interface {
	GetAuthenticatedUser(w http.ResponseWriter, r *http.Request)
	GetUserByID(w http.ResponseWriter, r *http.Request)
	UpdateAccount(w http.ResponseWriter, r *http.Request)
	ChangePassword(w http.ResponseWriter, r *http.Request)
}

type user struct {
	service   service.User
	validator *validator.Validate
}

func NewUser(service service.User, validator *validator.Validate) User {
	return &user{
		service,
		validator,
	}
}

func (u *user) GetAuthenticatedUser(w http.ResponseWriter, r *http.Request) {
	user := rcontext.GetUser(r)

	if err := json.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"user": user,
	}, nil); err != nil {
		httperror.ServerErrorResponse(w, r, err)
	}
}

func (u *user) GetUserByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		httperror.NotFoundResponse(w, r)
		return
	}

	user, err := u.service.GetUserByID(r.Context(), id)
	if err != nil {
		httperror.NotFoundResponse(w, r)
		return
	}

	if err := json.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"user": user,
	}, nil); err != nil {
		httperror.ServerErrorResponse(w, r, err)
	}
}

func (u *user) UpdateAccount(w http.ResponseWriter, r *http.Request) {
	var req dto.UpdateAccountRequest

	err := json.ReadJSON(w, r, &req)
	if err != nil {
		httperror.BadRequestResponse(w, r, err)
		return
	}

	if err = u.validator.Struct(req); err != nil {
		httperror.FailedValidationResponse(w, r, err)
		return
	}

	user := rcontext.GetUser(r)

	err = u.service.UpdateAccount(r.Context(), user, req)
	if err != nil {
		httperror.BadRequestResponse(w, r, err)
		return
	}

	if err = json.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"message": "profile updated successfully",
	}, nil); err != nil {
		httperror.ServerErrorResponse(w, r, err)
	}
}

func (u *user) ChangePassword(w http.ResponseWriter, r *http.Request) {
	var req dto.ChangePasswordRequest

	err := json.ReadJSON(w, r, &req)
	if err != nil {
		httperror.BadRequestResponse(w, r, err)
		return
	}

	if err = u.validator.Struct(req); err != nil {
		httperror.FailedValidationResponse(w, r, err)
		return
	}

	user := rcontext.GetUser(r)

	err = u.service.ChangePassword(r.Context(), user.ID, req)
	if err != nil {
		if errors.Is(err, service.ErrWrongPassword) {
			httperror.BadRequestResponse(w, r, err)
		} else {
			httperror.ServerErrorResponse(w, r, err)
		}
		return
	}

	if err = json.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"message": "password changed successfully",
	}, nil); err != nil {
		httperror.ServerErrorResponse(w, r, err)
	}
}
