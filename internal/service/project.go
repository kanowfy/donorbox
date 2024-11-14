package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/kanowfy/donorbox/internal/db"
	"github.com/kanowfy/donorbox/internal/dto"
	"github.com/kanowfy/donorbox/internal/filters"
	"github.com/kanowfy/donorbox/internal/model"
)

var (
	ErrProjectNotFound = errors.New("project not found")
	ErrNotOwner        = errors.New("user does not own the project")
)

type Project interface {
	GetAllProjects(ctx context.Context, pageNum, pageSize, categoryIndex int) ([]model.Project, filters.Metadata, error)
	SearchProjects(ctx context.Context, query string, pageNum, pageSize int) ([]model.Project, filters.Metadata, error)
	GetProjectsForUser(ctx context.Context, userID uuid.UUID) ([]model.Project, error)
	GetEndedProjects(ctx context.Context) ([]model.Project, error)
	GetProjectDetails(ctx context.Context, projectID uuid.UUID) (*model.Project, []model.Milestone, []model.Backing, []model.ProjectUpdate, *model.User, error)
	CreateProject(ctx context.Context, userID uuid.UUID, req dto.CreateProjectRequest) (*dto.CreateProjectResponse, error)
	UpdateProject(ctx context.Context, userID, projectID uuid.UUID, req dto.UpdateProjectRequest) error
	DeleteProject(ctx context.Context, userID, projectID uuid.UUID) error
	GetAllCategories(ctx context.Context) ([]model.Category, error)
	GetProjectUpdates(ctx context.Context, projectID uuid.UUID) ([]model.ProjectUpdate, error)
	CreateProjectUpdate(ctx context.Context, userID uuid.UUID, req dto.CreateProjectUpdateRequest) (*model.ProjectUpdate, error)
	GetUnresolvedMilestones(ctx context.Context) ([]model.Milestone, error)
	//CheckAndUpdateFinishedProjects(ctx context.Context) error
}

type project struct {
	repository     db.Querier
	backingService Backing
	userService    User
}

func NewProject(repository db.Querier, backingService Backing, userService User) Project {
	return &project{
		repository,
		backingService,
		userService,
	}
}

func (p *project) GetAllProjects(ctx context.Context, pageNum, pageSize, categoryIndex int) ([]model.Project, filters.Metadata, error) {
	f := filters.Filters{
		Category: categoryIndex,
		Page:     pageNum,
		PageSize: pageSize,
	}

	var args db.GetAllProjectsParams

	args.Category = int32(categoryIndex)
	args.PageLimit = int32(f.Limit())
	args.TotalOffset = int32(f.Offset())

	dbProjects, err := p.repository.GetAllProjects(ctx, args)
	if err != nil {
		return nil, filters.Metadata{}, err
	}

	metadata := filters.CalculateMetadata(len(dbProjects), f.Page, f.PageSize)

	var projects []model.Project

	for _, p := range dbProjects {
		projects = append(projects, model.Project{
			ID:             p.ID,
			UserID:         p.UserID,
			CategoryID:     p.CategoryID,
			Title:          p.Title,
			Description:    p.Description,
			TotalFund:      p.TotalFund,
			CoverPicture:   p.CoverPicture,
			ReceiverName:   p.ReceiverName,
			ReceiverNumber: p.ReceiverNumber,
			Address:        p.Address,
			District:       p.District,
			City:           p.City,
			Country:        p.Country,
			StartDate:      p.StartDate,
			EndDate:        p.EndDate,
			BackingCount:   &p.BackingCount,
		})
	}

	return projects, metadata, nil
}

