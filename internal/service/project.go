package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"slices"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/kanowfy/donorbox/internal/convert"
	"github.com/kanowfy/donorbox/internal/db"
	"github.com/kanowfy/donorbox/internal/dto"
	"github.com/kanowfy/donorbox/internal/model"
)

var (
	ErrProjectNotFound = errors.New("project not found")
	ErrNotOwner        = errors.New("user does not own the project")
)

type Project interface {
	GetAllProjects(ctx context.Context, categoryID int) ([]model.Project, error)
	SearchProjects(ctx context.Context, query string) ([]model.Project, error)
	GetProjectsForUser(ctx context.Context, userID int64) ([]model.Project, error)
	GetEndedProjects(ctx context.Context) ([]model.Project, error)
	GetPendingProjects(ctx context.Context) ([]dto.PendingProjectResponse, error)
	GetProjectDetails(ctx context.Context, projectID int64) (*model.Project, []model.Milestone, []model.Backing, *model.User, error)
	CreateProject(ctx context.Context, userID int64, req dto.CreateProjectRequest) (*dto.CreateProjectResponse, error)
	UpdateProject(ctx context.Context, userID, projectID int64, req dto.UpdateProjectRequest) error
	DeleteProject(ctx context.Context, userID, projectID int64) error
	GetAllCategories(ctx context.Context) ([]model.Category, error)
	GetCategoryByName(ctx context.Context, name string) (*model.Category, error)
	GetFundedMilestones(ctx context.Context) ([]dto.FundedMilestoneDto, error)
	CreateMilestoneProof(ctx context.Context, userID int64, req dto.CreateMilestoneProofRequest) error
	CheckUpdateRefutedMilestones(ctx context.Context) error
	CreateProjectReport(ctx context.Context, projectID int64, req dto.CreateProjectReportRequest) error
	GetProjectReports(ctx context.Context) ([]model.ProjecReport, error)
	GetDisputedProjects(ctx context.Context) ([]*dto.DisputedProject, error)
	//CheckAndUpdateFinishedProjects(ctx context.Context) error
}

type project struct {
	repository     db.Querier
	backingService Backing
	userService    User
	auditSvc       AuditTrail
}

const PROOF_PERIOD_DAY = 10

func NewProject(repository db.Querier, backingService Backing, userService User, auditSvc AuditTrail) Project {
	return &project{
		repository,
		backingService,
		userService,
		auditSvc,
	}
}

func (p *project) GetAllProjects(ctx context.Context, categoryIndex int) ([]model.Project, error) {
	dbProjects, err := p.repository.GetAllProjects(ctx, int32(categoryIndex))
	if err != nil {
		return nil, fmt.Errorf("get all projects: %w", err)
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
			BackingCount:   &p.BackingCount,
			Status:         convertProjectStatus(p.Status),
		})
	}

	return projects, nil
}

func (p *project) SearchProjects(ctx context.Context, query string) ([]model.Project, error) {
	dbProjects, err := p.repository.SearchProjects(ctx, query)
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
			Status:         convertProjectStatus(p.Status),
			CreatedAt:      convert.MustPgTimestampToTime(p.CreatedAt),
			EndDate:        convert.MustPgTimestampToTime(p.EndDate),
			BackingCount:   &p.BackingCount,
		})
	}

	return projects, nil
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
			Status:          model.MilestoneStatus(m.Status),
		}

		hasFundStatus := []model.MilestoneStatus{model.MilestoneStatusFundReleased, model.MilestoneStatusCompleted, model.MilestoneStatusRefuted}

		if slices.Contains(hasFundStatus, ms.Status) {
			ms.Completion = &model.MilestoneCompletion{
				TransferAmount: *m.TransferAmount,
				TransferNote:   m.FundReleasedNote,
				TransferImage:  m.FundReleasedImage,
				CreatedAt:      convert.MustPgTimestampToTime(m.FundReleasedAt),
			}
		}

		dbProofs, err := p.repository.GetSpendingProofsForMilestone(ctx, m.ID)
		if err != nil {
			return nil, err
		}

		if len(dbProofs) > 0 {
			var proofs []model.SpendingProof
			for _, pr := range dbProofs {
				proofs = append(proofs, model.SpendingProof{
					ID:            pr.ID,
					TransferImage: pr.TransferImage,
					Description:   pr.Description,
					ProofMedia:    pr.ProofMediaUrl,
					Status:        model.ProofStatus(pr.Status),
					RejectedCause: pr.RejectedCause,
					CreatedAt:     convert.MustPgTimestampToTime(pr.CreatedAt),
				})
			}

			ms.SpendingProofs = proofs
		}

		milestones = append(milestones, ms)
	}

	return milestones, nil
}

