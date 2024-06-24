package handler

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/kanowfy/donorbox/internal/service"
	"github.com/kanowfy/donorbox/pkg/httperror"
	"github.com/kanowfy/donorbox/pkg/json"
)

type Transaction interface {
	GetAllTransactions(w http.ResponseWriter, r *http.Request)
	GetOneTransaction(w http.ResponseWriter, r *http.Request)
	GetTransactionAudit(w http.ResponseWriter, r *http.Request)
}

type transaction struct {
	service   service.Transaction
	validator *validator.Validate
}

func NewTransaction(service service.Transaction, validator *validator.Validate) Transaction {
	return &transaction{
		service,
		validator,
	}
}

func (t *transaction) GetAllTransactions(w http.ResponseWriter, r *http.Request) {
	transactions, err := t.service.GetAllTransactions(r.Context())
	if err != nil {
		httperror.ServerErrorResponse(w, r, err)
		return
	}

	if err = json.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"transactions": transactions,
	}, nil); err != nil {
		httperror.ServerErrorResponse(w, r, err)
	}
}

func (t *transaction) GetOneTransaction(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		httperror.NotFoundResponse(w, r)
		return
	}

	transaction, err := t.service.GetOneTransaction(r.Context(), id)
	if err != nil {
		httperror.NotFoundResponse(w, r)
		return
	}

	if err = json.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"transaction": transaction,
	}, nil); err != nil {
		httperror.ServerErrorResponse(w, r, err)
	}
}

func (t *transaction) GetTransactionAudit(w http.ResponseWriter, r *http.Request) {
	pid, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		httperror.NotFoundResponse(w, r)
		return
	}

	transactions, err := t.service.GetTransactionAudit(r.Context(), pid)
	if err != nil {
		httperror.ServerErrorResponse(w, r, err)
		return
	}

	if err = json.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"transactions": transactions,
	}, nil); err != nil {
		httperror.ServerErrorResponse(w, r, err)
	}
}