func (p *project) SearchProjects(ctx context.Context, query string, pageNum, pageSize int) ([]model.Project, filters.Metadata, error) {
	f := filters.Filters{
		Page:     pageNum,
		PageSize: pageSize,
	}

	args := db.SearchProjectsParams{
		SearchQuery: query,
		PageLimit:   int32(f.Limit()),
		TotalOffset: int32(f.Offset()),
	}

	dbProjects, err := p.repository.SearchProjects(ctx, args)
	if err != nil {
		return nil, filters.Metadata{}, err
	}

	metadata := filters.CalculateMetadata(len(dbProjects), f.Page, f.PageSize)

	var projects []model.Project

	for _, p := range dbProjects {
		projects = append(projects, model.Project{
			ID:             p.ID,
			UserID:         p.UserID,
			CategoryID:     p.CategoryID,
			Title:          p.Title,
			Description:    p.Description,
			TotalFund:      p.TotalFund,
			CoverPicture:   p.CoverPicture,
			ReceiverName:   p.ReceiverName,
			ReceiverNumber: p.ReceiverNumber,
			Address:        p.Address,
			District:       p.District,
			City:           p.City,
			Country:        p.Country,
			StartDate:      p.StartDate,
			EndDate:        p.EndDate,
			BackingCount:   &p.BackingCount,
		})
	}

	return projects, metadata, nil
}

func (p *project) GetProjectsForUser(ctx context.Context, userID uuid.UUID) ([]model.Project, error) {
	dbProjects, err := p.repository.GetProjectsForUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	var projects []model.Project

	for _, p := range dbProjects {
		projects = append(projects, model.Project{
			ID:             p.ID,
			UserID:         p.UserID,
			CategoryID:     p.CategoryID,
			Title:          p.Title,
			Description:    p.Description,
			TotalFund:      p.TotalFund,
			CoverPicture:   p.CoverPicture,
			ReceiverName:   p.ReceiverName,
			ReceiverNumber: p.ReceiverNumber,
			Address:        p.Address,
			District:       p.District,
			City:           p.City,
			Country:        p.Country,
			StartDate:      p.StartDate,
			EndDate:        p.EndDate,
		})
	}

	return projects, nil
}

func (p *project) GetEndedProjects(ctx context.Context) ([]model.Project, error) {
	dbProjects, err := p.repository.GetFinishedProjects(ctx)
	if err != nil {
		return nil, err
	}

	var projects []model.Project

	for _, p := range dbProjects {
		projects = append(projects, model.Project{
			ID:             p.ID,
			UserID:         p.UserID,
			CategoryID:     p.CategoryID,
			Title:          p.Title,
			Description:    p.Description,
			TotalFund:      p.TotalFund,
			CoverPicture:   p.CoverPicture,
			ReceiverName:   p.ReceiverName,
			ReceiverNumber: p.ReceiverNumber,
			Address:        p.Address,
			District:       p.District,
			City:           p.City,
			Country:        p.Country,
			StartDate:      p.StartDate,
			EndDate:        p.EndDate,
		})
	}

	return projects, nil
}

func (p *project) GetMilestonesForProject(ctx context.Context, projectID uuid.UUID) ([]model.Milestone, error) {
	project, err := p.repository.GetProjectByID(ctx, projectID)
	if err != nil {
		return nil, ErrProjectNotFound
	}

	dbMilestones, err := p.repository.GetMilestoneForProject(ctx, project.ID)
	if err != nil {
		return nil, err
	}

	var milestones []model.Milestone
	for _, m := range dbMilestones {
		milestones = append(milestones, model.Milestone{
			ID:              m.ID,
			ProjectID:       m.ProjectID,
			Title:           m.Title,
			Description:     m.Description,
			FundGoal:        m.FundGoal,
			CurrentFund:     m.CurrentFund,
			BankDescription: m.BankDescription,
			Completed:       m.Completed,
		})
	}

	return milestones, nil
}