// TODO: refactor into one struct
func (p *project) GetProjectDetails(ctx context.Context, projectID int64) (*model.Project, []model.Milestone, []model.Backing, *model.User, error) {
	project, err := p.repository.GetProjectByID(ctx, projectID)
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("fetching project: %w", err)
	}

	milestones, err := p.GetMilestonesForProject(ctx, projectID)
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("fetching milestones: %w", err)
	}

	backings, err := p.backingService.GetBackingsForProject(ctx, project.ID)
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("fetching backings: %w", err)
	}

	user, err := p.userService.GetUserByID(ctx, project.UserID)
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("fetching user: %w", err)
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
	}, milestones, backings, user, nil
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

func (p *project) GetFundedMilestones(ctx context.Context) ([]dto.FundedMilestoneDto, error) {
	dbMilestones, err := p.repository.GetFundedMilestones(ctx)
	if err != nil {
		return nil, err
	}

	var milestones []dto.FundedMilestoneDto
	for _, m := range dbMilestones {
		ms := dto.FundedMilestoneDto{
			Milestone: model.Milestone{
				ID:              m.ID,
				ProjectID:       m.ProjectID,
				Title:           m.Title,
				Description:     m.Description,
				CurrentFund:     m.CurrentFund,
				FundGoal:        m.FundGoal,
				BankDescription: m.BankDescription,
				Status:          model.MilestoneStatus(m.Status),
			},
			Address:        m.Address,
			District:       m.District,
			City:           m.City,
			Country:        m.Country,
			ReceiverName:   m.ReceiverName,
			ReceiverNumber: m.ReceiverNumber,
		}

		if ms.Milestone.Status == model.MilestoneStatusFundReleased {
			ms.Milestone.Completion = &model.MilestoneCompletion{
				TransferAmount: *m.TransferAmount,
				TransferNote:   m.FundReleasedNote,
				TransferImage:  m.FundReleasedImage,
				CreatedAt:      convert.MustPgTimestampToTime(m.FundReleasedAt),
			}
		}

		dbProofs, err := p.repository.GetSpendingProofsForMilestone(ctx, m.ID)
		if err != nil {
			return nil, err
		}

		if len(dbProofs) > 0 {
			var proofs []model.SpendingProof
			for _, pr := range dbProofs {
				proofs = append(proofs, model.SpendingProof{
					ID:            pr.ID,
					TransferImage: pr.TransferImage,
					Description:   pr.Description,
					ProofMedia:    pr.ProofMediaUrl,
					Status:        model.ProofStatus(pr.Status),
					RejectedCause: pr.RejectedCause,
					CreatedAt:     convert.MustPgTimestampToTime(pr.CreatedAt),
				})
			}

			ms.Milestone.SpendingProofs = proofs
		}

		milestones = append(milestones, ms)
	}

	return milestones, nil
}

func (p *project) CreateMilestoneProof(ctx context.Context, userID int64, req dto.CreateMilestoneProofRequest) error {
	proof, err := p.repository.CreateSpendingProof(ctx, db.CreateSpendingProofParams{
		MilestoneID:   req.MilestoneID,
		TransferImage: req.Receipt,
		ProofMediaUrl: req.Media,
		Description:   req.Description,
	})
	if err != nil {
		return fmt.Errorf("creating spending proof: %w", err)
	}

	if err := p.auditSvc.LogAction(ctx, LogActionParams{
		UserID:        &userID,
		EntityType:    "spending_proof",
		EntityID:      &proof.ID,
		OperationType: "CREATE",
		NewValue:      proof,
	}); err != nil {
		return err
	}

	return nil
}

func (p *project) CheckUpdateRefutedMilestones(ctx context.Context) error {
	milestones, err := p.repository.GetAllMilestones(ctx)
	if err != nil {
		return err
	}

	for _, m := range milestones {
		if m.Status == db.MilestoneStatusFundReleased {
			proofs, err := p.repository.GetSpendingProofsForMilestone(ctx, m.ID)
			if err != nil {
				return err
			}

			if len(proofs) == 0 {
				completion, err := p.repository.GetMilestoneCompletionByMilestoneID(ctx, m.ID)
				if err != nil {
					return err
				}

				if time.Since(convert.MustPgTimestampToTime(completion.CreatedAt)).Hours() > PROOF_PERIOD_DAY*24 {
					if err := p.repository.UpdateMilestoneStatus(ctx, db.UpdateMilestoneStatusParams{
						ID:     m.ID,
						Status: db.MilestoneStatusRefuted,
					}); err != nil {
						return err
					}
				}
			} else {
				latestProof := proofs[0]
				if time.Since(convert.MustPgTimestampToTime(latestProof.CreatedAt)).Hours() > PROOF_PERIOD_DAY*24 {
					if err := p.repository.UpdateMilestoneStatus(ctx, db.UpdateMilestoneStatusParams{
						ID:     m.ID,
						Status: db.MilestoneStatusRefuted,
					}); err != nil {
						return err
					}

				}

			}

		}
	}

	return nil
}

