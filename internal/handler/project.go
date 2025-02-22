package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/kanowfy/donorbox/internal/dto"
	"github.com/kanowfy/donorbox/internal/rcontext"
	"github.com/kanowfy/donorbox/internal/service"
	"github.com/kanowfy/donorbox/pkg/helper"
	"github.com/kanowfy/donorbox/pkg/httperror"
	"github.com/kanowfy/donorbox/pkg/json"
)

type Project interface {
	GetAllProjects(w http.ResponseWriter, r *http.Request)
	SearchProjects(w http.ResponseWriter, r *http.Request)
	GetProjectsForUser(w http.ResponseWriter, r *http.Request)
	GetEndedProjects(w http.ResponseWriter, r *http.Request)
	GetPendingProjects(w http.ResponseWriter, r *http.Request)
	GetProjectDetails(w http.ResponseWriter, r *http.Request)
	CreateProject(w http.ResponseWriter, r *http.Request)
	UpdateProject(w http.ResponseWriter, r *http.Request)
	DeleteProject(w http.ResponseWriter, r *http.Request)
	GetAllCategories(w http.ResponseWriter, r *http.Request)
	GetCategoryByName(w http.ResponseWriter, r *http.Request)
	GetFundedMilestones(w http.ResponseWriter, r *http.Request)
	CreateMilestoneProof(w http.ResponseWriter, r *http.Request)
	CreateProjectReport(w http.ResponseWriter, r *http.Request)
	GetProjectReports(w http.ResponseWriter, r *http.Request)
	GetDisputedProjects(w http.ResponseWriter, r *http.Request)
}

type project struct {
	service   service.Project
	validator *validator.Validate
}

func NewProject(service service.Project, validator *validator.Validate) Project {
	return &project{
		service,
		validator,
	}
}

func (p *project) GetAllProjects(w http.ResponseWriter, r *http.Request) {
	qs := r.URL.Query()

	category, _ := helper.ReadInt(qs, "category", 0)

	projects, err := p.service.GetAllProjects(r.Context(), category)
	if err != nil {
		httperror.ServerErrorResponse(w, r, err)
		return
	}

	if err = json.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"projects": projects,
	}, nil); err != nil {
		httperror.ServerErrorResponse(w, r, err)
	}
}

func (p *project) SearchProjects(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Query string `json:"query" validate:"required"`
	}

	err := json.ReadJSON(w, r, &req)
	if err != nil {
		httperror.BadRequestResponse(w, r, err)
		return
	}

	if err = p.validator.Struct(req); err != nil {
		httperror.FailedValidationResponse(w, r, err)
		return
	}

	projects, err := p.service.SearchProjects(r.Context(), req.Query)
	if err != nil {
		httperror.ServerErrorResponse(w, r, err)
		return
	}

	if err = json.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"projects": projects,
	}, nil); err != nil {
		httperror.ServerErrorResponse(w, r, err)
	}
}

func (p *project) GetProjectsForUser(w http.ResponseWriter, r *http.Request) {
	user := rcontext.GetUser(r)

	projects, err := p.service.GetProjectsForUser(r.Context(), user.ID)
	if err != nil {
		httperror.ServerErrorResponse(w, r, err)
		return
	}

	if err = json.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"projects": projects,
	}, nil); err != nil {
		httperror.ServerErrorResponse(w, r, err)
	}
}

func (p *project) GetEndedProjects(w http.ResponseWriter, r *http.Request) {
	projects, err := p.service.GetEndedProjects(r.Context())
	if err != nil {
		httperror.ServerErrorResponse(w, r, err)
		return
	}

	if err = json.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"projects": projects,
	}, nil); err != nil {
		httperror.ServerErrorResponse(w, r, err)
	}
}

func (p *project) GetPendingProjects(w http.ResponseWriter, r *http.Request) {
	projects, err := p.service.GetPendingProjects(r.Context())
	if err != nil {
		httperror.ServerErrorResponse(w, r, err)
		return
	}

	if err = json.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"result": projects,
	}, nil); err != nil {
		httperror.ServerErrorResponse(w, r, err)
	}
}

func (p *project) GetProjectDetails(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		httperror.NotFoundResponse(w, r)
		return
	}

	project, milestones, backings, user, err := p.service.GetProjectDetails(r.Context(), id)
	if err != nil {
		httperror.NotFoundResponse(w, r)
		return
	}

	if err = json.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"project":    project,
		"milestones": milestones,
		"backings":   backings,
		"user":       user,
	}, nil); err != nil {
		httperror.ServerErrorResponse(w, r, err)
	}
}

func (p *project) CreateProject(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateProjectRequest

	err := json.ReadJSON(w, r, &req)
	if err != nil {
		httperror.BadRequestResponse(w, r, err)
		return
	}

	if err = p.validator.Struct(req); err != nil {
		httperror.FailedValidationResponse(w, r, err)
		return
	}

	user := rcontext.GetUser(r)

	project, err := p.service.CreateProject(r.Context(), user.ID, req)
	if err != nil {
		httperror.ServerErrorResponse(w, r, err)
	}

	if err = json.WriteJSON(w, http.StatusCreated, map[string]interface{}{
		"result": project,
	}, nil); err != nil {
		httperror.ServerErrorResponse(w, r, err)
	}
}

