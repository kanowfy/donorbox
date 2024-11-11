// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: project.sql

package db

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createProject = `-- name: CreateProject :one
INSERT INTO projects (
    user_id, category_id, title, description, cover_picture, end_date, receiver_number, receiver_name, address, district, city, country
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
)
RETURNING id, user_id, title, description, cover_picture, category_id, start_date, end_date, receiver_number, receiver_name, address, district, city, country, status
`

type CreateProjectParams struct {
	UserID         uuid.UUID
	CategoryID     int32
	Title          string
	Description    string
	CoverPicture   string
	EndDate        time.Time
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
		&i.StartDate,
		&i.EndDate,
		&i.ReceiverNumber,
		&i.ReceiverName,
		&i.Address,
		&i.District,
		&i.City,
		&i.Country,
		&i.Status,
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
	ProjectID       uuid.UUID
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

func (q *Queries) DeleteProjectByID(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteProjectByID, id)
	return err
}

const deleteProjectUpdate = `-- name: DeleteProjectUpdate :exec
DELETE FROM project_updates
WHERE id = $1
`

func (q *Queries) DeleteProjectUpdate(ctx context.Context, id uuid.UUID) error {
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
SELECT projects.id, projects.user_id, projects.title, projects.description, projects.cover_picture, projects.category_id, projects.start_date, projects.end_date, projects.receiver_number, projects.receiver_name, projects.address, projects.district, projects.city, projects.country, projects.status, COUNT(backings.project_id) as backing_count
FROM projects
LEFT JOIN backings ON projects.ID = backings.project_id
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
	ID             uuid.UUID
	UserID         uuid.UUID
	Title          string
	Description    string
	CoverPicture   string
	CategoryID     int32
	StartDate      time.Time
	EndDate        time.Time
	ReceiverNumber string
	ReceiverName   string
	Address        string
	District       string
	City           string
	Country        string
	Status         NullProjectStatus
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
			&i.StartDate,
			&i.EndDate,
			&i.ReceiverNumber,
			&i.ReceiverName,
			&i.Address,
			&i.District,
			&i.City,
			&i.Country,
			&i.Status,
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

const getFinishedProjects = `-- name: GetFinishedProjects :many
SELECT projects.id, projects.user_id, projects.title, projects.description, projects.cover_picture, projects.category_id, projects.start_date, projects.end_date, projects.receiver_number, projects.receiver_name, projects.address, projects.district, projects.city, projects.country, projects.status, COUNT(backings.project_id) as backing_count
FROM projects
JOIN backings ON projects.ID = backings.project_id
WHERE projects.status = 'finished'
GROUP BY projects.ID
ORDER BY end_date DESC
`

type GetFinishedProjectsRow struct {
	ID             uuid.UUID
	UserID         uuid.UUID
	Title          string
	Description    string
	CoverPicture   string
	CategoryID     int32
	StartDate      time.Time
	EndDate        time.Time
	ReceiverNumber string
	ReceiverName   string
	Address        string
	District       string
	City           string
	Country        string
	Status         NullProjectStatus
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
			&i.StartDate,
			&i.EndDate,
			&i.ReceiverNumber,
			&i.ReceiverName,
			&i.Address,
			&i.District,
			&i.City,
			&i.Country,
			&i.Status,
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

const getProjectByID = `-- name: GetProjectByID :one
SELECT id, user_id, title, description, cover_picture, category_id, start_date, end_date, receiver_number, receiver_name, address, district, city, country, status FROM projects
WHERE id = $1
`

func (q *Queries) GetProjectByID(ctx context.Context, id uuid.UUID) (Project, error) {
	row := q.db.QueryRow(ctx, getProjectByID, id)
	var i Project
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.Description,
		&i.CoverPicture,
		&i.CategoryID,
		&i.StartDate,
		&i.EndDate,
		&i.ReceiverNumber,
		&i.ReceiverName,
		&i.Address,
		&i.District,
		&i.City,
		&i.Country,
		&i.Status,
	)
	return i, err
}

const getProjectUpdates = `-- name: GetProjectUpdates :many
SELECT id, project_id, attachment_photo, description, created_at FROM project_updates
WHERE project_id = $1
ORDER BY created_at DESC
`

func (q *Queries) GetProjectUpdates(ctx context.Context, projectID uuid.UUID) ([]ProjectUpdate, error) {
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
SELECT id, user_id, title, description, cover_picture, category_id, start_date, end_date, receiver_number, receiver_name, address, district, city, country, status FROM projects
WHERE user_id = $1
ORDER BY start_date DESC
`

func (q *Queries) GetProjectsForUser(ctx context.Context, userID uuid.UUID) ([]Project, error) {
	rows, err := q.db.Query(ctx, getProjectsForUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Project
	for rows.Next() {
		var i Project
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Title,
			&i.Description,
			&i.CoverPicture,
			&i.CategoryID,
			&i.StartDate,
			&i.EndDate,
			&i.ReceiverNumber,
			&i.ReceiverName,
			&i.Address,
			&i.District,
			&i.City,
			&i.Country,
			&i.Status,
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
SELECT projects.id, projects.user_id, projects.title, projects.description, projects.cover_picture, projects.category_id, projects.start_date, projects.end_date, projects.receiver_number, projects.receiver_name, projects.address, projects.district, projects.city, projects.country, projects.status, COUNT(backings.project_id) as backing_count
FROM projects
LEFT JOIN backings ON projects.ID = backings.project_id
WHERE 
    to_tsvector('english', title || ' ' || description || ' ' || province || ' ' || country) @@ plainto_tsquery('english', $1::text)
GROUP BY projects.ID
ORDER BY backing_count DESC
LIMIT $3::integer OFFSET $2::integer
`

type SearchProjectsParams struct {
	SearchQuery string
	TotalOffset int32
	PageLimit   int32
}

type SearchProjectsRow struct {
	ID             uuid.UUID
	UserID         uuid.UUID
	Title          string
	Description    string
	CoverPicture   string
	CategoryID     int32
	StartDate      time.Time
	EndDate        time.Time
	ReceiverNumber string
	ReceiverName   string
	Address        string
	District       string
	City           string
	Country        string
	Status         NullProjectStatus
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
			&i.StartDate,
			&i.EndDate,
			&i.ReceiverNumber,
			&i.ReceiverName,
			&i.Address,
			&i.District,
			&i.City,
			&i.Country,
			&i.Status,
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

const updateProjectByID = `-- name: UpdateProjectByID :exec
UPDATE projects
SET title = $2, description = $3, cover_picture = $4, receiver_number=$5, receiver_name=$6, address=$7, district=$8, city=$9, country = $10, end_date = $11
WHERE id = $1
`

type UpdateProjectByIDParams struct {
	ID             uuid.UUID
	Title          string
	Description    string
	CoverPicture   string
	ReceiverNumber string
	ReceiverName   string
	Address        string
	District       string
	City           string
	Country        string
	EndDate        time.Time
}

func (q *Queries) UpdateProjectByID(ctx context.Context, arg UpdateProjectByIDParams) error {
	_, err := q.db.Exec(ctx, updateProjectByID,
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
	return err
}

const updateProjectFund = `-- name: UpdateProjectFund :exec
UPDATE projects SET current_amount = current_amount + $2::bigint
WHERE id = $1
`

type UpdateProjectFundParams struct {
	ID            uuid.UUID
	BackingAmount int64
}

func (q *Queries) UpdateProjectFund(ctx context.Context, arg UpdateProjectFundParams) error {
	_, err := q.db.Exec(ctx, updateProjectFund, arg.ID, arg.BackingAmount)
	return err
}

const updateProjectStatus = `-- name: UpdateProjectStatus :exec
UPDATE projects SET status = $2
WHERE id = $1
`

type UpdateProjectStatusParams struct {
	ID     uuid.UUID
	Status NullProjectStatus
}

func (q *Queries) UpdateProjectStatus(ctx context.Context, arg UpdateProjectStatusParams) error {
	_, err := q.db.Exec(ctx, updateProjectStatus, arg.ID, arg.Status)
	return err
}
