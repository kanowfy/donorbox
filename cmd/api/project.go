package main

import (
	"net/http"

	"github.com/kanowfy/donorbox/internal/convert"
	"github.com/kanowfy/donorbox/internal/db"
	"github.com/kanowfy/donorbox/internal/models"
)

func (app *application) getAllProjectsHandler(w http.ResponseWriter, r *http.Request) {
	qs := r.URL.Query()

	page, _ := readInt(qs, "page", 1)
	pageSize, _ := readInt(qs, "page_size", 6)
	category, _ := readInt(qs, "category", 0)

	filters := models.Filters{
		Category: category,
		Page:     page,
		PageSize: pageSize,
	}

	var args db.GetAllProjectsParams

	args.Category = int32(category)
	args.PageLimit = int32(filters.Limit())
	args.TotalOffset = int32(filters.Offset())

	projects, err := app.repository.GetAllProjects(r.Context(), args)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	metadata := models.CalculateMetadata(len(projects), filters.Page, filters.PageSize)

	if projects == nil {
		projects = []db.GetAllProjectsRow{}
	}

	if err = app.writeJSON(w, http.StatusOK, map[string]interface{}{
		"projects": projects,
		"metadata": metadata,
	}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) searchProjectsHandler(w http.ResponseWriter, r *http.Request) {
	qs := r.URL.Query()

	page, _ := readInt(qs, "page", 1)
	pageSize, _ := readInt(qs, "page_size", 6)
	category, _ := readInt(qs, "category", 0)
	searchQuery := readString(qs, "query", "")
	province := readString(qs, "province", "")
	country := readString(qs, "country", "")
	closeToGoal, _ := readInt(qs, "close_to_goal", 0)
	recent, _ := readInt(qs, "recently_launched", 0)

	filters := models.Filters{
		Category: category,
		Page:     page,
		PageSize: pageSize,
	}

	args := db.SearchProjectsParams{
		Category:    int32(category),
		SearchQuery: searchQuery,
		Province:    province,
		Country:     country,
		CloseToGoal: int32(closeToGoal),
		Recent:      int32(recent),
		PageLimit:   int32(filters.Limit()),
		TotalOffset: int32(filters.Offset()),
	}

	projects, err := app.repository.SearchProjects(r.Context(), args)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	metadata := models.CalculateMetadata(len(projects), filters.Page, filters.PageSize)

	if projects == nil {
		projects = []db.SearchProjectsRow{}
	}

	if err = app.writeJSON(w, http.StatusOK, map[string]interface{}{
		"projects": projects,
		"metadata": metadata,
	}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getOneProjectHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := convert.StringToPgxUUID(idStr)
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
		UserID:       convert.MustStringToPgxUUID(req.UserID),
		CategoryID:   int32(req.CategoryID),
		Title:        req.Title,
		Description:  req.Description,
		CoverPicture: req.CoverPicture,
		GoalAmount:   convert.MustStringToInt64(req.GoalAmount),
		Country:      req.Country,
		Province:     req.Province,
		EndDate:      convert.MustTimeToPgxTimestamp(req.EndDate),
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
	id, err := convert.StringToPgxUUID(idStr)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	project, err := app.repository.GetProjectByID(r.Context(), id)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	//TODO: check permission of requesting user and whether the current amount is 0

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
		updateParams.GoalAmount = convert.MustStringToInt64(*payload.GoalAmount)
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
		updateParams.EndDate = convert.MustTimeToPgxTimestamp(*payload.EndDate)
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
	id, err := convert.StringToPgxUUID(idStr)
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

func (app *application) getAllCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	categories, err := app.repository.GetAllCategories(r.Context())
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	if err = app.writeJSON(w, http.StatusOK, map[string]interface{}{
		"categories": categories,
	}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) createProjectUpdateHandler(w http.ResponseWriter, r *http.Request) {
	var req models.CreateProjectUpdateRequest

	err := app.readJSON(w, r, &req)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err = app.validator.Struct(req); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	pid := convert.MustStringToPgxUUID(req.ProjectID)

	// check if projectID is valid
	project, err := app.repository.GetProjectByID(r.Context(), pid)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	// Check permission of requesting user
	user := app.contextGetUser(r)
	if user.ID != project.UserID {
		app.authenticationRequiredResponse(w, r)
		return
	}

	args := db.CreateProjectUpdateParams{
		ProjectID:   pid,
		Description: req.Description,
	}

	update, err := app.repository.CreateProjectUpdate(r.Context(), args)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	if err := app.writeJSON(w, http.StatusCreated, map[string]interface{}{
		"project_update": update,
	}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) createProjectCommentHandler(w http.ResponseWriter, r *http.Request) {
	var req models.CreateProjectCommentRequest

	err := app.readJSON(w, r, &req)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err = app.validator.Struct(req); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	pid := convert.MustStringToPgxUUID(req.ProjectID)

	// check if projectID is valid, backerID is validated through middleware
	if _, err = app.repository.GetProjectByID(r.Context(), pid); err != nil {
		app.notFoundResponse(w, r)
		return
	}

	//TODO: check permission of requesting user(must be owner or backer of the project)

	args := db.CreateProjectCommentParams{
		ProjectID: pid,
		BackerID:  convert.MustStringToPgxUUID(req.BackerID),
		Content:   req.Content,
	}

	if req.ParentCommentID != nil {
		args.ParentCommentID = convert.MustStringToPgxUUID(*req.ParentCommentID)
	}

	comment, err := app.repository.CreateProjectComment(r.Context(), args)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	if err := app.writeJSON(w, http.StatusCreated, map[string]interface{}{
		"project_comment": comment,
	}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}

}