// TODO: refactor into one struct
func (p *project) GetProjectDetails(ctx context.Context, projectID uuid.UUID) (*model.Project, []model.Milestone, []model.Backing, []model.ProjectUpdate, *model.User, error) {
	project, err := p.repository.GetProjectByID(ctx, projectID)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	milestones, err := p.GetMilestonesForProject(ctx, projectID)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	backings, err := p.backingService.GetBackingsForProject(ctx, project.ID)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	updates, err := p.GetProjectUpdates(ctx, project.ID)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	user, err := p.userService.GetUserByID(ctx, project.UserID)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	return &model.Project{
		ID:             project.ID,
		UserID:         project.UserID,
		CategoryID:     project.CategoryID,
		Title:          project.Title,
		Description:    project.Description,
		TotalFund:      project.TotalFund,
		CoverPicture:   project.CoverPicture,
		ReceiverName:   project.ReceiverName,
		ReceiverNumber: project.ReceiverNumber,
		Address:        project.Address,
		District:       project.District,
		City:           project.City,
		Country:        project.Country,
		StartDate:      project.StartDate,
		EndDate:        project.EndDate,
	}, milestones, backings, updates, user, nil
}

// TODO: put the queries in transaction
func (p *project) CreateProject(ctx context.Context, userID uuid.UUID, req dto.CreateProjectRequest) (*dto.CreateProjectResponse, error) {
	projectArgs := db.CreateProjectParams{
		UserID:         userID,
		CategoryID:     int32(req.CategoryID),
		Title:          req.Title,
		Description:    req.Description,
		CoverPicture:   req.CoverPicture,
		ReceiverNumber: req.ReceiverNumber,
		ReceiverName:   req.ReceiverName,
		Address:        req.Address,
		District:       req.District,
		City:           req.City,
		Country:        req.Country,
		EndDate:        req.EndDate,
	}

	project, err := p.repository.CreateProject(ctx, projectArgs)
	if err != nil {
		return nil, err
	}

	var milestones []model.Milestone

	for _, m := range req.Milestones {
		milestone, err := p.repository.CreateMilestone(ctx, db.CreateMilestoneParams{
			ProjectID:       project.ID,
			Title:           m.Title,
			Description:     m.Description,
			BankDescription: m.BankDescription,
		})

		if err != nil {
			return nil, err
		}

		milestones = append(milestones, model.Milestone{
			ID:              milestone.ID,
			ProjectID:       milestone.ProjectID,
			Title:           milestone.Title,
			Description:     milestone.Description,
			FundGoal:        milestone.FundGoal,
			BankDescription: milestone.BankDescription,
		})

	}

	return &dto.CreateProjectResponse{
		Project: model.Project{
			ID:             project.ID,
			UserID:         project.UserID,
			CategoryID:     project.CategoryID,
			Title:          project.Title,
			Description:    project.Description,
			TotalFund:      project.TotalFund,
			CoverPicture:   project.CoverPicture,
			ReceiverNumber: project.ReceiverNumber,
			ReceiverName:   project.ReceiverName,
			Address:        project.Address,
			District:       project.District,
			City:           project.City,
			Country:        project.Country,
			StartDate:      project.StartDate,
			EndDate:        project.EndDate,
			Status:         model.ProjectStatusPending,
		},
		Milestones: milestones,
	}, nil
}

func (p *project) UpdateProject(ctx context.Context, userID, projectID uuid.UUID, req dto.UpdateProjectRequest) error {
	project, err := p.repository.GetProjectByID(ctx, projectID)
	if err != nil {
		return ErrProjectNotFound
	}

	// check permission of requesting user and whether the current amount is 0

	if userID != project.UserID {
		return ErrNotOwner
	}

	var updateParams db.UpdateProjectByIDParams
	updateParams.ID = project.ID

	if req.Title != nil {
		updateParams.Title = *req.Title
	} else {
		updateParams.Title = project.Title
	}

	if req.Description != nil {
		updateParams.Description = *req.Description
	} else {
		updateParams.Description = project.Description
	}

	if req.CoverPicture != nil {
		updateParams.CoverPicture = *req.CoverPicture
	} else {
		updateParams.CoverPicture = project.CoverPicture
	}

	if req.ReceiverName != nil {
		updateParams.ReceiverName = *req.ReceiverName
	} else {
		updateParams.ReceiverName = project.ReceiverName
	}

	if req.ReceiverNumber != nil {
		updateParams.ReceiverNumber = *req.ReceiverNumber
	} else {
		updateParams.ReceiverNumber = project.ReceiverNumber
	}

	if req.Address != nil {
		updateParams.Address = *req.Address
	} else {
		updateParams.Address = project.Address
	}

	if req.District != nil {
		updateParams.District = *req.District
	} else {
		updateParams.District = project.District
	}

	if req.City != nil {
		updateParams.City = *req.City
	} else {
		updateParams.City = project.City
	}

	if req.Country != nil {
		updateParams.Country = *req.Country
	} else {
		updateParams.Country = project.Country
	}

	if req.EndDate != nil {
		updateParams.EndDate = *req.EndDate
	} else {
		updateParams.EndDate = project.EndDate
	}

	if err = p.repository.UpdateProjectByID(ctx, updateParams); err != nil {
		return err
	}

	return nil
}