func (p *project) UpdateProject(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		httperror.NotFoundResponse(w, r)
		return
	}

	var req dto.UpdateProjectRequest

	err = json.ReadJSON(w, r, &req)
	if err != nil {
		httperror.BadRequestResponse(w, r, err)
		return
	}

	if err = p.validator.Struct(req); err != nil {
		httperror.FailedValidationResponse(w, r, err)
		return
	}

	user := rcontext.GetUser(r)

	err = p.service.UpdateProject(r.Context(), user.ID, id, req)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrProjectNotFound):
			httperror.NotFoundResponse(w, r)
		case errors.Is(err, service.ErrNotOwner):
			httperror.AuthenticationRequiredResponse(w, r)
		}
		return
	}

	if err = json.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"message": "project has successfully been updated",
	}, nil); err != nil {
		httperror.ServerErrorResponse(w, r, err)
	}
}

func (p *project) DeleteProject(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		httperror.NotFoundResponse(w, r)
		return
	}

	user := rcontext.GetUser(r)

	if err = p.service.DeleteProject(r.Context(), user.ID, id); err != nil {
		httperror.ServerErrorResponse(w, r, err)
		return
	}

	if err = json.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"message": "project has successfully been deleted",
	}, nil); err != nil {
		httperror.ServerErrorResponse(w, r, err)
	}

}

func (p *project) GetAllCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := p.service.GetAllCategories(r.Context())
	if err != nil {
		httperror.ServerErrorResponse(w, r, err)
		return
	}

	if err = json.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"categories": categories,
	}, nil); err != nil {
		httperror.ServerErrorResponse(w, r, err)
	}
}

func (p *project) GetCategoryByName(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	if name == "" {
		httperror.NotFoundResponse(w, r)
		return
	}

	category, err := p.service.GetCategoryByName(r.Context(), name)
	if err != nil {
		httperror.NotFoundResponse(w, r)
	}

	if err = json.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"category": category,
	}, nil); err != nil {
		httperror.ServerErrorResponse(w, r, err)
	}
}

func (p *project) GetFundedMilestones(w http.ResponseWriter, r *http.Request) {
	milestones, err := p.service.GetFundedMilestones(r.Context())
	if err != nil {
		httperror.ServerErrorResponse(w, r, err)
		return
	}

	if err := json.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"milestones": milestones,
	}, nil); err != nil {
		httperror.ServerErrorResponse(w, r, err)
	}
}

func (p *project) CreateMilestoneProof(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateMilestoneProofRequest

	err := json.ReadJSON(w, r, &req)
	if err != nil {
		httperror.BadRequestResponse(w, r, err)
		return
	}

	if err = p.validator.Struct(req); err != nil {
		httperror.FailedValidationResponse(w, r, err)
		return
	}

	user := rcontext.GetUser(r)

	if err := p.service.CreateMilestoneProof(r.Context(), user.ID, req); err != nil {
		httperror.ServerErrorResponse(w, r, err)
	}

	if err := json.WriteJSON(w, http.StatusCreated, map[string]interface{}{
		"message": "proof created",
	}, nil); err != nil {
		httperror.ServerErrorResponse(w, r, err)
	}
}

func (p *project) CreateProjectReport(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		httperror.NotFoundResponse(w, r)
		return
	}

	var req dto.CreateProjectReportRequest

	if err := json.ReadJSON(w, r, &req); err != nil {
		httperror.BadRequestResponse(w, r, err)
		return
	}

	if err = p.validator.Struct(req); err != nil {
		httperror.FailedValidationResponse(w, r, err)
		return
	}

	if err := p.service.CreateProjectReport(r.Context(), id, req); err != nil {
		httperror.ServerErrorResponse(w, r, err)
		return
	}

	if err := json.WriteJSON(w, http.StatusCreated, map[string]interface{}{
		"message": "report created",
	}, nil); err != nil {
		httperror.ServerErrorResponse(w, r, err)
	}
}

func (p *project) GetProjectReports(w http.ResponseWriter, r *http.Request) {
	reports, err := p.service.GetProjectReports(r.Context())
	if err != nil {
		httperror.ServerErrorResponse(w, r, err)
		return
	}

	if err := json.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"reports": reports,
	}, nil); err != nil {
		httperror.ServerErrorResponse(w, r, err)
	}
}

func (p *project) GetDisputedProjects(w http.ResponseWriter, r *http.Request) {
	projects, err := p.service.GetDisputedProjects(r.Context())
	if err != nil {
		httperror.ServerErrorResponse(w, r, err)
		return
	}

	if err := json.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"projects": projects,
	}, nil); err != nil {
		httperror.ServerErrorResponse(w, r, err)
	}

}
