// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	ActivateUser(ctx context.Context, id uuid.UUID) error
	CreateBacking(ctx context.Context, arg CreateBackingParams) (Backing, error)
	CreateMilestone(ctx context.Context, arg CreateMilestoneParams) (Milestone, error)
	CreateProject(ctx context.Context, arg CreateProjectParams) (Project, error)
	CreateProjectUpdate(ctx context.Context, arg CreateProjectUpdateParams) (ProjectUpdate, error)
	CreateSocialLoginUser(ctx context.Context, arg CreateSocialLoginUserParams) (User, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteProjectByID(ctx context.Context, id uuid.UUID) error
	DeleteProjectUpdate(ctx context.Context, id uuid.UUID) error
	GetAllCategories(ctx context.Context) ([]Category, error)
	GetAllProjects(ctx context.Context, arg GetAllProjectsParams) ([]GetAllProjectsRow, error)
	GetAllUsers(ctx context.Context) ([]GetAllUsersRow, error)
	GetBackingByID(ctx context.Context, id uuid.UUID) (Backing, error)
	GetBackingCountForProject(ctx context.Context, projectID uuid.UUID) (int64, error)
	GetBackingsForProject(ctx context.Context, projectID uuid.UUID) ([]GetBackingsForProjectRow, error)
	GetBackingsForUser(ctx context.Context, userID uuid.UUID) ([]Backing, error)
	GetEscrowUser(ctx context.Context) (EscrowUser, error)
	GetEscrowUserByEmail(ctx context.Context, email string) (EscrowUser, error)
	GetEscrowUserByID(ctx context.Context, id uuid.UUID) (EscrowUser, error)
	GetFinishedProjects(ctx context.Context) ([]GetFinishedProjectsRow, error)
	GetFirstBackingDonor(ctx context.Context, projectID uuid.UUID) (GetFirstBackingDonorRow, error)
	GetMilestoneByID(ctx context.Context, id uuid.UUID) (Milestone, error)
	GetMilestoneForProject(ctx context.Context, projectID uuid.UUID) ([]Milestone, error)
	GetMostBackingDonor(ctx context.Context, projectID uuid.UUID) (GetMostBackingDonorRow, error)
	GetMostRecentBackingDonor(ctx context.Context, projectID uuid.UUID) (GetMostRecentBackingDonorRow, error)
	GetProjectByID(ctx context.Context, id uuid.UUID) (Project, error)
	GetProjectUpdates(ctx context.Context, projectID uuid.UUID) ([]ProjectUpdate, error)
	GetProjectsForUser(ctx context.Context, userID uuid.UUID) ([]Project, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (User, error)
	SearchProjects(ctx context.Context, arg SearchProjectsParams) ([]SearchProjectsRow, error)
	UpdateEscrowUserByID(ctx context.Context, arg UpdateEscrowUserByIDParams) error
	UpdateProjectByID(ctx context.Context, arg UpdateProjectByIDParams) error
	UpdateProjectStatus(ctx context.Context, arg UpdateProjectStatusParams) error
	UpdateUserByID(ctx context.Context, arg UpdateUserByIDParams) error
	UpdateUserPassword(ctx context.Context, arg UpdateUserPasswordParams) error
}

var _ Querier = (*Queries)(nil)
