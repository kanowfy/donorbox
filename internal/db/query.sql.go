// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const activateUser = `-- name: ActivateUser :exec
UPDATE users
SET activated = TRUE
WHERE id = $1
`

func (q *Queries) ActivateUser(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, activateUser, id)
	return err
}

const createBacking = `-- name: CreateBacking :one
INSERT INTO backings (
    project_id, backer_id, amount
) VALUES (
    $1, $2, $3
) 
RETURNING id, project_id, backer_id, amount, created_at, status
`

type CreateBackingParams struct {
	ProjectID pgtype.UUID `json:"project_id"`
	BackerID  pgtype.UUID `json:"backer_id"`
	Amount    int64       `json:"amount"`
}

func (q *Queries) CreateBacking(ctx context.Context, arg CreateBackingParams) (Backing, error) {
	row := q.db.QueryRow(ctx, createBacking, arg.ProjectID, arg.BackerID, arg.Amount)
	var i Backing
	err := row.Scan(
		&i.ID,
		&i.ProjectID,
		&i.BackerID,
		&i.Amount,
		&i.CreatedAt,
		&i.Status,
	)
	return i, err
}

const createProject = `-- name: CreateProject :one
INSERT INTO projects (
    user_id, category_id, title, description, cover_picture, goal_amount, country, province, end_date
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9
)
RETURNING id, user_id, category_id, title, description, cover_picture, goal_amount, current_amount, country, province, start_date, end_date, payment_id, is_active
`

type CreateProjectParams struct {
	UserID       pgtype.UUID        `json:"user_id"`
	CategoryID   int32              `json:"category_id"`
	Title        string             `json:"title"`
	Description  string             `json:"description"`
	CoverPicture string             `json:"cover_picture"`
	GoalAmount   int64              `json:"goal_amount"`
	Country      string             `json:"country"`
	Province     string             `json:"province"`
	EndDate      pgtype.Timestamptz `json:"end_date"`
}

func (q *Queries) CreateProject(ctx context.Context, arg CreateProjectParams) (Project, error) {
	row := q.db.QueryRow(ctx, createProject,
		arg.UserID,
		arg.CategoryID,
		arg.Title,
		arg.Description,
		arg.CoverPicture,
		arg.GoalAmount,
		arg.Country,
		arg.Province,
		arg.EndDate,
	)
	var i Project
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.CategoryID,
		&i.Title,
		&i.Description,
		&i.CoverPicture,
		&i.GoalAmount,
		&i.CurrentAmount,
		&i.Country,
		&i.Province,
		&i.StartDate,
		&i.EndDate,
		&i.PaymentID,
		&i.IsActive,
	)
	return i, err
}

const createProjectComment = `-- name: CreateProjectComment :one
INSERT INTO project_comments (
    project_id, backer_id, parent_comment_id, content
) VALUES (
    $1, $2, $3, $4
)
RETURNING id, project_id, backer_id, parent_comment_id, content, commented_at
`

type CreateProjectCommentParams struct {
	ProjectID       pgtype.UUID `json:"project_id"`
	BackerID        pgtype.UUID `json:"backer_id"`
	ParentCommentID pgtype.UUID `json:"parent_comment_id"`
	Content         string      `json:"content"`
}

func (q *Queries) CreateProjectComment(ctx context.Context, arg CreateProjectCommentParams) (ProjectComment, error) {
	row := q.db.QueryRow(ctx, createProjectComment,
		arg.ProjectID,
		arg.BackerID,
		arg.ParentCommentID,
		arg.Content,
	)
	var i ProjectComment
	err := row.Scan(
		&i.ID,
		&i.ProjectID,
		&i.BackerID,
		&i.ParentCommentID,
		&i.Content,
		&i.CommentedAt,
	)
	return i, err
}

const createProjectUpdate = `-- name: CreateProjectUpdate :one
INSERT INTO project_updates (
    project_id, description
) VALUES (
    $1, $2
)
RETURNING id, project_id, description, created_at
`

type CreateProjectUpdateParams struct {
	ProjectID   pgtype.UUID `json:"project_id"`
	Description string      `json:"description"`
}

func (q *Queries) CreateProjectUpdate(ctx context.Context, arg CreateProjectUpdateParams) (ProjectUpdate, error) {
	row := q.db.QueryRow(ctx, createProjectUpdate, arg.ProjectID, arg.Description)
	var i ProjectUpdate
	err := row.Scan(
		&i.ID,
		&i.ProjectID,
		&i.Description,
		&i.CreatedAt,
	)
	return i, err
}

const createTransaction = `-- name: CreateTransaction :one
INSERT INTO transactions (
    backing_id, transaction_type, amount, initiator_id, recipient_id
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING id, backing_id, transaction_type, amount, initiator_id, recipient_id, status, created_at
`

type CreateTransactionParams struct {
	BackingID       pgtype.UUID     `json:"backing_id"`
	TransactionType TransactionType `json:"transaction_type"`
	Amount          int64           `json:"amount"`
	InitiatorID     pgtype.UUID     `json:"initiator_id"`
	RecipientID     pgtype.UUID     `json:"recipient_id"`
}

func (q *Queries) CreateTransaction(ctx context.Context, arg CreateTransactionParams) (Transaction, error) {
	row := q.db.QueryRow(ctx, createTransaction,
		arg.BackingID,
		arg.TransactionType,
		arg.Amount,
		arg.InitiatorID,
		arg.RecipientID,
	)
	var i Transaction
	err := row.Scan(
		&i.ID,
		&i.BackingID,
		&i.TransactionType,
		&i.Amount,
		&i.InitiatorID,
		&i.RecipientID,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (
    email, hashed_password, first_name, last_name 
) VALUES (
    $1, $2, $3, $4
)
RETURNING id, email, hashed_password, first_name, last_name, profile_picture, activated, user_type, created_at
`

type CreateUserParams struct {
	Email          string `json:"email"`
	HashedPassword string `json:"hashed_password"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.Email,
		arg.HashedPassword,
		arg.FirstName,
		arg.LastName,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.HashedPassword,
		&i.FirstName,
		&i.LastName,
		&i.ProfilePicture,
		&i.Activated,
		&i.UserType,
		&i.CreatedAt,
	)
	return i, err
}

const deleteProjectByID = `-- name: DeleteProjectByID :exec
DELETE FROM projects WHERE id = $1
`

func (q *Queries) DeleteProjectByID(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteProjectByID, id)
	return err
}

const deleteProjectComment = `-- name: DeleteProjectComment :exec
DELETE FROM project_comments
WHERE id = $1
`

func (q *Queries) DeleteProjectComment(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteProjectComment, id)
	return err
}

const deleteProjectUpdate = `-- name: DeleteProjectUpdate :exec
DELETE FROM project_updates
WHERE id = $1
`

func (q *Queries) DeleteProjectUpdate(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteProjectUpdate, id)
	return err
}

const getAllCategories = `-- name: GetAllCategories :many
SELECT id, name FROM categories
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
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
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

SELECT projects.id, projects.user_id, projects.category_id, projects.title, projects.description, projects.cover_picture, projects.goal_amount, projects.current_amount, projects.country, projects.province, projects.start_date, projects.end_date, projects.payment_id, projects.is_active, COUNT(backings.project_id) as backing_count
FROM projects
LEFT JOIN backings ON projects.ID = backings.project_id
WHERE category_id = 
    CASE WHEN $1::integer > 0 THEN $1::integer ELSE category_id END
GROUP BY projects.ID
ORDER BY backing_count DESC
LIMIT $3::integer OFFSET $2::integer
`

type GetAllProjectsParams struct {
	Category    int32 `json:"category"`
	TotalOffset int32 `json:"total_offset"`
	PageLimit   int32 `json:"page_limit"`
}

type GetAllProjectsRow struct {
	ID            pgtype.UUID        `json:"id"`
	UserID        pgtype.UUID        `json:"user_id"`
	CategoryID    int32              `json:"category_id"`
	Title         string             `json:"title"`
	Description   string             `json:"description"`
	CoverPicture  string             `json:"cover_picture"`
	GoalAmount    int64              `json:"goal_amount"`
	CurrentAmount int64              `json:"current_amount"`
	Country       string             `json:"country"`
	Province      string             `json:"province"`
	StartDate     pgtype.Timestamptz `json:"start_date"`
	EndDate       pgtype.Timestamptz `json:"end_date"`
	PaymentID     pgtype.Text        `json:"payment_id"`
	IsActive      bool               `json:"is_active"`
	BackingCount  int64              `json:"backing_count"`
}

// :::::::::: PROJECT ::::::::::--
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
			&i.CategoryID,
			&i.Title,
			&i.Description,
			&i.CoverPicture,
			&i.GoalAmount,
			&i.CurrentAmount,
			&i.Country,
			&i.Province,
			&i.StartDate,
			&i.EndDate,
			&i.PaymentID,
			&i.IsActive,
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

const getAllTransactions = `-- name: GetAllTransactions :many

SELECT id, backing_id, transaction_type, amount, initiator_id, recipient_id, status, created_at FROM transactions
ORDER BY created_at DESC
`

// :::::::::: TRANSACTION ::::::::::--
func (q *Queries) GetAllTransactions(ctx context.Context) ([]Transaction, error) {
	rows, err := q.db.Query(ctx, getAllTransactions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Transaction
	for rows.Next() {
		var i Transaction
		if err := rows.Scan(
			&i.ID,
			&i.BackingID,
			&i.TransactionType,
			&i.Amount,
			&i.InitiatorID,
			&i.RecipientID,
			&i.Status,
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

const getAllUsers = `-- name: GetAllUsers :many

SELECT id, email, first_name, last_name, profile_picture, activated, user_type, created_at FROM users
`

type GetAllUsersRow struct {
	ID             pgtype.UUID        `json:"id"`
	Email          string             `json:"email"`
	FirstName      string             `json:"first_name"`
	LastName       string             `json:"last_name"`
	ProfilePicture pgtype.Text        `json:"profile_picture"`
	Activated      bool               `json:"activated"`
	UserType       UserType           `json:"user_type"`
	CreatedAt      pgtype.Timestamptz `json:"created_at"`
}

// :::::::::: USER ::::::::::--
func (q *Queries) GetAllUsers(ctx context.Context) ([]GetAllUsersRow, error) {
	rows, err := q.db.Query(ctx, getAllUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllUsersRow
	for rows.Next() {
		var i GetAllUsersRow
		if err := rows.Scan(
			&i.ID,
			&i.Email,
			&i.FirstName,
			&i.LastName,
			&i.ProfilePicture,
			&i.Activated,
			&i.UserType,
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

const getBackingByID = `-- name: GetBackingByID :one
SELECT id, project_id, backer_id, amount, created_at, status FROM backings
WHERE id = $1
`

func (q *Queries) GetBackingByID(ctx context.Context, id pgtype.UUID) (Backing, error) {
	row := q.db.QueryRow(ctx, getBackingByID, id)
	var i Backing
	err := row.Scan(
		&i.ID,
		&i.ProjectID,
		&i.BackerID,
		&i.Amount,
		&i.CreatedAt,
		&i.Status,
	)
	return i, err
}

const getBackingsForProject = `-- name: GetBackingsForProject :many

SELECT id, project_id, backer_id, amount, created_at, status FROM backings
WHERE project_id = $1
ORDER BY created_at DESC
`

// :::::::::: BACKING ::::::::::--
func (q *Queries) GetBackingsForProject(ctx context.Context, projectID pgtype.UUID) ([]Backing, error) {
	rows, err := q.db.Query(ctx, getBackingsForProject, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Backing
	for rows.Next() {
		var i Backing
		if err := rows.Scan(
			&i.ID,
			&i.ProjectID,
			&i.BackerID,
			&i.Amount,
			&i.CreatedAt,
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

const getBackingsForUser = `-- name: GetBackingsForUser :many
SELECT id, project_id, backer_id, amount, created_at, status FROM backings
WHERE backer_id = $1
`

func (q *Queries) GetBackingsForUser(ctx context.Context, backerID pgtype.UUID) ([]Backing, error) {
	rows, err := q.db.Query(ctx, getBackingsForUser, backerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Backing
	for rows.Next() {
		var i Backing
		if err := rows.Scan(
			&i.ID,
			&i.ProjectID,
			&i.BackerID,
			&i.Amount,
			&i.CreatedAt,
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

const getEscrowUserByEmail = `-- name: GetEscrowUserByEmail :one
SELECT id, email, hashed_password, user_type, payment_id, created_at FROM escrow_users
WHERE email = $1
`

func (q *Queries) GetEscrowUserByEmail(ctx context.Context, email string) (EscrowUser, error) {
	row := q.db.QueryRow(ctx, getEscrowUserByEmail, email)
	var i EscrowUser
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.HashedPassword,
		&i.UserType,
		&i.PaymentID,
		&i.CreatedAt,
	)
	return i, err
}

const getEscrowUserByID = `-- name: GetEscrowUserByID :one
SELECT id, email, hashed_password, user_type, payment_id, created_at FROM escrow_users
WHERE id = $1
`

func (q *Queries) GetEscrowUserByID(ctx context.Context, id pgtype.UUID) (EscrowUser, error) {
	row := q.db.QueryRow(ctx, getEscrowUserByID, id)
	var i EscrowUser
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.HashedPassword,
		&i.UserType,
		&i.PaymentID,
		&i.CreatedAt,
	)
	return i, err
}

const getProjectByID = `-- name: GetProjectByID :one
SELECT id, user_id, category_id, title, description, cover_picture, goal_amount, current_amount, country, province, start_date, end_date, payment_id, is_active FROM projects
WHERE id = $1
`

func (q *Queries) GetProjectByID(ctx context.Context, id pgtype.UUID) (Project, error) {
	row := q.db.QueryRow(ctx, getProjectByID, id)
	var i Project
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.CategoryID,
		&i.Title,
		&i.Description,
		&i.CoverPicture,
		&i.GoalAmount,
		&i.CurrentAmount,
		&i.Country,
		&i.Province,
		&i.StartDate,
		&i.EndDate,
		&i.PaymentID,
		&i.IsActive,
	)
	return i, err
}

const getProjectComments = `-- name: GetProjectComments :many

SELECT id, project_id, backer_id, parent_comment_id, content, commented_at FROM project_comments
WHERE project_id = $1
`

// :::::::::: PROJECT COMMENT ::::::::::--
func (q *Queries) GetProjectComments(ctx context.Context, projectID pgtype.UUID) ([]ProjectComment, error) {
	rows, err := q.db.Query(ctx, getProjectComments, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ProjectComment
	for rows.Next() {
		var i ProjectComment
		if err := rows.Scan(
			&i.ID,
			&i.ProjectID,
			&i.BackerID,
			&i.ParentCommentID,
			&i.Content,
			&i.CommentedAt,
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

const getProjectUpdates = `-- name: GetProjectUpdates :many

SELECT id, project_id, description, created_at FROM project_updates
WHERE project_id = $1
ORDER BY created_at DESC
`

// :::::::::: PROJECT UPDATE ::::::::::--
func (q *Queries) GetProjectUpdates(ctx context.Context, projectID pgtype.UUID) ([]ProjectUpdate, error) {
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

const getTransactionAudit = `-- name: GetTransactionAudit :many
SELECT id, backing_id, transaction_type, amount, initiator_id, recipient_id, status, created_at FROM transactions
WHERE backing_id = $1 ORDER BY created_at ASC
`

func (q *Queries) GetTransactionAudit(ctx context.Context, backingID pgtype.UUID) ([]Transaction, error) {
	rows, err := q.db.Query(ctx, getTransactionAudit, backingID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Transaction
	for rows.Next() {
		var i Transaction
		if err := rows.Scan(
			&i.ID,
			&i.BackingID,
			&i.TransactionType,
			&i.Amount,
			&i.InitiatorID,
			&i.RecipientID,
			&i.Status,
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

const getTransactionByID = `-- name: GetTransactionByID :one
SELECT id, backing_id, transaction_type, amount, initiator_id, recipient_id, status, created_at FROM transactions
WHERE id = $1
`

func (q *Queries) GetTransactionByID(ctx context.Context, id pgtype.UUID) (Transaction, error) {
	row := q.db.QueryRow(ctx, getTransactionByID, id)
	var i Transaction
	err := row.Scan(
		&i.ID,
		&i.BackingID,
		&i.TransactionType,
		&i.Amount,
		&i.InitiatorID,
		&i.RecipientID,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, email, hashed_password, first_name, last_name, profile_picture, activated, user_type, created_at FROM users
WHERE email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.HashedPassword,
		&i.FirstName,
		&i.LastName,
		&i.ProfilePicture,
		&i.Activated,
		&i.UserType,
		&i.CreatedAt,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, email, hashed_password, first_name, last_name, profile_picture, activated, user_type, created_at FROM users
WHERE id = $1
`

func (q *Queries) GetUserByID(ctx context.Context, id pgtype.UUID) (User, error) {
	row := q.db.QueryRow(ctx, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.HashedPassword,
		&i.FirstName,
		&i.LastName,
		&i.ProfilePicture,
		&i.Activated,
		&i.UserType,
		&i.CreatedAt,
	)
	return i, err
}

const searchProjects = `-- name: SearchProjects :many
SELECT id, user_id, category_id, title, description, cover_picture, goal_amount, current_amount, country, province, start_date, end_date, payment_id, is_active
FROM projects
WHERE 
    to_tsvector('english', title || ' ' || description || ' ' || province || ' ' || country) @@ plainto_tsquery('english', $1::text)
LIMIT $3::integer OFFSET $2::integer
`

type SearchProjectsParams struct {
	SearchQuery string `json:"search_query"`
	TotalOffset int32  `json:"total_offset"`
	PageLimit   int32  `json:"page_limit"`
}

func (q *Queries) SearchProjects(ctx context.Context, arg SearchProjectsParams) ([]Project, error) {
	rows, err := q.db.Query(ctx, searchProjects, arg.SearchQuery, arg.TotalOffset, arg.PageLimit)
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
			&i.CategoryID,
			&i.Title,
			&i.Description,
			&i.CoverPicture,
			&i.GoalAmount,
			&i.CurrentAmount,
			&i.Country,
			&i.Province,
			&i.StartDate,
			&i.EndDate,
			&i.PaymentID,
			&i.IsActive,
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

const updateEscrowUserPaymentID = `-- name: UpdateEscrowUserPaymentID :exec
UPDATE escrow_users
SET payment_id = $2
WHERE id = $1
`

type UpdateEscrowUserPaymentIDParams struct {
	ID        pgtype.UUID `json:"id"`
	PaymentID pgtype.Text `json:"payment_id"`
}

func (q *Queries) UpdateEscrowUserPaymentID(ctx context.Context, arg UpdateEscrowUserPaymentIDParams) error {
	_, err := q.db.Exec(ctx, updateEscrowUserPaymentID, arg.ID, arg.PaymentID)
	return err
}

const updateProjectBackingStatus = `-- name: UpdateProjectBackingStatus :exec
UPDATE backings
SET status = $2
WHERE project_id = $1
`

type UpdateProjectBackingStatusParams struct {
	ProjectID pgtype.UUID   `json:"project_id"`
	Status    BackingStatus `json:"status"`
}

func (q *Queries) UpdateProjectBackingStatus(ctx context.Context, arg UpdateProjectBackingStatusParams) error {
	_, err := q.db.Exec(ctx, updateProjectBackingStatus, arg.ProjectID, arg.Status)
	return err
}

const updateProjectByID = `-- name: UpdateProjectByID :exec
UPDATE projects
SET title = $2, description = $3, cover_picture = $4, goal_amount = $5, country = $6, province = $7, end_date = $8
WHERE id = $1
`

type UpdateProjectByIDParams struct {
	ID           pgtype.UUID        `json:"id"`
	Title        string             `json:"title"`
	Description  string             `json:"description"`
	CoverPicture string             `json:"cover_picture"`
	GoalAmount   int64              `json:"goal_amount"`
	Country      string             `json:"country"`
	Province     string             `json:"province"`
	EndDate      pgtype.Timestamptz `json:"end_date"`
}

func (q *Queries) UpdateProjectByID(ctx context.Context, arg UpdateProjectByIDParams) error {
	_, err := q.db.Exec(ctx, updateProjectByID,
		arg.ID,
		arg.Title,
		arg.Description,
		arg.CoverPicture,
		arg.GoalAmount,
		arg.Country,
		arg.Province,
		arg.EndDate,
	)
	return err
}

const updateProjectFund = `-- name: UpdateProjectFund :exec
UPDATE projects SET current_amount = current_amount + $2::bigint
WHERE id = $1
`

type UpdateProjectFundParams struct {
	ID            pgtype.UUID `json:"id"`
	BackingAmount int64       `json:"backing_amount"`
}

func (q *Queries) UpdateProjectFund(ctx context.Context, arg UpdateProjectFundParams) error {
	_, err := q.db.Exec(ctx, updateProjectFund, arg.ID, arg.BackingAmount)
	return err
}

const updateProjectPaymentID = `-- name: UpdateProjectPaymentID :exec
UPDATE projects SET payment_id = $2
WHERE id = $1
`

type UpdateProjectPaymentIDParams struct {
	ID        pgtype.UUID `json:"id"`
	PaymentID pgtype.Text `json:"payment_id"`
}

func (q *Queries) UpdateProjectPaymentID(ctx context.Context, arg UpdateProjectPaymentIDParams) error {
	_, err := q.db.Exec(ctx, updateProjectPaymentID, arg.ID, arg.PaymentID)
	return err
}

const updateTransactionStatus = `-- name: UpdateTransactionStatus :exec
UPDATE transactions
SET status = $2
WHERE id = $1
`

type UpdateTransactionStatusParams struct {
	ID     pgtype.UUID       `json:"id"`
	Status TransactionStatus `json:"status"`
}

func (q *Queries) UpdateTransactionStatus(ctx context.Context, arg UpdateTransactionStatusParams) error {
	_, err := q.db.Exec(ctx, updateTransactionStatus, arg.ID, arg.Status)
	return err
}

const updateUserByID = `-- name: UpdateUserByID :exec
UPDATE users
SET email = $2, first_name = $3, last_name = $4, profile_picture = $5, activated = $6
WHERE id = $1
`

type UpdateUserByIDParams struct {
	ID             pgtype.UUID `json:"id"`
	Email          string      `json:"email"`
	FirstName      string      `json:"first_name"`
	LastName       string      `json:"last_name"`
	ProfilePicture pgtype.Text `json:"profile_picture"`
	Activated      bool        `json:"activated"`
}

func (q *Queries) UpdateUserByID(ctx context.Context, arg UpdateUserByIDParams) error {
	_, err := q.db.Exec(ctx, updateUserByID,
		arg.ID,
		arg.Email,
		arg.FirstName,
		arg.LastName,
		arg.ProfilePicture,
		arg.Activated,
	)
	return err
}

const updateUserPassword = `-- name: UpdateUserPassword :exec
UPDATE users
SET hashed_password = $2
WHERE id = $1
`

type UpdateUserPasswordParams struct {
	ID             pgtype.UUID `json:"id"`
	HashedPassword string      `json:"hashed_password"`
}

func (q *Queries) UpdateUserPassword(ctx context.Context, arg UpdateUserPasswordParams) error {
	_, err := q.db.Exec(ctx, updateUserPassword, arg.ID, arg.HashedPassword)
	return err
}
