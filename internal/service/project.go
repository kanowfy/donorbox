package service

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/kanowfy/donorbox/internal/convert"
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
	GetProjectsForUser(ctx context.Context, userID int64) ([]model.Project, error)
	GetEndedProjects(ctx context.Context) ([]model.Project, error)
	GetPendingProjects(ctx context.Context) ([]dto.PendingProjectResponse, error)
	GetProjectDetails(ctx context.Context, projectID int64) (*model.Project, []model.Milestone, []model.Backing, []model.ProjectUpdate, *model.User, error)
	CreateProject(ctx context.Context, userID int64, req dto.CreateProjectRequest) (*dto.CreateProjectResponse, error)
	UpdateProject(ctx context.Context, userID, projectID int64, req dto.UpdateProjectRequest) error
	DeleteProject(ctx context.Context, userID, projectID int64) error
	GetAllCategories(ctx context.Context) ([]model.Category, error)
	GetCategoryByName(ctx context.Context, name string) (*model.Category, error)
	GetProjectUpdates(ctx context.Context, projectID int64) ([]model.ProjectUpdate, error)
	CreateProjectUpdate(ctx context.Context, userID int64, req dto.CreateProjectUpdateRequest) (*model.ProjectUpdate, error)
	GetUnresolvedMilestones(ctx context.Context) ([]dto.UnresolvedMilestoneDto, error)
	//CheckAndUpdateFinishedProjects(ctx context.Context) error
}

type project struct {
	repository     db.Querier
	backingService Backing
	userService    User
	auditSvc       AuditTrail
}

func NewProject(repository db.Querier, backingService Backing, userService User, auditSvc AuditTrail) Project {
	return &project{
		repository,
		backingService,
		userService,
		auditSvc,
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
		return nil, filters.Metadata{}, fmt.Errorf("GetAllProjects svc: %w", err)
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
			FundGoal:       p.FundGoal,
			TotalFund:      p.TotalFund,
			CoverPicture:   p.CoverPicture,
			ReceiverName:   p.ReceiverName,
			ReceiverNumber: p.ReceiverNumber,
			Address:        p.Address,
			District:       p.District,
			City:           p.City,
			Country:        p.Country,
			CreatedAt:      convert.MustPgTimestampToTime(p.CreatedAt),
			EndDate:        convert.MustPgTimestampToTime(p.EndDate),
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
			FundGoal:       p.FundGoal,
			TotalFund:      p.TotalFund,
			CoverPicture:   p.CoverPicture,
			ReceiverName:   p.ReceiverName,
			ReceiverNumber: p.ReceiverNumber,
			Address:        p.Address,
			District:       p.District,
			City:           p.City,
			Country:        p.Country,
			CreatedAt:      convert.MustPgTimestampToTime(p.CreatedAt),
			EndDate:        convert.MustPgTimestampToTime(p.EndDate),
			BackingCount:   &p.BackingCount,
		})
	}

	return projects, metadata, nil
}

func (p *project) GetProjectsForUser(ctx context.Context, userID int64) ([]model.Project, error) {
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
			FundGoal:       p.FundGoal,
			TotalFund:      p.TotalFund,
			CoverPicture:   p.CoverPicture,
			ReceiverName:   p.ReceiverName,
			ReceiverNumber: p.ReceiverNumber,
			Address:        p.Address,
			District:       p.District,
			City:           p.City,
			Country:        p.Country,
			CreatedAt:      convert.MustPgTimestampToTime(p.CreatedAt),
			EndDate:        convert.MustPgTimestampToTime(p.EndDate),
			Status:         model.ProjectStatus(p.Status),
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
			FundGoal:       p.FundGoal,
			TotalFund:      p.TotalFund,
			CoverPicture:   p.CoverPicture,
			ReceiverName:   p.ReceiverName,
			ReceiverNumber: p.ReceiverNumber,
			Address:        p.Address,
			District:       p.District,
			City:           p.City,
			Country:        p.Country,
			CreatedAt:      convert.MustPgTimestampToTime(p.CreatedAt),
			EndDate:        convert.MustPgTimestampToTime(p.EndDate),
		})
	}

	return projects, nil
}

