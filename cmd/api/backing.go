package main

import (
	"net/http"

	"github.com/kanowfy/donorbox/internal/convert"
	"github.com/kanowfy/donorbox/internal/db"
	"github.com/kanowfy/donorbox/internal/models"
	"github.com/kanowfy/donorbox/internal/service"
)

func (app *application) getBackingsForProjectHandler(w http.ResponseWriter, r *http.Request) {
	id, err := convert.StringToPgxUUID(r.PathValue("id"))
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	backings, err := app.repository.GetBackingsForProject(r.Context(), id)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	if err := app.writeJSON(w, http.StatusOK, map[string]interface{}{
		"backings": backings,
	}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getBackingsForUserHandler(w http.ResponseWriter, r *http.Request) {
	user := app.contextGetUser(r)

	backings, err := app.repository.GetBackingsForUser(r.Context(), user.ID)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	if err := app.writeJSON(w, http.StatusOK, map[string]interface{}{
		"backings": backings,
	}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) createProjectBackingHandler(w http.ResponseWriter, r *http.Request) {
	pid, err := convert.StringToPgxUUID(r.PathValue("id"))
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	var req models.BackingRequest

	err = app.readJSON(w, r, &req)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err = app.validator.Struct(req); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	user := app.contextGetUser(r)

	arg := db.CreateBackingParams{
		ProjectID: pid,
		BackerID:  user.ID,
		Amount:    convert.MustStringToInt64(req.Amount),
	}

	if err := service.AcceptBacking(r.Context(), app.repository.pool, app.repository.Queries, arg); err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	if err := app.writeJSON(w, http.StatusOK, map[string]interface{}{
		"message": "fund accepted",
	}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
