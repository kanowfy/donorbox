package handler

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/kanowfy/donorbox/internal/dto"
	"github.com/kanowfy/donorbox/internal/rcontext"
	"github.com/kanowfy/donorbox/internal/service"
	"github.com/kanowfy/donorbox/pkg/httperror"
	"github.com/kanowfy/donorbox/pkg/json"
)

type Escrow interface {
	Login(w http.ResponseWriter, r *http.Request)
	GetAuthenticatedEscrow(w http.ResponseWriter, r *http.Request)
}

type escrow struct {
	service   service.Escrow
	validator *validator.Validate
}

func NewEscrow(service service.Escrow, validator *validator.Validate) Escrow {
	return &escrow{
		service,
		validator,
	}
}

func (e *escrow) Login(w http.ResponseWriter, r *http.Request) {
	var req dto.LoginRequest

	err := json.ReadJSON(w, r, &req)
	if err != nil {
		httperror.BadRequestResponse(w, r, err)
		return
	}

	if err = e.validator.Struct(req); err != nil {
		httperror.FailedValidationResponse(w, r, err)
		return
	}

	token, err := e.service.Login(r.Context(), req)
	if err != nil {
		httperror.InvalidCredentialsResponse(w, r)
		return
	}

	// return token
	if err := json.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"token": token,
	}, nil); err != nil {
		httperror.ServerErrorResponse(w, r, err)
	}
}

func (e *escrow) GetAuthenticatedEscrow(w http.ResponseWriter, r *http.Request) {
	escrow := rcontext.GetEscrowUser(r)

	if err := json.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"user": escrow,
	}, nil); err != nil {
		httperror.ServerErrorResponse(w, r, err)
	}
}
