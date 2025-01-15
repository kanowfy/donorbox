// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: backing.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createBacking = `-- name: CreateBacking :one
INSERT INTO backings (
    project_id, user_id, amount, word_of_support
) VALUES (
    $1, $2, $3, $4
) 
RETURNING id, user_id, project_id, amount, word_of_support, created_at
`

type CreateBackingParams struct {
	ProjectID     int64
	UserID        *int64
	Amount        int64
	WordOfSupport *string
}

func (q *Queries) CreateBacking(ctx context.Context, arg CreateBackingParams) (Backing, error) {
	row := q.db.QueryRow(ctx, createBacking,
		arg.ProjectID,
		arg.UserID,
		arg.Amount,
		arg.WordOfSupport,
	)
	var i Backing
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.ProjectID,
		&i.Amount,
		&i.WordOfSupport,
		&i.CreatedAt,
	)
	return i, err
}

const getBackingByID = `-- name: GetBackingByID :one
SELECT id, user_id, project_id, amount, word_of_support, created_at FROM backings
WHERE id = $1
`

func (q *Queries) GetBackingByID(ctx context.Context, id int64) (Backing, error) {
	row := q.db.QueryRow(ctx, getBackingByID, id)
	var i Backing
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.ProjectID,
		&i.Amount,
		&i.WordOfSupport,
		&i.CreatedAt,
	)
	return i, err
}

const getBackingCountForProject = `-- name: GetBackingCountForProject :one
SELECT COUNT(*) AS backing_count
FROM backings
WHERE project_id = $1
`

func (q *Queries) GetBackingCountForProject(ctx context.Context, projectID int64) (int64, error) {
	row := q.db.QueryRow(ctx, getBackingCountForProject, projectID)
	var backing_count int64
	err := row.Scan(&backing_count)
	return backing_count, err
}

const getBackingsForProject = `-- name: GetBackingsForProject :many
SELECT backings.id, backings.user_id, backings.project_id, backings.amount, backings.word_of_support, backings.created_at, users.first_name, users.last_name, users.profile_picture
FROM backings
LEFT JOIN users
ON backings.user_id = users.id
WHERE project_id = $1
ORDER BY backings.created_at DESC
`

type GetBackingsForProjectRow struct {
	ID             int64
	UserID         *int64
	ProjectID      int64
	Amount         int64
	WordOfSupport  *string
	CreatedAt      pgtype.Timestamptz
	FirstName      *string
	LastName       *string
	ProfilePicture *string
}

func (q *Queries) GetBackingsForProject(ctx context.Context, projectID int64) ([]GetBackingsForProjectRow, error) {
	rows, err := q.db.Query(ctx, getBackingsForProject, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetBackingsForProjectRow
	for rows.Next() {
		var i GetBackingsForProjectRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.ProjectID,
			&i.Amount,
			&i.WordOfSupport,
			&i.CreatedAt,
			&i.FirstName,
			&i.LastName,
			&i.ProfilePicture,
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

const getBackingsForUser = `-- name: GetBackingsForUser :many
SELECT b.id, b.user_id, b.project_id, b.amount, b.word_of_support, b.created_at, p.title, p.cover_picture FROM backings b
JOIN projects p ON b.project_id = p.id
WHERE b.user_id = $1
ORDER BY b.created_at DESC
`

type GetBackingsForUserRow struct {
	ID            int64
	UserID        *int64
	ProjectID     int64
	Amount        int64
	WordOfSupport *string
	CreatedAt     pgtype.Timestamptz
	Title         string
	CoverPicture  string
}

func (q *Queries) GetBackingsForUser(ctx context.Context, userID *int64) ([]GetBackingsForUserRow, error) {
	rows, err := q.db.Query(ctx, getBackingsForUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetBackingsForUserRow
	for rows.Next() {
		var i GetBackingsForUserRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.ProjectID,
			&i.Amount,
			&i.WordOfSupport,
			&i.CreatedAt,
			&i.Title,
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

const getFirstBackingDonor = `-- name: GetFirstBackingDonor :one
SELECT backings.id, backings.user_id, backings.project_id, backings.amount, backings.word_of_support, backings.created_at, users.first_name, users.last_name, users.profile_picture
FROM backings
LEFT JOIN users
ON backings.user_id = users.id
WHERE project_id = $1
ORDER BY backings.created_at
LIMIT 1
`

type GetFirstBackingDonorRow struct {
	ID             int64
	UserID         *int64
	ProjectID      int64
	Amount         int64
	WordOfSupport  *string
	CreatedAt      pgtype.Timestamptz
	FirstName      *string
	LastName       *string
	ProfilePicture *string
}

func (q *Queries) GetFirstBackingDonor(ctx context.Context, projectID int64) (GetFirstBackingDonorRow, error) {
	row := q.db.QueryRow(ctx, getFirstBackingDonor, projectID)
	var i GetFirstBackingDonorRow
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.ProjectID,
		&i.Amount,
		&i.WordOfSupport,
		&i.CreatedAt,
		&i.FirstName,
		&i.LastName,
		&i.ProfilePicture,
	)
	return i, err
}

const getMostBackingDonor = `-- name: GetMostBackingDonor :one
SELECT backings.id, backings.user_id, backings.project_id, backings.amount, backings.word_of_support, backings.created_at, users.first_name, users.last_name, users.profile_picture
FROM backings
LEFT JOIN users
ON backings.user_id = users.id
WHERE project_id = $1
ORDER BY backings.amount DESC
LIMIT 1
`

type GetMostBackingDonorRow struct {
	ID             int64
	UserID         *int64
	ProjectID      int64
	Amount         int64
	WordOfSupport  *string
	CreatedAt      pgtype.Timestamptz
	FirstName      *string
	LastName       *string
	ProfilePicture *string
}

func (q *Queries) GetMostBackingDonor(ctx context.Context, projectID int64) (GetMostBackingDonorRow, error) {
	row := q.db.QueryRow(ctx, getMostBackingDonor, projectID)
	var i GetMostBackingDonorRow
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.ProjectID,
		&i.Amount,
		&i.WordOfSupport,
		&i.CreatedAt,
		&i.FirstName,
		&i.LastName,
		&i.ProfilePicture,
	)
	return i, err
}

const getMostRecentBackingDonor = `-- name: GetMostRecentBackingDonor :one
SELECT backings.id, backings.user_id, backings.project_id, backings.amount, backings.word_of_support, backings.created_at, users.first_name, users.last_name, users.profile_picture
FROM backings
LEFT JOIN users
ON backings.user_id = users.id
WHERE project_id = $1
ORDER BY backings.created_at DESC
LIMIT 1
`

type GetMostRecentBackingDonorRow struct {
	ID             int64
	UserID         *int64
	ProjectID      int64
	Amount         int64
	WordOfSupport  *string
	CreatedAt      pgtype.Timestamptz
	FirstName      *string
	LastName       *string
	ProfilePicture *string
}

func (q *Queries) GetMostRecentBackingDonor(ctx context.Context, projectID int64) (GetMostRecentBackingDonorRow, error) {
	row := q.db.QueryRow(ctx, getMostRecentBackingDonor, projectID)
	var i GetMostRecentBackingDonorRow
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.ProjectID,
		&i.Amount,
		&i.WordOfSupport,
		&i.CreatedAt,
		&i.FirstName,
		&i.LastName,
		&i.ProfilePicture,
	)
	return i, err
}

const getTotalBackingByMonth = `-- name: GetTotalBackingByMonth :many
WITH months AS (
    SELECT generate_series(
            date_trunc('month', (SELECT MIN(created_at) FROM backings)), 
            date_trunc('month', current_date), 
            interval '1 month') AS month
)
SELECT 
    to_char(months.month, 'YYYY-MM') AS month,
    COALESCE(SUM(backings.amount), 0)::bigint AS total_donated
FROM months
LEFT JOIN backings ON date_trunc('month', backings.created_at) = months.month
GROUP BY months.month
ORDER BY months.month
`

type GetTotalBackingByMonthRow struct {
	Month        string
	TotalDonated int64
}

func (q *Queries) GetTotalBackingByMonth(ctx context.Context) ([]GetTotalBackingByMonthRow, error) {
	rows, err := q.db.Query(ctx, getTotalBackingByMonth)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetTotalBackingByMonthRow
	for rows.Next() {
		var i GetTotalBackingByMonthRow
		if err := rows.Scan(&i.Month, &i.TotalDonated); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
