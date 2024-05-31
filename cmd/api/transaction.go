package main

import (
	"net/http"

	"github.com/kanowfy/donorbox/internal/convert"
)

func (app *application) getAllTransactionsHandler(w http.ResponseWriter, r *http.Request) {
	transactions, err := app.repository.GetAllTransactions(r.Context())
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	if err = app.writeJSON(w, http.StatusOK, map[string]interface{}{
		"transactions": transactions,
	}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getOneTransactionHandler(w http.ResponseWriter, r *http.Request) {
	id, err := convert.StringToPgxUUID(r.PathValue("id"))
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	transaction, err := app.repository.GetTransactionByID(r.Context(), id)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	if err = app.writeJSON(w, http.StatusOK, map[string]interface{}{
		"transaction": transaction,
	}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getTransactionAuditHandler(w http.ResponseWriter, r *http.Request) {
	backingID, err := convert.StringToPgxUUID(r.PathValue("id"))
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	transactions, err := app.repository.GetTransactionsForProject(r.Context(), backingID)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	if err = app.writeJSON(w, http.StatusOK, map[string]interface{}{
		"transactions": transactions,
	}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
