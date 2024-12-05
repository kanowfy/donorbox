package handler

import (
	"net/http"
	"strconv"

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
	ApproveOfProject(w http.ResponseWriter, r *http.Request)
	ResolveMilestone(w http.ResponseWriter, r *http.Request)
	ApproveUserVerification(w http.ResponseWriter, r *http.Request)
	GetStatistics(w http.ResponseWriter, r *http.Request)
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

func (e *escrow) ApproveOfProject(w http.ResponseWriter, r *http.Request) {
	var req dto.ProjectApprovalRequest

	err := json.ReadJSON(w, r, &req)
	if err != nil {
		httperror.BadRequestResponse(w, r, err)
		return
	}

	if err = e.validator.Struct(req); err != nil {
		httperror.FailedValidationResponse(w, r, err)
		return
	}

	escrow := rcontext.GetEscrowUser(r)

	if err := e.service.ApproveOfProject(r.Context(), escrow.ID, req); err != nil {
		httperror.ServerErrorResponse(w, r, err)
		return
	}

	if err = json.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"message": "project approved successfully",
	}, nil); err != nil {
		httperror.ServerErrorResponse(w, r, err)
	}
}

func (e *escrow) ResolveMilestone(w http.ResponseWriter, r *http.Request) {
	mid, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		httperror.NotFoundResponse(w, r)
		return
	}

	var req dto.ResolveMilestoneRequest

	err = json.ReadJSON(w, r, &req)
	if err != nil {
		httperror.BadRequestResponse(w, r, err)
		return
	}

	if err = e.validator.Struct(req); err != nil {
		httperror.FailedValidationResponse(w, r, err)
		return
	}

	escrow := rcontext.GetEscrowUser(r)

	if err := e.service.ResolveMilestone(r.Context(), escrow.ID, mid, req); err != nil {
		httperror.ServerErrorResponse(w, r, err)
		return
	}

	if err = json.WriteJSON(w, http.StatusCreated, map[string]interface{}{
		"message": "confirming milestone resolution",
	}, nil); err != nil {
		httperror.ServerErrorResponse(w, r, err)
	}
}

func (e *escrow) ApproveUserVerification(w http.ResponseWriter, r *http.Request) {
	var req dto.VerificationApprovalRequest

	err := json.ReadJSON(w, r, &req)
	if err != nil {
		httperror.BadRequestResponse(w, r, err)
		return
	}

	if err = e.validator.Struct(req); err != nil {
		httperror.FailedValidationResponse(w, r, err)
		return
	}

	escrow := rcontext.GetEscrowUser(r)

	if err := e.service.ApproveUserVerification(r.Context(), escrow.ID, req); err != nil {
		httperror.ServerErrorResponse(w, r, err)
		return
	}

	if err = json.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"message": "project approved successfully",
	}, nil); err != nil {
		httperror.ServerErrorResponse(w, r, err)
	}
}

func (e *escrow) GetStatistics(w http.ResponseWriter, r *http.Request) {
	stats, err := e.service.GetStatistics(r.Context())
	if err != nil {
		httperror.ServerErrorResponse(w, r, err)
		return
	}

	if err = json.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"statistics": stats,
	}, nil); err != nil {
		httperror.ServerErrorResponse(w, r, err)
	}
}
