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
RETURNING id, project_id, title, description, fund_goal, current_fund, bank_description, status, created_at
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
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const createMilestoneCompletion = `-- name: CreateMilestoneCompletion :one
INSERT INTO escrow_milestone_completions (
    milestone_id, transfer_amount, transfer_note, transfer_image
) VALUES (
    $1, $2, $3, $4
)
RETURNING id, milestone_id, transfer_amount, transfer_note, transfer_image, created_at
`

type CreateMilestoneCompletionParams struct {
	MilestoneID    int64
	TransferAmount int64
	TransferNote   *string
	TransferImage  *string
}

func (q *Queries) CreateMilestoneCompletion(ctx context.Context, arg CreateMilestoneCompletionParams) (EscrowMilestoneCompletion, error) {
	row := q.db.QueryRow(ctx, createMilestoneCompletion,
		arg.MilestoneID,
		arg.TransferAmount,
		arg.TransferNote,
		arg.TransferImage,
	)
	var i EscrowMilestoneCompletion
	err := row.Scan(
		&i.ID,
		&i.MilestoneID,
		&i.TransferAmount,
		&i.TransferNote,
		&i.TransferImage,
		&i.CreatedAt,
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

const createSpendingProof = `-- name: CreateSpendingProof :one
INSERT INTO user_spending_proofs (
    milestone_id, transfer_image, proof_media_url, description
) VALUES (
    $1, $2, $3, $4
) 
RETURNING id, milestone_id, transfer_image, proof_media_url, description, status, rejected_cause, created_at
`

type CreateSpendingProofParams struct {
	MilestoneID   int64
	TransferImage string
	ProofMediaUrl string
	Description   string
}

func (q *Queries) CreateSpendingProof(ctx context.Context, arg CreateSpendingProofParams) (UserSpendingProof, error) {
	row := q.db.QueryRow(ctx, createSpendingProof,
		arg.MilestoneID,
		arg.TransferImage,
		arg.ProofMediaUrl,
		arg.Description,
	)
	var i UserSpendingProof
	err := row.Scan(
		&i.ID,
		&i.MilestoneID,
		&i.TransferImage,
		&i.ProofMediaUrl,
		&i.Description,
		&i.Status,
		&i.RejectedCause,
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

const getAllMilestones = `-- name: GetAllMilestones :many
SELECT m.id, m.project_id, m.title, m.description, m.fund_goal, m.current_fund, m.bank_description, m.status, m.created_at, c.transfer_amount, c.transfer_note AS fund_released_note, c.transfer_image AS fund_released_image, c.created_at AS fund_released_at,
p.address, p.district, p.city, p.country, p.receiver_name, p.receiver_number
FROM milestones m
LEFT JOIN escrow_milestone_completions c ON m.id = c.milestone_id
JOIN projects p ON m.project_id = p.id
ORDER BY m.id
`

type GetAllMilestonesRow struct {
	ID                int64
	ProjectID         int64
	Title             string
	Description       *string
	FundGoal          int64
	CurrentFund       int64
	BankDescription   string
	Status            MilestoneStatus
	CreatedAt         pgtype.Timestamptz
	TransferAmount    *int64
	FundReleasedNote  *string
	FundReleasedImage *string
	FundReleasedAt    pgtype.Timestamptz
	Address           string
	District          string
	City              string
	Country           string
	ReceiverName      string
	ReceiverNumber    string
}

func (q *Queries) GetAllMilestones(ctx context.Context) ([]GetAllMilestonesRow, error) {
	rows, err := q.db.Query(ctx, getAllMilestones)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllMilestonesRow
	for rows.Next() {
		var i GetAllMilestonesRow
		if err := rows.Scan(
			&i.ID,
			&i.ProjectID,
			&i.Title,
			&i.Description,
			&i.FundGoal,
			&i.CurrentFund,
			&i.BankDescription,
			&i.Status,
			&i.CreatedAt,
			&i.TransferAmount,
			&i.FundReleasedNote,
			&i.FundReleasedImage,
			&i.FundReleasedAt,
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

const getAllProjects = `-- name: GetAllProjects :many
WITH aggregated_backings AS (
    SELECT project_id, COUNT(*) AS backing_count
    FROM backings
    GROUP BY project_id
),
aggregated_milestones AS (
    SELECT project_id, 
           SUM(current_fund) AS total_fund, 
           SUM(fund_goal) AS fund_goal
    FROM milestones
    GROUP BY project_id
)
SELECT p.id, p.user_id, p.title, p.description, p.cover_picture, p.category_id, p.end_date, p.receiver_number, p.receiver_name, p.address, p.district, p.city, p.country, p.status, p.created_at, 
       COALESCE(m.total_fund, 0) AS total_fund,
       COALESCE(m.fund_goal, 0) AS fund_goal,
       COALESCE(b.backing_count, 0) AS backing_count
FROM projects p
LEFT JOIN aggregated_backings b ON p.id = b.project_id
LEFT JOIN aggregated_milestones m ON p.id = m.project_id
WHERE p.category_id =
    CASE WHEN $1::integer > 0 THEN $1::integer ELSE p.category_id END
ORDER BY backing_count DESC
`

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

func (q *Queries) GetAllProjects(ctx context.Context, category int32) ([]GetAllProjectsRow, error) {
	rows, err := q.db.Query(ctx, getAllProjects, category)
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

const getDisputedProjects = `-- name: GetDisputedProjects :many
SELECT p.id, p.user_id, p.title, p.description, p.cover_picture, p.category_id, p.end_date, p.receiver_number, p.receiver_name, p.address, p.district, p.city, p.country, p.status, p.created_at, SUM(m.current_fund) AS total_fund,
SUM(m.fund_goal) AS fund_goal, COUNT(b.project_id) as backing_count
FROM projects p
LEFT JOIN backings b ON p.ID = b.project_id
LEFT JOIN milestones m ON p.ID = m.project_id
WHERE p.status = 'disputed'
GROUP BY p.ID
ORDER BY p.created_at
`

type GetDisputedProjectsRow struct {
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

func (q *Queries) GetDisputedProjects(ctx context.Context) ([]GetDisputedProjectsRow, error) {
	rows, err := q.db.Query(ctx, getDisputedProjects)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetDisputedProjectsRow
	for rows.Next() {
		var i GetDisputedProjectsRow
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

const getFundedMilestones = `-- name: GetFundedMilestones :many
SELECT m.id, m.project_id, m.title, m.description, m.fund_goal, m.current_fund, m.bank_description, m.status, m.created_at, c.transfer_amount, c.transfer_note AS fund_released_note, c.transfer_image AS fund_released_image, c.created_at AS fund_released_at,
p.address, p.district, p.city, p.country, p.receiver_name, p.receiver_number
FROM milestones m
JOIN projects p ON m.project_id = p.id
LEFT JOIN escrow_milestone_completions c ON m.id = c.milestone_id
WHERE current_fund >= fund_goal
ORDER BY m.id
`

type GetFundedMilestonesRow struct {
	ID                int64
	ProjectID         int64
	Title             string
	Description       *string
	FundGoal          int64
	CurrentFund       int64
	BankDescription   string
	Status            MilestoneStatus
	CreatedAt         pgtype.Timestamptz
	TransferAmount    *int64
	FundReleasedNote  *string
	FundReleasedImage *string
	FundReleasedAt    pgtype.Timestamptz
	Address           string
	District          string
	City              string
	Country           string
	ReceiverName      string
	ReceiverNumber    string
}

func (q *Queries) GetFundedMilestones(ctx context.Context) ([]GetFundedMilestonesRow, error) {
	rows, err := q.db.Query(ctx, getFundedMilestones)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetFundedMilestonesRow
	for rows.Next() {
		var i GetFundedMilestonesRow
		if err := rows.Scan(
			&i.ID,
			&i.ProjectID,
			&i.Title,
			&i.Description,
			&i.FundGoal,
			&i.CurrentFund,
			&i.BankDescription,
			&i.Status,
			&i.CreatedAt,
			&i.TransferAmount,
			&i.FundReleasedNote,
			&i.FundReleasedImage,
			&i.FundReleasedAt,
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

const getMilestoneAndProofs = `-- name: GetMilestoneAndProofs :many
SELECT p.id, p.milestone_id, p.transfer_image, p.proof_media_url, p.description, p.status, p.rejected_cause, p.created_at, m.title AS milestone_title, m.description AS milestone_description, m.fund_goal, m.current_fund, 
m.bank_description, m.status AS milestone_status, m.created_at AS milestone_created_at 
FROM user_spending_proofs p
JOIN milestones m ON m.id = p.milestone_id
ORDER BY p.created_at
`

type GetMilestoneAndProofsRow struct {
	ID                   int64
	MilestoneID          int64
	TransferImage        string
	ProofMediaUrl        string
	Description          string
	Status               ProofStatus
	RejectedCause        *string
	CreatedAt            pgtype.Timestamptz
	MilestoneTitle       string
	MilestoneDescription *string
	FundGoal             int64
	CurrentFund          int64
	BankDescription      string
	MilestoneStatus      MilestoneStatus
	MilestoneCreatedAt   pgtype.Timestamptz
}

func (q *Queries) GetMilestoneAndProofs(ctx context.Context) ([]GetMilestoneAndProofsRow, error) {
	rows, err := q.db.Query(ctx, getMilestoneAndProofs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetMilestoneAndProofsRow
	for rows.Next() {
		var i GetMilestoneAndProofsRow
		if err := rows.Scan(
			&i.ID,
			&i.MilestoneID,
			&i.TransferImage,
			&i.ProofMediaUrl,
			&i.Description,
			&i.Status,
			&i.RejectedCause,
			&i.CreatedAt,
			&i.MilestoneTitle,
			&i.MilestoneDescription,
			&i.FundGoal,
			&i.CurrentFund,
			&i.BankDescription,
			&i.MilestoneStatus,
			&i.MilestoneCreatedAt,
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
SELECT m.id, m.project_id, m.title, m.description, m.fund_goal, m.current_fund, m.bank_description, m.status, m.created_at, c.transfer_amount, c.transfer_note AS fund_released_note, c.transfer_image AS fund_released_image,
c.created_at AS fund_released_at
FROM milestones m
LEFT JOIN escrow_milestone_completions c ON m.id = c.milestone_id
WHERE m.id = $1
`

type GetMilestoneByIDRow struct {
	ID                int64
	ProjectID         int64
	Title             string
	Description       *string
	FundGoal          int64
	CurrentFund       int64
	BankDescription   string
	Status            MilestoneStatus
	CreatedAt         pgtype.Timestamptz
	TransferAmount    *int64
	FundReleasedNote  *string
	FundReleasedImage *string
	FundReleasedAt    pgtype.Timestamptz
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
		&i.Status,
		&i.CreatedAt,
		&i.TransferAmount,
		&i.FundReleasedNote,
		&i.FundReleasedImage,
		&i.FundReleasedAt,
	)
	return i, err
}

const getMilestoneCompletionByMilestoneID = `-- name: GetMilestoneCompletionByMilestoneID :one
SELECT id, milestone_id, transfer_amount, transfer_note, transfer_image, created_at FROM escrow_milestone_completions
WHERE milestone_id = $1
`

func (q *Queries) GetMilestoneCompletionByMilestoneID(ctx context.Context, milestoneID int64) (EscrowMilestoneCompletion, error) {
	row := q.db.QueryRow(ctx, getMilestoneCompletionByMilestoneID, milestoneID)
	var i EscrowMilestoneCompletion
	err := row.Scan(
		&i.ID,
		&i.MilestoneID,
		&i.TransferAmount,
		&i.TransferNote,
		&i.TransferImage,
		&i.CreatedAt,
	)
	return i, err
}

const getMilestoneForProject = `-- name: GetMilestoneForProject :many
SELECT m.id, m.project_id, m.title, m.description, m.fund_goal, m.current_fund, m.bank_description, m.status, m.created_at, c.transfer_amount, c.transfer_note AS fund_released_note, c.transfer_image AS fund_released_image,
c.created_at AS fund_released_at
FROM milestones m
LEFT JOIN escrow_milestone_completions c ON m.id = c.milestone_id
WHERE m.project_id = $1
ORDER BY m.id
`

type GetMilestoneForProjectRow struct {
	ID                int64
	ProjectID         int64
	Title             string
	Description       *string
	FundGoal          int64
	CurrentFund       int64
	BankDescription   string
	Status            MilestoneStatus
	CreatedAt         pgtype.Timestamptz
	TransferAmount    *int64
	FundReleasedNote  *string
	FundReleasedImage *string
	FundReleasedAt    pgtype.Timestamptz
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
			&i.Status,
			&i.CreatedAt,
			&i.TransferAmount,
			&i.FundReleasedNote,
			&i.FundReleasedImage,
			&i.FundReleasedAt,
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
WHERE projects.status = 'pending'
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

const getSpendingProofByID = `-- name: GetSpendingProofByID :one
SELECT id, milestone_id, transfer_image, proof_media_url, description, status, rejected_cause, created_at FROM user_spending_proofs
WHERE id = $1
`

func (q *Queries) GetSpendingProofByID(ctx context.Context, id int64) (UserSpendingProof, error) {
	row := q.db.QueryRow(ctx, getSpendingProofByID, id)
	var i UserSpendingProof
	err := row.Scan(
		&i.ID,
		&i.MilestoneID,
		&i.TransferImage,
		&i.ProofMediaUrl,
		&i.Description,
		&i.Status,
		&i.RejectedCause,
		&i.CreatedAt,
	)
	return i, err
}

const getSpendingProofsForMilestone = `-- name: GetSpendingProofsForMilestone :many
SELECT id, milestone_id, transfer_image, proof_media_url, description, status, rejected_cause, created_at FROM user_spending_proofs
WHERE milestone_id = $1
ORDER BY created_at DESC
`

func (q *Queries) GetSpendingProofsForMilestone(ctx context.Context, milestoneID int64) ([]UserSpendingProof, error) {
	rows, err := q.db.Query(ctx, getSpendingProofsForMilestone, milestoneID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []UserSpendingProof
	for rows.Next() {
		var i UserSpendingProof
		if err := rows.Scan(
			&i.ID,
			&i.MilestoneID,
			&i.TransferImage,
			&i.ProofMediaUrl,
			&i.Description,
			&i.Status,
			&i.RejectedCause,
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

const searchProjects = `-- name: SearchProjects :many
WITH aggregated_backings AS (
    SELECT project_id, COUNT(*) AS backing_count
    FROM backings
    GROUP BY project_id
),
aggregated_milestones AS (
    SELECT project_id, 
           SUM(current_fund) AS total_fund, 
           SUM(fund_goal) AS fund_goal
    FROM milestones
    GROUP BY project_id
)
SELECT p.id, p.user_id, p.title, p.description, p.cover_picture, p.category_id, p.end_date, p.receiver_number, p.receiver_name, p.address, p.district, p.city, p.country, p.status, p.created_at, 
       COALESCE(m.total_fund, 0) AS total_fund,
       COALESCE(m.fund_goal, 0) AS fund_goal,
       COALESCE(b.backing_count, 0) AS backing_count
FROM projects p
LEFT JOIN aggregated_backings b ON p.id = b.project_id
LEFT JOIN aggregated_milestones m ON p.id = m.project_id
WHERE
    to_tsvector('english', p.title || ' ' || p.description || ' ' || city || ' ' || country) @@ plainto_tsquery('english', $1::text)
ORDER BY backing_count DESC
`

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

func (q *Queries) SearchProjects(ctx context.Context, searchQuery string) ([]SearchProjectsRow, error) {
	rows, err := q.db.Query(ctx, searchProjects, searchQuery)
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
SET status = $2
WHERE id = $1
`

type UpdateMilestoneStatusParams struct {
	ID     int64
	Status MilestoneStatus
}

func (q *Queries) UpdateMilestoneStatus(ctx context.Context, arg UpdateMilestoneStatusParams) error {
	_, err := q.db.Exec(ctx, updateMilestoneStatus, arg.ID, arg.Status)
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

const updateSpendingProofStatus = `-- name: UpdateSpendingProofStatus :exec
UPDATE user_spending_proofs
SET status = $2, rejected_cause = $3
WHERE id = $1
`

type UpdateSpendingProofStatusParams struct {
	ID            int64
	Status        ProofStatus
	RejectedCause *string
}

func (q *Queries) UpdateSpendingProofStatus(ctx context.Context, arg UpdateSpendingProofStatusParams) error {
	_, err := q.db.Exec(ctx, updateSpendingProofStatus, arg.ID, arg.Status, arg.RejectedCause)
	return err
}
