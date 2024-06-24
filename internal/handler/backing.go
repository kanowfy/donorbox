package handler

import (
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/kanowfy/donorbox/internal/dto"
	"github.com/kanowfy/donorbox/internal/service"
	"github.com/kanowfy/donorbox/pkg/httperror"
	"github.com/kanowfy/donorbox/pkg/json"
)

type Backing interface {
	GetBackingsForProject(w http.ResponseWriter, r *http.Request)
	CreateProjectBacking(w http.ResponseWriter, r *http.Request)
	GetProjectBackingStats(w http.ResponseWriter, r *http.Request)
}

type backing struct {
	service   service.Backing
	validator *validator.Validate
}

func NewBacking(service service.Backing, validator *validator.Validate) Backing {
	return &backing{
		service,
		validator,
	}
}

func (b *backing) GetBackingsForProject(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		httperror.NotFoundResponse(w, r)
		return
	}

	backings, err := b.service.GetBackingsForProject(r.Context(), id)
	if err != nil {
		if errors.Is(err, service.ErrProjectNotFound) {
			httperror.NotFoundResponse(w, r)
		} else {
			httperror.ServerErrorResponse(w, r, err)
		}
		return
	}

	if err := json.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"backings": backings,
	}, nil); err != nil {
		httperror.ServerErrorResponse(w, r, err)
	}
}

func (b *backing) CreateProjectBacking(w http.ResponseWriter, r *http.Request) {
	pid, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		httperror.NotFoundResponse(w, r)
		return
	}

	var req dto.BackingRequest

	err = json.ReadJSON(w, r, &req)
	if err != nil {
		httperror.BadRequestResponse(w, r, err)
		return
	}

	if err = b.validator.Struct(req); err != nil {
		httperror.FailedValidationResponse(w, r, err)
		return
	}

	if err := b.service.AcceptBacking(r.Context(), pid, req); err != nil {
		httperror.ServerErrorResponse(w, r, err)
		return
	}

	if err := json.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"message": "fund accepted",
	}, nil); err != nil {
		httperror.ServerErrorResponse(w, r, err)
	}
}

func (b *backing) GetProjectBackingStats(w http.ResponseWriter, r *http.Request) {
	pid, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		httperror.NotFoundResponse(w, r)
		return
	}

	backingAggregation, err := b.service.GetProjectBackingAggregation(r.Context(), pid)
	if err != nil {
		httperror.ServerErrorResponse(w, r, err)
		return
	}

	if err := json.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"most_backing":   backingAggregation.MostAmountBacking,
		"first_backing":  backingAggregation.FirstBacking,
		"recent_backing": backingAggregation.RecentBacking,
		"backing_count":  backingAggregation.TotalBacking,
	}, nil); err != nil {
		httperror.ServerErrorResponse(w, r, err)
	}

}