func (p *project) GetPendingProjects(ctx context.Context) ([]dto.PendingProjectResponse, error) {
	dbProjects, err := p.repository.GetPendingProjects(ctx)
	if err != nil {
		return nil, err
	}

	var projects []dto.PendingProjectResponse

	for _, pj := range dbProjects {
		dbMilestones, err := p.repository.GetMilestoneForProject(ctx, pj.ID)
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
				BankDescription: m.BankDescription,
			})
		}

		projects = append(projects, dto.PendingProjectResponse{
			Project: model.Project{
				ID:             pj.ID,
				UserID:         pj.UserID,
				CategoryID:     pj.CategoryID,
				Title:          pj.Title,
				Description:    pj.Description,
				FundGoal:       pj.FundGoal,
				CoverPicture:   pj.CoverPicture,
				ReceiverName:   pj.ReceiverName,
				ReceiverNumber: pj.ReceiverNumber,
				Address:        pj.Address,
				District:       pj.District,
				City:           pj.City,
				Country:        pj.Country,
				CreatedAt:      convert.MustPgTimestampToTime(pj.CreatedAt),
				EndDate:        convert.MustPgTimestampToTime(pj.EndDate),
			},
			Milestones: milestones,
		})
	}

	return projects, nil
}

func (p *project) GetMilestonesForProject(ctx context.Context, projectID int64) ([]model.Milestone, error) {
	project, err := p.repository.GetProjectByID(ctx, projectID)
	if err != nil {
		return nil, ErrProjectNotFound
	}

	dbMilestones, err := p.repository.GetMilestoneForProject(ctx, project.ID)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var milestones []model.Milestone
	for _, m := range dbMilestones {
		ms := model.Milestone{
			ID:              m.ID,
			ProjectID:       m.ProjectID,
			Title:           m.Title,
			Description:     m.Description,
			FundGoal:        m.FundGoal,
			CurrentFund:     m.CurrentFund,
			BankDescription: m.BankDescription,
			Completed:       m.Completed,
		}

		if ms.Completed {
			ms.Completion = &model.MilestoneCompletion{
				TransferAmount: *m.TransferAmount,
				TransferNote:   m.TransferNote,
				TransferImage:  m.TransferImage,
				CompletedAt:    convert.MustPgTimestampToTime(m.CompletedAt),
			}
		}

		milestones = append(milestones, ms)
	}

	return milestones, nil
}

// TODO: refactor into one struct
func (p *project) GetProjectDetails(ctx context.Context, projectID int64) (*model.Project, []model.Milestone, []model.Backing, []model.ProjectUpdate, *model.User, error) {
	project, err := p.repository.GetProjectByID(ctx, projectID)
	if err != nil {
		return nil, nil, nil, nil, nil, fmt.Errorf("fetching project: %w", err)
	}

	milestones, err := p.GetMilestonesForProject(ctx, projectID)
	if err != nil {
		return nil, nil, nil, nil, nil, fmt.Errorf("fetching milestones: %w", err)
	}

	backings, err := p.backingService.GetBackingsForProject(ctx, project.ID)
	if err != nil {
		return nil, nil, nil, nil, nil, fmt.Errorf("fetching backings: %w", err)
	}

	updates, err := p.GetProjectUpdates(ctx, project.ID)
	if err != nil {
		return nil, nil, nil, nil, nil, fmt.Errorf("fetching project updates: %w", err)
	}

	user, err := p.userService.GetUserByID(ctx, project.UserID)
	if err != nil {
		return nil, nil, nil, nil, nil, fmt.Errorf("fetching user: %w", err)
	}

	return &model.Project{
		ID:             project.ID,
		UserID:         project.UserID,
		CategoryID:     project.CategoryID,
		Title:          project.Title,
		Description:    project.Description,
		FundGoal:       project.FundGoal,
		TotalFund:      project.TotalFund,
		CoverPicture:   project.CoverPicture,
		ReceiverName:   project.ReceiverName,
		ReceiverNumber: project.ReceiverNumber,
		Address:        project.Address,
		District:       project.District,
		City:           project.City,
		Country:        project.Country,
		CreatedAt:      convert.MustPgTimestampToTime(project.CreatedAt),
		EndDate:        convert.MustPgTimestampToTime(project.EndDate),
		Status:         convertProjectStatus(project.Status),
	}, milestones, backings, updates, user, nil
}

