package main

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/kanowfy/donorbox/internal/db"
	"github.com/kanowfy/donorbox/internal/models"
)

func (app *application) getAllProjectsHandler(w http.ResponseWriter, r *http.Request) {
	qs := r.URL.Query()

	page, err := readInt(qs, "page", 1)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	pageSize, err := readInt(qs, "page_size", 20)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	sort := readString(qs, "sort", "-end_date")

	filters := models.Filters{
		Page:         page,
		PageSize:     pageSize,
		Sort:         sort,
		SortSafeList: []string{"end_date", "current_amount", "-end_date", "-current_amount"},
	}

	var args db.GetAllProjectsParams
	switch sort {
	case "end_date":
		args.EndDateAsc = 1
	case "current_amount":
		args.CurrentAmountAsc = 1
	case "-current_amount":
		args.CurrentAmountDesc = 1
	default:
		args.EndDateDesc = 1
	}

	args.PageLimit = int32(filters.Limit())
	args.TotalOffset = int32(filters.Offset())

	projects, err := app.repository.GetAllProjects(r.Context(), args)
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
	id, err := stringToPgxUUID(idStr)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	project, err := app.repository.GetProjectByID(r.Context(), id)
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

	if err = app.validator.Struct(req); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	args := db.CreateProjectParams{
		UserID:       mustStringToPgxUUID(req.UserID),
		Title:        req.Title,
		Description:  req.Description,
		CoverPicture: req.CoverPicture,
		GoalAmount:   mustStringToPgxNumeric(req.GoalAmount),
		Country:      req.Country,
		Province:     req.Province,
		EndDate:      mustTimeToPgxTimestamp(req.EndDate),
	}

	project, err := app.repository.CreateProject(r.Context(), args)
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

func (app *application) updateProjectHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := stringToPgxUUID(idStr)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	project, err := app.repository.GetProjectByID(r.Context(), id)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	var payload models.UpdateProjectRequest

	err = app.readJSON(w, r, &payload)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err = app.validator.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	var updateParams db.UpdateProjectByIDParams
	updateParams.ID = project.ID

	if payload.Title != nil {
		updateParams.Title = *payload.Title
	} else {
		updateParams.Title = project.Title
	}

	if payload.Description != nil {
		updateParams.Description = *payload.Description
	} else {
		updateParams.Description = project.Description
	}

	if payload.CoverPicture != nil {
		updateParams.CoverPicture = *payload.CoverPicture
	} else {
		updateParams.CoverPicture = project.CoverPicture
	}

	if payload.GoalAmount != nil {
		updateParams.GoalAmount = mustStringToPgxNumeric(*payload.GoalAmount)
	} else {
		updateParams.GoalAmount = project.GoalAmount
	}

	if payload.Country != nil {
		updateParams.Country = *payload.Country
	} else {
		updateParams.Country = project.Country
	}

	if payload.Province != nil {
		updateParams.Province = *payload.Province
	} else {
		updateParams.Province = project.Province
	}

	if payload.EndDate != nil {
		updateParams.EndDate = mustTimeToPgxTimestamp(*payload.EndDate)
	} else {
		updateParams.EndDate = project.EndDate
	}

	if err = app.repository.UpdateProjectByID(r.Context(), updateParams); err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	if err = app.writeJSON(w, http.StatusOK, map[string]interface{}{
		"message": "project has successfully been updated",
	}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deleteProjectHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	var id pgtype.UUID
	err := id.Scan(idStr)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	if err = app.repository.DeleteProjectByID(r.Context(), id); err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	if err = app.writeJSON(w, http.StatusOK, map[string]interface{}{
		"message": "project has successfully been deleted",
	}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}

}
