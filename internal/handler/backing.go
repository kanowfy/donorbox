package handler

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/kanowfy/donorbox/internal/dto"
	"github.com/kanowfy/donorbox/internal/service"
	"github.com/kanowfy/donorbox/pkg/httperror"
	"github.com/kanowfy/donorbox/pkg/json"
	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/paymentintent"
)

type Backing interface {
	CreatePaymentIntent(w http.ResponseWriter, r *http.Request)
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
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
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

func (b *backing) CreatePaymentIntent(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Amount int64 `json:"amount"`
	}

	if err := json.ReadJSON(w, r, &req); err != nil {
		httperror.BadRequestResponse(w, r, err)
		return
	}

	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(req.Amount),
		Currency: stripe.String(string(stripe.CurrencyVND)),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
	}

	pi, err := paymentintent.New(params)
	if err != nil {
		httperror.ServerErrorResponse(w, r, err)
		return
	}

	if err := json.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"client_secret": pi.ClientSecret,
	}, nil); err != nil {
		httperror.ServerErrorResponse(w, r, err)
	}
}

func (b *backing) CreateProjectBacking(w http.ResponseWriter, r *http.Request) {
	pid, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
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

	log.Printf("%+v\n", req)

	if err := b.service.CreateBacking(r.Context(), pid, req); err != nil {
		httperror.ServerErrorResponse(w, r, err)
		return
	}

	if err := json.WriteJSON(w, http.StatusCreated, map[string]interface{}{
		"message": "backing created",
	}, nil); err != nil {
		httperror.ServerErrorResponse(w, r, err)
	}
}

func (b *backing) GetProjectBackingStats(w http.ResponseWriter, r *http.Request) {
	pid, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		httperror.NotFoundResponse(w, r)
		return
	}

	mostBacking, firstBacking, recentBacking, backingCount, err := b.service.GetProjectBackingStats(r.Context(), pid)
	if err != nil {
		httperror.ServerErrorResponse(w, r, err)
		return
	}

	if err := json.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"most_backing":   mostBacking,
		"first_backing":  firstBacking,
		"recent_backing": recentBacking,
		"backing_count":  backingCount,
	}, nil); err != nil {
		httperror.ServerErrorResponse(w, r, err)
	}

}
