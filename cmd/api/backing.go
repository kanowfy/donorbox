package main

import (
	"net/http"

	"github.com/kanowfy/donorbox/internal/convert"
	"github.com/kanowfy/donorbox/internal/models"
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

	if err := app.service.AcceptBacking(r.Context(), pid, user.ID, req); err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	if err := app.writeJSON(w, http.StatusOK, map[string]interface{}{
		"message": "fund accepted",
	}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getProjectBackingStats(w http.ResponseWriter, r *http.Request) {
	pid, err := convert.StringToPgxUUID(r.PathValue("id"))
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	mostBacking, err := app.repository.GetMostBackingDonor(r.Context(), pid)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	firstBacking, err := app.repository.GetFirstBackingDonor(r.Context(), pid)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	recentBacking, err := app.repository.GetMostRecentBackingDonor(r.Context(), pid)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	count, err := app.repository.GetBackingCountForProject(r.Context(), pid)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	if err := app.writeJSON(w, http.StatusOK, map[string]interface{}{
		"most_backing":   mostBacking,
		"first_backing":  firstBacking,
		"recent_backing": recentBacking,
		"backing_count":  int32(count),
	}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}

}