func (p *project) CreateProjectReport(ctx context.Context, projectID int64, req dto.CreateProjectReportRequest) error {
	_, err := p.repository.CreateProjectReport(ctx, db.CreateProjectReportParams{
		ProjectID:   projectID,
		Email:       req.Email,
		FullName:    req.FullName,
		PhoneNumber: req.PhoneNumber,
		Relation:    req.Relation,
		Reason:      req.Reason,
		Details:     req.Details,
	})

	if err != nil {
		return fmt.Errorf("create project report: %w", err)
	}

	return nil
}

func (p *project) GetProjectReports(ctx context.Context) ([]model.ProjecReport, error) {
	dbReports, err := p.repository.GetAllProjectReports(ctx)
	if err != nil {
		return nil, fmt.Errorf("get project reports: %w", err)
	}

	var reports []model.ProjecReport
	for _, r := range dbReports {
		reports = append(reports, model.ProjecReport{
			ID:          r.ID,
			ProjectID:   r.ProjectID,
			Email:       r.Email,
			FullName:    r.FullName,
			PhoneNumber: r.PhoneNumber,
			Relation:    r.Relation,
			Reason:      r.Reason,
			Details:     r.Details,
			Status:      model.ReportStatus(r.Status),
			CreatedAt:   convert.MustPgTimestampToTime(r.CreatedAt),
		})
	}

	return reports, nil
}

func (p *project) GetDisputedProjects(ctx context.Context) ([]*dto.DisputedProject, error) {
	dbProjects, err := p.repository.GetDisputedProjects(ctx)
	if err != nil {
		return nil, fmt.Errorf("get disputed projects: %w", err)
	}

	projects := make([]*dto.DisputedProject, len(dbProjects))

	for i, pr := range dbProjects {
		project := &dto.DisputedProject{
			Project: model.Project{
				ID:             pr.ID,
				UserID:         pr.UserID,
				CategoryID:     pr.CategoryID,
				Title:          pr.Title,
				Description:    pr.Description,
				FundGoal:       pr.FundGoal,
				TotalFund:      pr.TotalFund,
				CoverPicture:   pr.CoverPicture,
				ReceiverName:   pr.ReceiverName,
				ReceiverNumber: pr.ReceiverNumber,
				Address:        pr.Address,
				District:       pr.District,
				City:           pr.City,
				Country:        pr.Country,
				//EndDate:        time.Time{}, do I need this?
				Status:       convertProjectStatus(pr.Status),
				BackingCount: &pr.BackingCount,
				CreatedAt:    convert.MustPgTimestampToTime(pr.CreatedAt),
			},
		}

		milestones, err := p.GetMilestonesForProject(ctx, pr.ID)
		if err != nil {
			return nil, fmt.Errorf("get disputed projects: %w", err)
		}

		var refutedMilestones []model.Milestone
		for _, m := range milestones {
			if m.Status == model.MilestoneStatusRefuted {
				refutedMilestones = append(refutedMilestones, m)
			}
		}

		project.Milestones = refutedMilestones

		dbReports, err := p.repository.GetResolvedProjectReportsForProject(ctx, pr.ID)
		if err != nil {
			return nil, fmt.Errorf("get disputed projects: %w", err)
		}

		if len(dbReports) != 0 {
			project.IsReported = true

			var reports []model.ProjecReport
			for _, r := range dbReports {
				reports = append(reports, model.ProjecReport{
					ID:          r.ID,
					ProjectID:   r.ProjectID,
					Email:       r.Email,
					FullName:    r.FullName,
					PhoneNumber: r.PhoneNumber,
					Relation:    r.Relation,
					Reason:      r.Reason,
					Details:     r.Details,
					Status:      model.ReportStatus(r.Status),
					CreatedAt:   convert.MustPgTimestampToTime(r.CreatedAt),
				})
			}

			project.Reports = reports
		}

		user, err := p.userService.GetUserByID(ctx, project.Project.UserID)
		if err != nil {
			return nil, fmt.Errorf("get disputed projects: %w", err)
		}
		project.User = *user

		projects[i] = project
	}

	return projects, nil
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
	case db.ProjectStatusRejected:
		status = model.ProjectStatusRejected
	case db.ProjectStatusDisputed:
		status = model.ProjectStatusDisputed
	case db.ProjectStatusStopped:
		status = model.ProjectStatusStopped
	}
	return status
}
