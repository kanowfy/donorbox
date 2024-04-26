package main

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/kanowfy/donorbox/internal/convert"
	"github.com/kanowfy/donorbox/internal/db"
	"github.com/kanowfy/donorbox/internal/models"
	"github.com/kanowfy/donorbox/internal/service"
)

func (app *application) projectBackingHandler(w http.ResponseWriter, r *http.Request) {
	var req models.BackingRequest

	err := app.readJSON(w, r, &req)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	slog.Debug(fmt.Sprintf("%v", req))

	if err = app.validator.Struct(req); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	user := app.contextGetUser(r)

	arg := db.CreateBackingParams{
		ProjectID: convert.MustStringToPgxUUID(req.ProjectID),
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
