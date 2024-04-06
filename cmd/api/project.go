package main

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/kanowfy/donorbox/internal/db"
	"github.com/kanowfy/donorbox/internal/models"
)

func (app *application) getAllProjectsHandler(w http.ResponseWriter, r *http.Request) {
	projects, err := app.db.GetAllProjects(r.Context())
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	if projects == nil {
		projects = []db.Project{}
	}

	if err = app.writeJSON(w, http.StatusOK, map[string]interface{}{
		"projects": projects,
	}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getOneProjectHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	var id pgtype.UUID
	err := id.Scan(idStr)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	project, err := app.db.GetProjectByID(r.Context(), id)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	if err = app.writeJSON(w, http.StatusOK, map[string]interface{}{
		"project": project,
	}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) createProjectHandler(w http.ResponseWriter, r *http.Request) {
	var req models.CreateProjectRequest

	err := app.readJSON(w, r, &req)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	var uid pgtype.UUID
	err = uid.Scan(req.UserID)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	var goalAmount pgtype.Numeric
	err = goalAmount.Scan(req.GoalAmount)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	var endDate pgtype.Timestamptz
	err = endDate.Scan(req.EndDate)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	args := db.CreateProjectParams{
		UserID:       uid,
		Title:        req.Title,
		Description:  req.Description,
		CoverPicture: req.CoverPicture,
		GoalAmount:   goalAmount,
		Country:      req.Country,
		Province:     req.Province,
		EndDate:      endDate,
	}

	project, err := app.db.CreateProject(r.Context(), args)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	if err = app.writeJSON(w, http.StatusCreated, map[string]interface{}{
		"project": project,
	}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