// TODO: put the queries in transaction
func (p *project) CreateProject(ctx context.Context, userID int64, req dto.CreateProjectRequest) (*dto.CreateProjectResponse, error) {
	queries := p.repository.(*db.Queries)
	q, tx, err := queries.BeginTX(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

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
		EndDate:        convert.TimeToPgTimestamp(req.EndDate),
	}

	project, err := q.CreateProject(ctx, projectArgs)
	if err != nil {
		return nil, err
	}

	if err := p.auditSvc.LogAction(ctx, LogActionParams{
		UserID:        &userID,
		EntityType:    "project",
		EntityID:      &project.ID,
		OperationType: "CREATE",
		NewValue:      project,
	}); err != nil {
		return nil, err
	}

	var milestones []model.Milestone
	var fundGoal int64

	for _, m := range req.Milestones {
		milestone, err := q.CreateMilestone(ctx, db.CreateMilestoneParams{
			ProjectID:       project.ID,
			Title:           m.Title,
			Description:     m.Description,
			FundGoal:        convert.MustStringToInt64(m.FundGoal),
			BankDescription: m.BankDescription,
		})

		if err != nil {
			return nil, err
		}

		if err := p.auditSvc.LogAction(ctx, LogActionParams{
			UserID:        &userID,
			EntityType:    "milestone",
			EntityID:      &milestone.ID,
			OperationType: "CREATE",
			NewValue:      milestone,
		}); err != nil {
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

		fundGoal += milestone.FundGoal
	}

	return &dto.CreateProjectResponse{
		Project: model.Project{
			ID:             project.ID,
			UserID:         project.UserID,
			CategoryID:     project.CategoryID,
			Title:          project.Title,
			Description:    project.Description,
			FundGoal:       fundGoal,
			TotalFund:      0,
			CoverPicture:   project.CoverPicture,
			ReceiverNumber: project.ReceiverNumber,
			ReceiverName:   project.ReceiverName,
			Address:        project.Address,
			District:       project.District,
			City:           project.City,
			Country:        project.Country,
			CreatedAt:      convert.MustPgTimestampToTime(project.CreatedAt),
			EndDate:        convert.MustPgTimestampToTime(project.EndDate),
			Status:         model.ProjectStatusPending,
		},
		Milestones: milestones,
	}, tx.Commit(ctx)
}

func (p *project) UpdateProject(ctx context.Context, userID, projectID int64, req dto.UpdateProjectRequest) error {
	queries := p.repository.(*db.Queries)
	q, tx, err := queries.BeginTX(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	project, err := q.GetProjectByID(ctx, projectID)
	if err != nil {
		return ErrProjectNotFound
	}

	// check permission of requesting user and whether the current amount is 0

	if userID != project.UserID {
		return ErrNotOwner
	}

	trailParams := LogActionParams{
		UserID:        &userID,
		EntityType:    "project",
		EntityID:      &projectID,
		OperationType: "UPDATE",
		OldValue:      project,
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
		updateParams.EndDate = convert.TimeToPgTimestamp(*req.EndDate)
	} else {
		updateParams.EndDate = project.EndDate
	}

	updated, err := q.UpdateProjectByID(ctx, updateParams)
	if err != nil {
		return err
	}

	trailParams.NewValue = updated

	if err = p.auditSvc.LogAction(ctx, trailParams); err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (p *project) DeleteProject(ctx context.Context, userID, projectID int64) error {
	queries := p.repository.(*db.Queries)
	q, tx, err := queries.BeginTX(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	project, err := q.GetProjectByID(ctx, projectID)
	if err != nil {
		return ErrProjectNotFound
	}

	if userID != project.UserID {
		return ErrNotOwner
	}

	if err = q.DeleteProjectByID(ctx, project.ID); err != nil {
		return err
	}

	if err := p.auditSvc.LogAction(ctx, LogActionParams{
		UserID:        &userID,
		EntityType:    "project",
		EntityID:      &project.ID,
		OperationType: "DELETE",
	}); err != nil {
		return err
	}

	return tx.Commit(ctx)
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

func (p *project) GetCategoryByName(ctx context.Context, name string) (*model.Category, error) {
	category, err := p.repository.GetCategoryByName(ctx, name)
	if err != nil {
		return nil, err
	}

	return &model.Category{
		ID:           int(category.ID),
		Name:         category.Name,
		Description:  category.Description,
		CoverPicture: category.CoverPicture,
	}, nil
}

func (p *project) GetProjectUpdates(ctx context.Context, projectID int64) ([]model.ProjectUpdate, error) {
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
			CreatedAt:       convert.MustPgTimestampToTime(c.CreatedAt),
		})
	}

	return updates, nil
}

func (p *project) CreateProjectUpdate(ctx context.Context, userID int64, req dto.CreateProjectUpdateRequest) (*model.ProjectUpdate, error) {
	queries := p.repository.(*db.Queries)
	q, tx, err := queries.BeginTX(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	project, err := q.GetProjectByID(ctx, req.ProjectID)
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

	update, err := q.CreateProjectUpdate(ctx, args)
	if err != nil {
		return nil, err
	}

	if err := p.auditSvc.LogAction(ctx, LogActionParams{
		UserID:        &userID,
		EntityType:    "project_update",
		EntityID:      &update.ID,
		OperationType: "CREATE",
		NewValue:      update,
	}); err != nil {
		return nil, err
	}

	return &model.ProjectUpdate{
		ID:              update.ID,
		ProjectID:       update.ProjectID,
		AttachmentPhoto: update.AttachmentPhoto,
		Description:     update.Description,
		CreatedAt:       convert.MustPgTimestampToTime(update.CreatedAt),
	}, tx.Commit(ctx)
}

func (p *project) GetUnresolvedMilestones(ctx context.Context) ([]dto.UnresolvedMilestoneDto, error) {
	dbMilestones, err := p.repository.GetUnresolvedMilestones(ctx)
	if err != nil {
		return nil, err
	}

	var milestones []dto.UnresolvedMilestoneDto
	for _, m := range dbMilestones {
		milestones = append(milestones, dto.UnresolvedMilestoneDto{
			Milestone: model.Milestone{
				ID:              m.ID,
				ProjectID:       m.ProjectID,
				Title:           m.Title,
				Description:     m.Description,
				CurrentFund:     m.CurrentFund,
				FundGoal:        m.FundGoal,
				BankDescription: m.BankDescription,
			},
			Address:        m.Address,
			District:       m.District,
			City:           m.City,
			Country:        m.Country,
			ReceiverName:   m.ReceiverName,
			ReceiverNumber: m.ReceiverNumber,
		})
	}

	return milestones, nil
}

func convertProjectStatus(dbStatus db.ProjectStatus) model.ProjectStatus {
	var status model.ProjectStatus
	switch dbStatus {
	case db.ProjectStatusPending:
		status = model.ProjectStatusPending
	case db.ProjectStatusOngoing:
		status = model.ProjectStatusOngoing
	case db.ProjectStatusFinished:
		status = model.ProjectStatusFinished
	}
	return status
}
