package handler

import (
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/kanowfy/donorbox/internal/dto"
	"github.com/kanowfy/donorbox/internal/rcontext"
	"github.com/kanowfy/donorbox/internal/service"
	"github.com/kanowfy/donorbox/pkg/httperror"
	"github.com/kanowfy/donorbox/pkg/json"
)

type Card interface {
	GetCard(w http.ResponseWriter, r *http.Request)
	SetupEscrowCard(w http.ResponseWriter, r *http.Request)
	SetupProjectCard(w http.ResponseWriter, r *http.Request)
}

type card struct {
	service   service.Card
	validator *validator.Validate
}

func NewCard(service service.Card, validator *validator.Validate) Card {
	return &card{
		service,
		validator,
	}
}

func (c *card) GetCard(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	cid, err := uuid.Parse(idStr)
	if err != nil {
		httperror.NotFoundResponse(w, r)
		return
	}

	card, err := c.service.GetCardByID(r.Context(), cid)
	if err != nil {
		httperror.NotFoundResponse(w, r)
		return
	}

	if err = json.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"card": card,
	}, nil); err != nil {
		httperror.ServerErrorResponse(w, r, err)
	}
}

func (c *card) SetupEscrowCard(w http.ResponseWriter, r *http.Request) {
	var req dto.CardInformation

	err := json.ReadJSON(w, r, &req)
	if err != nil {
		httperror.BadRequestResponse(w, r, err)
		return
	}

	if err = c.validator.Struct(req); err != nil {
		httperror.FailedValidationResponse(w, r, err)
		return
	}

	escrow := rcontext.GetEscrowUser(r)

	if err := c.service.SetupEscrowCard(r.Context(), escrow.ID, req); err != nil {
		if errors.Is(err, service.ErrInvalidCardInfo) {
			httperror.BadRequestResponse(w, r, err)
		} else {
			httperror.ServerErrorResponse(w, r, err)
		}
		return
	}

	if err = json.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"message": "setup transfer successfully",
	}, nil); err != nil {
		httperror.ServerErrorResponse(w, r, err)
	}
}

func (c *card) SetupProjectCard(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	pid, err := uuid.Parse(idStr)
	if err != nil {
		httperror.NotFoundResponse(w, r)
		return
	}
	var req dto.CardInformation

	err = json.ReadJSON(w, r, &req)
	if err != nil {
		httperror.BadRequestResponse(w, r, err)
		return
	}

	if err = c.validator.Struct(req); err != nil {
		httperror.FailedValidationResponse(w, r, err)
		return
	}

	if err := c.service.SetupProjectCard(r.Context(), pid, req); err != nil {
		if errors.Is(err, service.ErrInvalidCardInfo) {
			httperror.BadRequestResponse(w, r, err)
		} else {
			httperror.ServerErrorResponse(w, r, err)
		}
		return
	}

	if err = json.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"message": "setup transfer successfully",
	}, nil); err != nil {
		httperror.ServerErrorResponse(w, r, err)
	}
}
