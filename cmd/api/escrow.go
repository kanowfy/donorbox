package main

import (
	"errors"
	"net/http"
	"time"

	"github.com/kanowfy/donorbox/internal/convert"
	"github.com/kanowfy/donorbox/internal/models"
	"github.com/kanowfy/donorbox/internal/service"
	"github.com/kanowfy/donorbox/internal/token"
	"golang.org/x/crypto/bcrypt"
)

func (app *application) escrowLoginHandler(w http.ResponseWriter, r *http.Request) {
	var req models.LoginRequest

	err := app.readJSON(w, r, &req)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err = app.validator.Struct(req); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	escrow, err := app.repository.GetEscrowUserByEmail(r.Context(), req.Email)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	// validate password
	if err = bcrypt.CompareHashAndPassword([]byte(escrow.HashedPassword), []byte(req.Password)); err != nil {
		app.invalidCredentialsResponse(w, r)
		return
	}

	token, err := token.GenerateToken(convert.PgxUUIDToString(escrow.ID), time.Hour*3*24)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	// return token
	if err := app.writeJSON(w, http.StatusOK, map[string]interface{}{
		"token": token,
	}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getAuthenticatedEscrowHandler(w http.ResponseWriter, r *http.Request) {
	escrow := app.contextGetEscrowUser(r)

	if err := app.writeJSON(w, http.StatusOK, map[string]interface{}{
		"user": escrow,
	}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) setupEscrowCardHandler(w http.ResponseWriter, r *http.Request) {
	var req models.CardInformation

	err := app.readJSON(w, r, &req)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err = app.validator.Struct(req); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	escrow := app.contextGetEscrowUser(r)

	if err := app.service.SetupEscrowCard(r.Context(), escrow.ID, req); err != nil {
		if errors.Is(err, service.ErrInvalidCardInfo) {
			app.badRequestResponse(w, r, err)
		} else {
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	if err = app.writeJSON(w, http.StatusOK, map[string]interface{}{
		"message": "setup transfer successfully",
	}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) payoutHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := convert.StringToPgxUUID(idStr)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	escrow := app.contextGetEscrowUser(r)

	_, err = app.repository.GetProjectByID(r.Context(), id)

	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	if err := app.service.Payout(r.Context(), id, escrow); err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	// send email to owner

	if err = app.writeJSON(w, http.StatusOK, map[string]interface{}{
		"message": "project payout successful",
	}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) refundHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := convert.StringToPgxUUID(idStr)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	escrow := app.contextGetEscrowUser(r)

	_, err = app.repository.GetProjectByID(r.Context(), id)

	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	if err := app.service.Refund(r.Context(), id, escrow); err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	// send email to owner

	if err = app.writeJSON(w, http.StatusOK, map[string]interface{}{
		"message": "project refund successful",
	}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getStatisticsHandler(w http.ResponseWriter, r *http.Request) {
	stats, err := app.repository.GetStatistics(r.Context())
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	transactions, err := app.repository.GetTransactionsStatsByWeek(r.Context())
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	if err = app.writeJSON(w, http.StatusOK, map[string]interface{}{
		"stats":        stats,
		"transactions": transactions,
	}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
