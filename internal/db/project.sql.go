// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: project.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createMilestone = `-- name: CreateMilestone :one
INSERT INTO milestones (
    project_id, title, description, fund_goal, bank_description
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING id, project_id, title, description, fund_goal, current_fund, bank_description, completed, created_at
`

type CreateMilestoneParams struct {
	ProjectID       int64
	Title           string
	Description     *string
	FundGoal        int64
	BankDescription string
}

func (q *Queries) CreateMilestone(ctx context.Context, arg CreateMilestoneParams) (Milestone, error) {
	row := q.db.QueryRow(ctx, createMilestone,
		arg.ProjectID,
		arg.Title,
		arg.Description,
		arg.FundGoal,
		arg.BankDescription,
	)
	var i Milestone
	err := row.Scan(
		&i.ID,
		&i.ProjectID,
		&i.Title,
		&i.Description,
		&i.FundGoal,
		&i.CurrentFund,
		&i.BankDescription,
		&i.Completed,
		&i.CreatedAt,
	)
	return i, err
}

const createMilestoneCompletion = `-- name: CreateMilestoneCompletion :one
INSERT INTO milestone_completions (
    milestone_id, transfer_amount, transfer_note, transfer_image
) VALUES (
    $1, $2, $3, $4
)
RETURNING id, milestone_id, transfer_amount, transfer_note, transfer_image, completed_at
`

type CreateMilestoneCompletionParams struct {
	MilestoneID    int64
	TransferAmount int64
	TransferNote   *string
	TransferImage  *string
}

func (q *Queries) CreateMilestoneCompletion(ctx context.Context, arg CreateMilestoneCompletionParams) (MilestoneCompletion, error) {
	row := q.db.QueryRow(ctx, createMilestoneCompletion,
		arg.MilestoneID,
		arg.TransferAmount,
		arg.TransferNote,
		arg.TransferImage,
	)
	var i MilestoneCompletion
	err := row.Scan(
		&i.ID,
		&i.MilestoneID,
		&i.TransferAmount,
		&i.TransferNote,
		&i.TransferImage,
		&i.CompletedAt,
	)
	return i, err
}

const createProject = `-- name: CreateProject :one
INSERT INTO projects (
    user_id, category_id, title, description, cover_picture, end_date, receiver_number, receiver_name, address, district, city, country
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
)
RETURNING id, user_id, title, description, cover_picture, category_id, end_date, receiver_number, receiver_name, address, district, city, country, status, created_at
`

type CreateProjectParams struct {
	UserID         int64
	CategoryID     int32
	Title          string
	Description    string
	CoverPicture   string
	EndDate        pgtype.Timestamptz
	ReceiverNumber string
	ReceiverName   string
	Address        string
	District       string
	City           string
	Country        string
}

func (q *Queries) CreateProject(ctx context.Context, arg CreateProjectParams) (Project, error) {
	row := q.db.QueryRow(ctx, createProject,
		arg.UserID,
		arg.CategoryID,
		arg.Title,
		arg.Description,
		arg.CoverPicture,
		arg.EndDate,
		arg.ReceiverNumber,
		arg.ReceiverName,
		arg.Address,
		arg.District,
		arg.City,
		arg.Country,
	)
	var i Project
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.Description,
		&i.CoverPicture,
		&i.CategoryID,
		&i.EndDate,
		&i.ReceiverNumber,
		&i.ReceiverName,
		&i.Address,
		&i.District,
		&i.City,
		&i.Country,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const createProjectUpdate = `-- name: CreateProjectUpdate :one
INSERT INTO project_updates (
    project_id, attachment_photo, description
) VALUES (
    $1, $2, $3
)
RETURNING id, project_id, attachment_photo, description, created_at
`

type CreateProjectUpdateParams struct {
	ProjectID       int64
	AttachmentPhoto *string
	Description     string
}

func (q *Queries) CreateProjectUpdate(ctx context.Context, arg CreateProjectUpdateParams) (ProjectUpdate, error) {
	row := q.db.QueryRow(ctx, createProjectUpdate, arg.ProjectID, arg.AttachmentPhoto, arg.Description)
	var i ProjectUpdate
	err := row.Scan(
		&i.ID,
		&i.ProjectID,
		&i.AttachmentPhoto,
		&i.Description,
		&i.CreatedAt,
	)
	return i, err
}

const deleteProjectByID = `-- name: DeleteProjectByID :exec
DELETE FROM projects WHERE id = $1
`

func (q *Queries) DeleteProjectByID(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteProjectByID, id)
	return err
}

const deleteProjectUpdate = `-- name: DeleteProjectUpdate :exec
DELETE FROM project_updates
WHERE id = $1
`

func (q *Queries) DeleteProjectUpdate(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteProjectUpdate, id)
	return err
}

const getAllCategories = `-- name: GetAllCategories :many
SELECT id, name, description, cover_picture FROM categories
`

func (q *Queries) GetAllCategories(ctx context.Context) ([]Category, error) {
	rows, err := q.db.Query(ctx, getAllCategories)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Category
	for rows.Next() {
		var i Category
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.CoverPicture,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllProjects = `-- name: GetAllProjects :many
SELECT projects.id, projects.user_id, projects.title, projects.description, projects.cover_picture, projects.category_id, projects.end_date, projects.receiver_number, projects.receiver_name, projects.address, projects.district, projects.city, projects.country, projects.status, projects.created_at, SUM(milestones.current_fund) AS total_fund,
SUM(milestones.fund_goal) AS fund_goal, COUNT(backings.project_id) as backing_count
FROM projects
LEFT JOIN backings ON projects.ID = backings.project_id
LEFT JOIN milestones ON projects.ID = milestones.project_id
WHERE category_id =
    CASE WHEN $1::integer > 0 THEN $1::integer ELSE category_id END
AND projects.status = 'ongoing'
GROUP BY projects.ID
ORDER BY backing_count DESC
LIMIT $3::integer OFFSET $2::integer
`

type GetAllProjectsParams struct {
	Category    int32
	TotalOffset int32
	PageLimit   int32
}

type GetAllProjectsRow struct {
	ID             int64
	UserID         int64
	Title          string
	Description    string
	CoverPicture   string
	CategoryID     int32
	EndDate        pgtype.Timestamptz
	ReceiverNumber string
	ReceiverName   string
	Address        string
	District       string
	City           string
	Country        string
	Status         ProjectStatus
	CreatedAt      pgtype.Timestamptz
	TotalFund      int64
	FundGoal       int64
	BackingCount   int64
}

func (q *Queries) GetAllProjects(ctx context.Context, arg GetAllProjectsParams) ([]GetAllProjectsRow, error) {
	rows, err := q.db.Query(ctx, getAllProjects, arg.Category, arg.TotalOffset, arg.PageLimit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllProjectsRow
	for rows.Next() {
		var i GetAllProjectsRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Title,
			&i.Description,
			&i.CoverPicture,
			&i.CategoryID,
			&i.EndDate,
			&i.ReceiverNumber,
			&i.ReceiverName,
			&i.Address,
			&i.District,
			&i.City,
			&i.Country,
			&i.Status,
			&i.CreatedAt,
			&i.TotalFund,
			&i.FundGoal,
			&i.BackingCount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCategoryByName = `-- name: GetCategoryByName :one
SELECT id, name, description, cover_picture FROM categories
WHERE name = $1
`

func (q *Queries) GetCategoryByName(ctx context.Context, name string) (Category, error) {
	row := q.db.QueryRow(ctx, getCategoryByName, name)
	var i Category
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.CoverPicture,
	)
	return i, err
}

const getFinishedProjects = `-- name: GetFinishedProjects :many
SELECT projects.id, projects.user_id, projects.title, projects.description, projects.cover_picture, projects.category_id, projects.end_date, projects.receiver_number, projects.receiver_name, projects.address, projects.district, projects.city, projects.country, projects.status, projects.created_at, SUM(milestones.current_fund) AS total_fund,
SUM(milestones.fund_goal) AS fund_goal, COUNT(backings.project_id) as backing_count
FROM projects
JOIN backings ON projects.ID = backings.project_id
LEFT JOIN milestones ON projects.ID = milestones.project_id
WHERE projects.status = 'finished'
GROUP BY projects.ID
ORDER BY end_date DESC
`

type GetFinishedProjectsRow struct {
	ID             int64
	UserID         int64
	Title          string
	Description    string
	CoverPicture   string
	CategoryID     int32
	EndDate        pgtype.Timestamptz
	ReceiverNumber string
	ReceiverName   string
	Address        string
	District       string
	City           string
	Country        string
	Status         ProjectStatus
	CreatedAt      pgtype.Timestamptz
	TotalFund      int64
	FundGoal       int64
	BackingCount   int64
}

func (q *Queries) GetFinishedProjects(ctx context.Context) ([]GetFinishedProjectsRow, error) {
	rows, err := q.db.Query(ctx, getFinishedProjects)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetFinishedProjectsRow
	for rows.Next() {
		var i GetFinishedProjectsRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Title,
			&i.Description,
			&i.CoverPicture,
			&i.CategoryID,
			&i.EndDate,
			&i.ReceiverNumber,
			&i.ReceiverName,
			&i.Address,
			&i.District,
			&i.City,
			&i.Country,
			&i.Status,
			&i.CreatedAt,
			&i.TotalFund,
			&i.FundGoal,
			&i.BackingCount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getMilestoneByID = `-- name: GetMilestoneByID :one
SELECT m.id, m.project_id, m.title, m.description, m.fund_goal, m.current_fund, m.bank_description, m.completed, m.created_at, c.transfer_amount, c.transfer_note, c.transfer_image, c.completed_at FROM milestones m
LEFT JOIN milestone_completions c ON m.id = c.milestone_id
WHERE m.id = $1
`

type GetMilestoneByIDRow struct {
	ID              int64
	ProjectID       int64
	Title           string
	Description     *string
	FundGoal        int64
	CurrentFund     int64
	BankDescription string
	Completed       bool
	CreatedAt       pgtype.Timestamptz
	TransferAmount  *int64
	TransferNote    *string
	TransferImage   *string
	CompletedAt     pgtype.Timestamptz
}

func (q *Queries) GetMilestoneByID(ctx context.Context, id int64) (GetMilestoneByIDRow, error) {
	row := q.db.QueryRow(ctx, getMilestoneByID, id)
	var i GetMilestoneByIDRow
	err := row.Scan(
		&i.ID,
		&i.ProjectID,
		&i.Title,
		&i.Description,
		&i.FundGoal,
		&i.CurrentFund,
		&i.BankDescription,
		&i.Completed,
		&i.CreatedAt,
		&i.TransferAmount,
		&i.TransferNote,
		&i.TransferImage,
		&i.CompletedAt,
	)
	return i, err
}

const getMilestoneForProject = `-- name: GetMilestoneForProject :many
SELECT m.id, m.project_id, m.title, m.description, m.fund_goal, m.current_fund, m.bank_description, m.completed, m.created_at, c.transfer_amount, c.transfer_note, c.transfer_image, c.completed_at FROM milestones m
LEFT JOIN milestone_completions c ON m.id = c.milestone_id
WHERE m.project_id = $1
ORDER BY m.id
`

type GetMilestoneForProjectRow struct {
	ID              int64
	ProjectID       int64
	Title           string
	Description     *string
	FundGoal        int64
	CurrentFund     int64
	BankDescription string
	Completed       bool
	CreatedAt       pgtype.Timestamptz
	TransferAmount  *int64
	TransferNote    *string
	TransferImage   *string
	CompletedAt     pgtype.Timestamptz
}

func (q *Queries) GetMilestoneForProject(ctx context.Context, projectID int64) ([]GetMilestoneForProjectRow, error) {
	rows, err := q.db.Query(ctx, getMilestoneForProject, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetMilestoneForProjectRow
	for rows.Next() {
		var i GetMilestoneForProjectRow
		if err := rows.Scan(
			&i.ID,
			&i.ProjectID,
			&i.Title,
			&i.Description,
			&i.FundGoal,
			&i.CurrentFund,
			&i.BankDescription,
			&i.Completed,
			&i.CreatedAt,
			&i.TransferAmount,
			&i.TransferNote,
			&i.TransferImage,
			&i.CompletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPendingProjects = `-- name: GetPendingProjects :many
SELECT projects.id, projects.user_id, projects.title, projects.description, projects.cover_picture, projects.category_id, projects.end_date, projects.receiver_number, projects.receiver_name, projects.address, projects.district, projects.city, projects.country, projects.status, projects.created_at, SUM(milestones.fund_goal) AS fund_goal
FROM projects
LEFT JOIN milestones ON projects.ID = milestones.project_id
WHERE status = 'pending'
GROUP BY projects.ID
ORDER BY projects.created_at DESC
`

type GetPendingProjectsRow struct {
	ID             int64
	UserID         int64
	Title          string
	Description    string
	CoverPicture   string
	CategoryID     int32
	EndDate        pgtype.Timestamptz
	ReceiverNumber string
	ReceiverName   string
	Address        string
	District       string
	City           string
	Country        string
	Status         ProjectStatus
	CreatedAt      pgtype.Timestamptz
	FundGoal       int64
}

func (q *Queries) GetPendingProjects(ctx context.Context) ([]GetPendingProjectsRow, error) {
	rows, err := q.db.Query(ctx, getPendingProjects)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetPendingProjectsRow
	for rows.Next() {
		var i GetPendingProjectsRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Title,
			&i.Description,
			&i.CoverPicture,
			&i.CategoryID,
			&i.EndDate,
			&i.ReceiverNumber,
			&i.ReceiverName,
			&i.Address,
			&i.District,
			&i.City,
			&i.Country,
			&i.Status,
			&i.CreatedAt,
			&i.FundGoal,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getProjectByID = `-- name: GetProjectByID :one
SELECT projects.id, projects.user_id, projects.title, projects.description, projects.cover_picture, projects.category_id, projects.end_date, projects.receiver_number, projects.receiver_name, projects.address, projects.district, projects.city, projects.country, projects.status, projects.created_at, SUM(milestones.current_fund) AS total_fund, SUM(milestones.fund_goal) AS fund_goal
FROM projects
LEFT JOIN milestones ON projects.ID = milestones.project_id
WHERE projects.ID = $1
GROUP BY projects.ID
`

type GetProjectByIDRow struct {
	ID             int64
	UserID         int64
	Title          string
	Description    string
	CoverPicture   string
	CategoryID     int32
	EndDate        pgtype.Timestamptz
	ReceiverNumber string
	ReceiverName   string
	Address        string
	District       string
	City           string
	Country        string
	Status         ProjectStatus
	CreatedAt      pgtype.Timestamptz
	TotalFund      int64
	FundGoal       int64
}

func (q *Queries) GetProjectByID(ctx context.Context, id int64) (GetProjectByIDRow, error) {
	row := q.db.QueryRow(ctx, getProjectByID, id)
	var i GetProjectByIDRow
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.Description,
		&i.CoverPicture,
		&i.CategoryID,
		&i.EndDate,
		&i.ReceiverNumber,
		&i.ReceiverName,
		&i.Address,
		&i.District,
		&i.City,
		&i.Country,
		&i.Status,
		&i.CreatedAt,
		&i.TotalFund,
		&i.FundGoal,
	)
	return i, err
}

const getProjectUpdates = `-- name: GetProjectUpdates :many
SELECT id, project_id, attachment_photo, description, created_at FROM project_updates
WHERE project_id = $1
ORDER BY created_at DESC
`

func (q *Queries) GetProjectUpdates(ctx context.Context, projectID int64) ([]ProjectUpdate, error) {
	rows, err := q.db.Query(ctx, getProjectUpdates, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ProjectUpdate
	for rows.Next() {
		var i ProjectUpdate
		if err := rows.Scan(
			&i.ID,
			&i.ProjectID,
			&i.AttachmentPhoto,
			&i.Description,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getProjectsForUser = `-- name: GetProjectsForUser :many
SELECT projects.id, projects.user_id, projects.title, projects.description, projects.cover_picture, projects.category_id, projects.end_date, projects.receiver_number, projects.receiver_name, projects.address, projects.district, projects.city, projects.country, projects.status, projects.created_at, SUM(milestones.current_fund) AS total_fund, SUM(milestones.fund_goal) AS fund_goal
FROM projects
LEFT JOIN milestones ON projects.ID = milestones.project_id
WHERE user_id = $1
GROUP BY projects.ID
ORDER BY projects.created_at DESC
`

type GetProjectsForUserRow struct {
	ID             int64
	UserID         int64
	Title          string
	Description    string
	CoverPicture   string
	CategoryID     int32
	EndDate        pgtype.Timestamptz
	ReceiverNumber string
	ReceiverName   string
	Address        string
	District       string
	City           string
	Country        string
	Status         ProjectStatus
	CreatedAt      pgtype.Timestamptz
	TotalFund      int64
	FundGoal       int64
}

func (q *Queries) GetProjectsForUser(ctx context.Context, userID int64) ([]GetProjectsForUserRow, error) {
	rows, err := q.db.Query(ctx, getProjectsForUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetProjectsForUserRow
	for rows.Next() {
		var i GetProjectsForUserRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Title,
			&i.Description,
			&i.CoverPicture,
			&i.CategoryID,
			&i.EndDate,
			&i.ReceiverNumber,
			&i.ReceiverName,
			&i.Address,
			&i.District,
			&i.City,
			&i.Country,
			&i.Status,
			&i.CreatedAt,
			&i.TotalFund,
			&i.FundGoal,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUnresolvedMilestones = `-- name: GetUnresolvedMilestones :many
SELECT milestones.id, milestones.project_id, milestones.title, milestones.description, milestones.fund_goal, milestones.current_fund, milestones.bank_description, milestones.completed, milestones.created_at, projects.address, projects.district, projects.city, projects.country, projects.receiver_name, projects.receiver_number
FROM milestones
JOIN projects ON milestones.project_id = projects.id
WHERE current_fund >= fund_goal
AND completed IS FALSE
`

type GetUnresolvedMilestonesRow struct {
	ID              int64
	ProjectID       int64
	Title           string
	Description     *string
	FundGoal        int64
	CurrentFund     int64
	BankDescription string
	Completed       bool
	CreatedAt       pgtype.Timestamptz
	Address         string
	District        string
	City            string
	Country         string
	ReceiverName    string
	ReceiverNumber  string
}

func (q *Queries) GetUnresolvedMilestones(ctx context.Context) ([]GetUnresolvedMilestonesRow, error) {
	rows, err := q.db.Query(ctx, getUnresolvedMilestones)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUnresolvedMilestonesRow
	for rows.Next() {
		var i GetUnresolvedMilestonesRow
		if err := rows.Scan(
			&i.ID,
			&i.ProjectID,
			&i.Title,
			&i.Description,
			&i.FundGoal,
			&i.CurrentFund,
			&i.BankDescription,
			&i.Completed,
			&i.CreatedAt,
			&i.Address,
			&i.District,
			&i.City,
			&i.Country,
			&i.ReceiverName,
			&i.ReceiverNumber,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const searchProjects = `-- name: SearchProjects :many
SELECT projects.id, projects.user_id, projects.title, projects.description, projects.cover_picture, projects.category_id, projects.end_date, projects.receiver_number, projects.receiver_name, projects.address, projects.district, projects.city, projects.country, projects.status, projects.created_at, SUM(milestones.current_fund) AS total_fund,
SUM(milestones.fund_goal) AS fund_goal, COUNT(backings.project_id) as backing_count
FROM projects
LEFT JOIN backings ON projects.ID = backings.project_id
LEFT JOIN milestones ON projects.ID = milestones.project_id
WHERE
    to_tsvector('english', projects.title || ' ' || projects.description || ' ' || city || ' ' || country) @@ plainto_tsquery('english', $1::text)
AND projects.status = 'ongoing'
GROUP BY projects.ID
LIMIT $3::integer OFFSET $2::integer
`

type SearchProjectsParams struct {
	SearchQuery string
	TotalOffset int32
	PageLimit   int32
}

type SearchProjectsRow struct {
	ID             int64
	UserID         int64
	Title          string
	Description    string
	CoverPicture   string
	CategoryID     int32
	EndDate        pgtype.Timestamptz
	ReceiverNumber string
	ReceiverName   string
	Address        string
	District       string
	City           string
	Country        string
	Status         ProjectStatus
	CreatedAt      pgtype.Timestamptz
	TotalFund      int64
	FundGoal       int64
	BackingCount   int64
}

func (q *Queries) SearchProjects(ctx context.Context, arg SearchProjectsParams) ([]SearchProjectsRow, error) {
	rows, err := q.db.Query(ctx, searchProjects, arg.SearchQuery, arg.TotalOffset, arg.PageLimit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SearchProjectsRow
	for rows.Next() {
		var i SearchProjectsRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Title,
			&i.Description,
			&i.CoverPicture,
			&i.CategoryID,
			&i.EndDate,
			&i.ReceiverNumber,
			&i.ReceiverName,
			&i.Address,
			&i.District,
			&i.City,
			&i.Country,
			&i.Status,
			&i.CreatedAt,
			&i.TotalFund,
			&i.FundGoal,
			&i.BackingCount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateMilestoneFund = `-- name: UpdateMilestoneFund :exec
UPDATE milestones
SET current_fund = current_fund + $2::bigint
WHERE id = $1
`

type UpdateMilestoneFundParams struct {
	ID     int64
	Amount int64
}

func (q *Queries) UpdateMilestoneFund(ctx context.Context, arg UpdateMilestoneFundParams) error {
	_, err := q.db.Exec(ctx, updateMilestoneFund, arg.ID, arg.Amount)
	return err
}

const updateMilestoneStatus = `-- name: UpdateMilestoneStatus :exec
UPDATE milestones
SET completed = TRUE
WHERE id = $1
`

func (q *Queries) UpdateMilestoneStatus(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, updateMilestoneStatus, id)
	return err
}

const updateProjectByID = `-- name: UpdateProjectByID :one
UPDATE projects
SET title = $2, description = $3, cover_picture = $4, receiver_number=$5, receiver_name=$6, address=$7, district=$8, city=$9, country = $10, end_date = $11
WHERE id = $1
RETURNING id, user_id, title, description, cover_picture, category_id, end_date, receiver_number, receiver_name, address, district, city, country, status, created_at
`

type UpdateProjectByIDParams struct {
	ID             int64
	Title          string
	Description    string
	CoverPicture   string
	ReceiverNumber string
	ReceiverName   string
	Address        string
	District       string
	City           string
	Country        string
	EndDate        pgtype.Timestamptz
}

func (q *Queries) UpdateProjectByID(ctx context.Context, arg UpdateProjectByIDParams) (Project, error) {
	row := q.db.QueryRow(ctx, updateProjectByID,
		arg.ID,
		arg.Title,
		arg.Description,
		arg.CoverPicture,
		arg.ReceiverNumber,
		arg.ReceiverName,
		arg.Address,
		arg.District,
		arg.City,
		arg.Country,
		arg.EndDate,
	)
	var i Project
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.Description,
		&i.CoverPicture,
		&i.CategoryID,
		&i.EndDate,
		&i.ReceiverNumber,
		&i.ReceiverName,
		&i.Address,
		&i.District,
		&i.City,
		&i.Country,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const updateProjectStatus = `-- name: UpdateProjectStatus :exec
UPDATE projects SET status = $2
WHERE id = $1
`

type UpdateProjectStatusParams struct {
	ID     int64
	Status ProjectStatus
}

func (q *Queries) UpdateProjectStatus(ctx context.Context, arg UpdateProjectStatusParams) error {
	_, err := q.db.Exec(ctx, updateProjectStatus, arg.ID, arg.Status)
	return err
}