func (p *project) DeleteProject(ctx context.Context, userID, projectID uuid.UUID) error {
	project, err := p.repository.GetProjectByID(ctx, projectID)
	if err != nil {
		return ErrProjectNotFound
	}

	if userID != project.UserID {
		return ErrNotOwner
	}

	if err = p.repository.DeleteProjectByID(ctx, project.ID); err != nil {
		return err
	}

	return nil
}

func (p *project) GetAllCategories(ctx context.Context) ([]model.Category, error) {
	dbCategories, err := p.repository.GetAllCategories(ctx)
	if err != nil {
		return nil, err
	}

	var categories []model.Category
	for _, c := range dbCategories {
		categories = append(categories, model.Category{
			ID:           int(c.ID),
			Name:         c.Name,
			Description:  c.Description,
			CoverPicture: c.CoverPicture,
		})
	}

	return categories, nil
}

func (p *project) GetProjectUpdates(ctx context.Context, projectID uuid.UUID) ([]model.ProjectUpdate, error) {
	project, err := p.repository.GetProjectByID(ctx, projectID)
	if err != nil {
		return nil, ErrProjectNotFound
	}

	dbUpdates, err := p.repository.GetProjectUpdates(ctx, project.ID)
	if err != nil {
		return nil, err
	}

	var updates []model.ProjectUpdate
	for _, c := range dbUpdates {
		updates = append(updates, model.ProjectUpdate{
			ID:              c.ID,
			ProjectID:       c.ProjectID,
			AttachmentPhoto: c.AttachmentPhoto,
			Description:     c.Description,
			CreatedAt:       c.CreatedAt,
		})
	}

	return updates, nil
}

func (p *project) CreateProjectUpdate(ctx context.Context, userID uuid.UUID, req dto.CreateProjectUpdateRequest) (*model.ProjectUpdate, error) {
	pid, err := uuid.Parse(req.ProjectID)
	if err != nil {
		return nil, ErrProjectNotFound
	}

	project, err := p.repository.GetProjectByID(ctx, pid)
	if err != nil {
		return nil, ErrProjectNotFound
	}

	if userID != project.UserID {
		return nil, ErrNotOwner
	}

	args := db.CreateProjectUpdateParams{
		ProjectID:       project.ID,
		AttachmentPhoto: req.AttachmentPhoto,
		Description:     req.Description,
	}

	update, err := p.repository.CreateProjectUpdate(ctx, args)
	if err != nil {
		return nil, err
	}

	return &model.ProjectUpdate{
		ID:              update.ID,
		ProjectID:       update.ProjectID,
		AttachmentPhoto: update.AttachmentPhoto,
		Description:     update.Description,
		CreatedAt:       update.CreatedAt,
	}, err
}

func (p *project) GetUnresolvedMilestones(ctx context.Context) ([]model.Milestone, error) {
	return nil, nil
}

/*
func convertProjectStatus(dbStatus db.ProjectStatus) model.ProjectStatus {
	var status model.ProjectStatus
	switch dbStatus {
	case db.ProjectStatusOngoing:
		status = model.ProjectStatusOngoing
	case db.ProjectStatusEnded:
		status = model.ProjectStatusEnded
	}
	return status
}
*/
