// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"
)

type Querier interface {
	ActivateUser(ctx context.Context, id int64) error
	CreateAuditLog(ctx context.Context, arg CreateAuditLogParams) (AuditTrail, error)
	CreateBacking(ctx context.Context, arg CreateBackingParams) (Backing, error)
	CreateEscrowUser(ctx context.Context, arg CreateEscrowUserParams) (EscrowUser, error)
	CreateMilestone(ctx context.Context, arg CreateMilestoneParams) (Milestone, error)
	CreateMilestoneCompletion(ctx context.Context, arg CreateMilestoneCompletionParams) (MilestoneCompletion, error)
	CreateNotification(ctx context.Context, arg CreateNotificationParams) (Notification, error)
	CreateProject(ctx context.Context, arg CreateProjectParams) (Project, error)
	CreateProjectUpdate(ctx context.Context, arg CreateProjectUpdateParams) (ProjectUpdate, error)
	CreateSocialLoginUser(ctx context.Context, arg CreateSocialLoginUserParams) (User, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteProjectByID(ctx context.Context, id int64) error
	DeleteProjectUpdate(ctx context.Context, id int64) error
	GetAllCategories(ctx context.Context) ([]Category, error)
	GetAllProjects(ctx context.Context, arg GetAllProjectsParams) ([]GetAllProjectsRow, error)
	GetAllUsers(ctx context.Context) ([]GetAllUsersRow, error)
	GetAuditHistory(ctx context.Context) ([]AuditTrail, error)
	GetBackingByID(ctx context.Context, id int64) (Backing, error)
	GetBackingCountForProject(ctx context.Context, projectID int64) (int64, error)
	GetBackingsForProject(ctx context.Context, projectID int64) ([]GetBackingsForProjectRow, error)
	GetBackingsForUser(ctx context.Context, userID *int64) ([]Backing, error)
	GetCategoriesCount(ctx context.Context) ([]GetCategoriesCountRow, error)
	GetEscrowUserByEmail(ctx context.Context, email string) (EscrowUser, error)
	GetEscrowUserByID(ctx context.Context, id int64) (EscrowUser, error)
	GetFinishedProjects(ctx context.Context) ([]GetFinishedProjectsRow, error)
	GetFirstBackingDonor(ctx context.Context, projectID int64) (GetFirstBackingDonorRow, error)
	GetMilestoneByID(ctx context.Context, id int64) (GetMilestoneByIDRow, error)
	GetMilestoneForProject(ctx context.Context, projectID int64) ([]GetMilestoneForProjectRow, error)
	GetMostBackingDonor(ctx context.Context, projectID int64) (GetMostBackingDonorRow, error)
	GetMostRecentBackingDonor(ctx context.Context, projectID int64) (GetMostRecentBackingDonorRow, error)
	GetNotificationsForUser(ctx context.Context, userID int64) ([]Notification, error)
	GetPendingProjects(ctx context.Context) ([]GetPendingProjectsRow, error)
	GetPendingVerificationUsers(ctx context.Context) ([]GetPendingVerificationUsersRow, error)
	GetProjectByID(ctx context.Context, id int64) (GetProjectByIDRow, error)
	GetProjectUpdates(ctx context.Context, projectID int64) ([]ProjectUpdate, error)
	GetProjectsForUser(ctx context.Context, userID int64) ([]GetProjectsForUserRow, error)
	GetStatsAggregation(ctx context.Context) (GetStatsAggregationRow, error)
	GetUnresolvedMilestones(ctx context.Context) ([]GetUnresolvedMilestonesRow, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	GetUserByID(ctx context.Context, id int64) (User, error)
	SearchProjects(ctx context.Context, arg SearchProjectsParams) ([]SearchProjectsRow, error)
	UpdateEscrowUserByID(ctx context.Context, arg UpdateEscrowUserByIDParams) error
	UpdateMilestoneFund(ctx context.Context, arg UpdateMilestoneFundParams) error
	UpdateMilestoneStatus(ctx context.Context, id int64) error
	UpdateProjectByID(ctx context.Context, arg UpdateProjectByIDParams) error
	UpdateProjectStatus(ctx context.Context, arg UpdateProjectStatusParams) error
	UpdateReadNotification(ctx context.Context, id int64) error
	UpdateUserByID(ctx context.Context, arg UpdateUserByIDParams) error
	UpdateUserPassword(ctx context.Context, arg UpdateUserPasswordParams) error
	UpdateVerificationStatus(ctx context.Context, arg UpdateVerificationStatusParams) error
}

var _ Querier = (*Queries)(nil)
