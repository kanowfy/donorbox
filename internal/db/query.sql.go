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
    project_id, backer_id, amount, word_of_support
) VALUES (
    $1, $2, $3, $4
) 
RETURNING id, project_id, backer_id, amount, word_of_support, status, created_at
`

type CreateBackingParams struct {
	ProjectID     pgtype.UUID `json:"project_id"`
	BackerID      pgtype.UUID `json:"backer_id"`
	Amount        int64       `json:"amount"`
	WordOfSupport *string     `json:"word_of_support"`
}

func (q *Queries) CreateBacking(ctx context.Context, arg CreateBackingParams) (Backing, error) {
	row := q.db.QueryRow(ctx, createBacking,
		arg.ProjectID,
		arg.BackerID,
		arg.Amount,
		arg.WordOfSupport,
	)
	var i Backing
	err := row.Scan(
		&i.ID,
		&i.ProjectID,
		&i.BackerID,
		&i.Amount,
		&i.WordOfSupport,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const createCard = `-- name: CreateCard :one
INSERT INTO cards (
    token, card_owner_name, last_four_digits, brand
) VALUES (
    $1, $2, $3, $4
)
RETURNING id, token, card_owner_name, last_four_digits, brand, created_at
`

type CreateCardParams struct {
	Token          string    `json:"token"`
	CardOwnerName  string    `json:"card_owner_name"`
	LastFourDigits string    `json:"last_four_digits"`
	Brand          CardBrand `json:"brand"`
}

func (q *Queries) CreateCard(ctx context.Context, arg CreateCardParams) (Card, error) {
	row := q.db.QueryRow(ctx, createCard,
		arg.Token,
		arg.CardOwnerName,
		arg.LastFourDigits,
		arg.Brand,
	)
	var i Card
	err := row.Scan(
		&i.ID,
		&i.Token,
		&i.CardOwnerName,
		&i.LastFourDigits,
		&i.Brand,
		&i.CreatedAt,
	)
	return i, err
}

const createProject = `-- name: CreateProject :one
INSERT INTO projects (
    user_id, category_id, title, description, cover_picture, goal_amount, country, province, end_date
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9
)
RETURNING id, user_id, category_id, title, description, cover_picture, goal_amount, current_amount, country, province, card_id, start_date, end_date, status
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
		&i.CardID,
		&i.StartDate,
		&i.EndDate,
		&i.Status,
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

const createSocialLoginUser = `-- name: CreateSocialLoginUser :one
INSERT INTO users (
    email, hashed_password, first_name, last_name, profile_picture, activated
) VALUES (
    $1, 'xxxxxxxx', $2, $3, $4, TRUE
)
RETURNING id, email, hashed_password, first_name, last_name, profile_picture, activated, user_type, created_at
`

type CreateSocialLoginUserParams struct {
	Email          string  `json:"email"`
	FirstName      string  `json:"first_name"`
	LastName       string  `json:"last_name"`
	ProfilePicture *string `json:"profile_picture"`
}

func (q *Queries) CreateSocialLoginUser(ctx context.Context, arg CreateSocialLoginUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createSocialLoginUser,
		arg.Email,
		arg.FirstName,
		arg.LastName,
		arg.ProfilePicture,
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

const createTransaction = `-- name: CreateTransaction :one
INSERT INTO transactions (
    project_id, transaction_type, amount, initiator_card_id, recipient_card_id
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING id, project_id, transaction_type, amount, initiator_card_id, recipient_card_id, status, created_at
`

type CreateTransactionParams struct {
	ProjectID       pgtype.UUID     `json:"project_id"`
	TransactionType TransactionType `json:"transaction_type"`
	Amount          int64           `json:"amount"`
	InitiatorCardID pgtype.UUID     `json:"initiator_card_id"`
	RecipientCardID pgtype.UUID     `json:"recipient_card_id"`
}

func (q *Queries) CreateTransaction(ctx context.Context, arg CreateTransactionParams) (Transaction, error) {
	row := q.db.QueryRow(ctx, createTransaction,
		arg.ProjectID,
		arg.TransactionType,
		arg.Amount,
		arg.InitiatorCardID,
		arg.RecipientCardID,
	)
	var i Transaction
	err := row.Scan(
		&i.ID,
		&i.ProjectID,
		&i.TransactionType,
		&i.Amount,
		&i.InitiatorCardID,
		&i.RecipientCardID,
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

const deleteProjectUpdate = `-- name: DeleteProjectUpdate :exec
DELETE FROM project_updates
WHERE id = $1
`

func (q *Queries) DeleteProjectUpdate(ctx context.Context, id pgtype.UUID) error {
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

SELECT projects.id, projects.user_id, projects.category_id, projects.title, projects.description, projects.cover_picture, projects.goal_amount, projects.current_amount, projects.country, projects.province, projects.card_id, projects.start_date, projects.end_date, projects.status, COUNT(backings.project_id) as backing_count
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
	CardID        pgtype.UUID        `json:"card_id"`
	StartDate     pgtype.Timestamptz `json:"start_date"`
	EndDate       pgtype.Timestamptz `json:"end_date"`
	Status        ProjectStatus      `json:"status"`
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
			&i.CardID,
			&i.StartDate,
			&i.EndDate,
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

const getAllTransactions = `-- name: GetAllTransactions :many

SELECT id, project_id, transaction_type, amount, initiator_card_id, recipient_card_id, status, created_at FROM transactions
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
			&i.ProjectID,
			&i.TransactionType,
			&i.Amount,
			&i.InitiatorCardID,
			&i.RecipientCardID,
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

SELECT id, email, first_name, last_name, profile_picture, created_at FROM users
`

type GetAllUsersRow struct {
	ID             pgtype.UUID        `json:"id"`
	Email          string             `json:"email"`
	FirstName      string             `json:"first_name"`
	LastName       string             `json:"last_name"`
	ProfilePicture *string            `json:"profile_picture"`
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
SELECT id, project_id, backer_id, amount, word_of_support, status, created_at FROM backings
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
		&i.WordOfSupport,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const getBackingCountForProject = `-- name: GetBackingCountForProject :one
SELECT COUNT(*) AS backing_count
FROM backings
WHERE project_id = $1
`

func (q *Queries) GetBackingCountForProject(ctx context.Context, projectID pgtype.UUID) (int64, error) {
	row := q.db.QueryRow(ctx, getBackingCountForProject, projectID)
	var backing_count int64
	err := row.Scan(&backing_count)
	return backing_count, err
}

const getBackingTransactionsForProject = `-- name: GetBackingTransactionsForProject :many
SELECT id, project_id, transaction_type, amount, initiator_card_id, recipient_card_id, status, created_at FROM transactions
WHERE project_id = $1 AND transaction_type == 'backing'
`

func (q *Queries) GetBackingTransactionsForProject(ctx context.Context, projectID pgtype.UUID) ([]Transaction, error) {
	rows, err := q.db.Query(ctx, getBackingTransactionsForProject, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Transaction
	for rows.Next() {
		var i Transaction
		if err := rows.Scan(
			&i.ID,
			&i.ProjectID,
			&i.TransactionType,
			&i.Amount,
			&i.InitiatorCardID,
			&i.RecipientCardID,
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

const getBackingsForProject = `-- name: GetBackingsForProject :many

SELECT backings.id, backings.project_id, backings.backer_id, backings.amount, backings.word_of_support, backings.status, backings.created_at, users.first_name, users.last_name, users.profile_picture
FROM backings
LEFT JOIN users
ON backings.backer_id = users.id
WHERE project_id = $1
`

type GetBackingsForProjectRow struct {
	ID             pgtype.UUID        `json:"id"`
	ProjectID      pgtype.UUID        `json:"project_id"`
	BackerID       pgtype.UUID        `json:"backer_id"`
	Amount         int64              `json:"amount"`
	WordOfSupport  *string            `json:"word_of_support"`
	Status         BackingStatus      `json:"status"`
	CreatedAt      pgtype.Timestamptz `json:"created_at"`
	FirstName      *string            `json:"first_name"`
	LastName       *string            `json:"last_name"`
	ProfilePicture *string            `json:"profile_picture"`
}

// :::::::::: BACKING ::::::::::--
func (q *Queries) GetBackingsForProject(ctx context.Context, projectID pgtype.UUID) ([]GetBackingsForProjectRow, error) {
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
			&i.ProjectID,
			&i.BackerID,
			&i.Amount,
			&i.WordOfSupport,
			&i.Status,
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
SELECT id, project_id, backer_id, amount, word_of_support, status, created_at FROM backings
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
			&i.WordOfSupport,
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

const getCardByID = `-- name: GetCardByID :one

SELECT id, token, card_owner_name, last_four_digits, brand, created_at FROM cards
WHERE id = $1
`

// :::::::::: CARD ::::::::::--
func (q *Queries) GetCardByID(ctx context.Context, id pgtype.UUID) (Card, error) {
	row := q.db.QueryRow(ctx, getCardByID, id)
	var i Card
	err := row.Scan(
		&i.ID,
		&i.Token,
		&i.CardOwnerName,
		&i.LastFourDigits,
		&i.Brand,
		&i.CreatedAt,
	)
	return i, err
}

const getEscrowUser = `-- name: GetEscrowUser :one
SELECT id, email, hashed_password, user_type, card_id, created_at FROM escrow_users
LIMIT 1
`

func (q *Queries) GetEscrowUser(ctx context.Context) (EscrowUser, error) {
	row := q.db.QueryRow(ctx, getEscrowUser)
	var i EscrowUser
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.HashedPassword,
		&i.UserType,
		&i.CardID,
		&i.CreatedAt,
	)
	return i, err
}

const getEscrowUserByEmail = `-- name: GetEscrowUserByEmail :one
SELECT id, email, hashed_password, user_type, card_id, created_at FROM escrow_users
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
		&i.CardID,
		&i.CreatedAt,
	)
	return i, err
}

const getEscrowUserByID = `-- name: GetEscrowUserByID :one
SELECT id, email, hashed_password, user_type, card_id, created_at FROM escrow_users
WHERE id = $1
`

// :::::::::: ESCROW USER ::::::::::--
func (q *Queries) GetEscrowUserByID(ctx context.Context, id pgtype.UUID) (EscrowUser, error) {
	row := q.db.QueryRow(ctx, getEscrowUserByID, id)
	var i EscrowUser
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.HashedPassword,
		&i.UserType,
		&i.CardID,
		&i.CreatedAt,
	)
	return i, err
}

const getFirstBackingDonor = `-- name: GetFirstBackingDonor :one
SELECT backings.id, backings.project_id, backings.backer_id, backings.amount, backings.word_of_support, backings.status, backings.created_at, users.first_name, users.last_name, users.profile_picture
FROM backings
LEFT JOIN users
ON backings.backer_id = users.id
WHERE project_id = $1
ORDER BY backings.created_at
LIMIT 1
`

type GetFirstBackingDonorRow struct {
	ID             pgtype.UUID        `json:"id"`
	ProjectID      pgtype.UUID        `json:"project_id"`
	BackerID       pgtype.UUID        `json:"backer_id"`
	Amount         int64              `json:"amount"`
	WordOfSupport  *string            `json:"word_of_support"`
	Status         BackingStatus      `json:"status"`
	CreatedAt      pgtype.Timestamptz `json:"created_at"`
	FirstName      *string            `json:"first_name"`
	LastName       *string            `json:"last_name"`
	ProfilePicture *string            `json:"profile_picture"`
}

func (q *Queries) GetFirstBackingDonor(ctx context.Context, projectID pgtype.UUID) (GetFirstBackingDonorRow, error) {
	row := q.db.QueryRow(ctx, getFirstBackingDonor, projectID)
	var i GetFirstBackingDonorRow
	err := row.Scan(
		&i.ID,
		&i.ProjectID,
		&i.BackerID,
		&i.Amount,
		&i.WordOfSupport,
		&i.Status,
		&i.CreatedAt,
		&i.FirstName,
		&i.LastName,
		&i.ProfilePicture,
	)
	return i, err
}

const getMostBackingDonor = `-- name: GetMostBackingDonor :one
SELECT backings.id, backings.project_id, backings.backer_id, backings.amount, backings.word_of_support, backings.status, backings.created_at, users.first_name, users.last_name, users.profile_picture
FROM backings
LEFT JOIN users
ON backings.backer_id = users.id
WHERE project_id = $1
ORDER BY backings.amount DESC
LIMIT 1
`

type GetMostBackingDonorRow struct {
	ID             pgtype.UUID        `json:"id"`
	ProjectID      pgtype.UUID        `json:"project_id"`
	BackerID       pgtype.UUID        `json:"backer_id"`
	Amount         int64              `json:"amount"`
	WordOfSupport  *string            `json:"word_of_support"`
	Status         BackingStatus      `json:"status"`
	CreatedAt      pgtype.Timestamptz `json:"created_at"`
	FirstName      *string            `json:"first_name"`
	LastName       *string            `json:"last_name"`
	ProfilePicture *string            `json:"profile_picture"`
}

func (q *Queries) GetMostBackingDonor(ctx context.Context, projectID pgtype.UUID) (GetMostBackingDonorRow, error) {
	row := q.db.QueryRow(ctx, getMostBackingDonor, projectID)
	var i GetMostBackingDonorRow
	err := row.Scan(
		&i.ID,
		&i.ProjectID,
		&i.BackerID,
		&i.Amount,
		&i.WordOfSupport,
		&i.Status,
		&i.CreatedAt,
		&i.FirstName,
		&i.LastName,
		&i.ProfilePicture,
	)
	return i, err
}

const getMostRecentBackingDonor = `-- name: GetMostRecentBackingDonor :one
SELECT backings.id, backings.project_id, backings.backer_id, backings.amount, backings.word_of_support, backings.status, backings.created_at, users.first_name, users.last_name, users.profile_picture
FROM backings
LEFT JOIN users
ON backings.backer_id = users.id
WHERE project_id = $1
ORDER BY backings.created_at DESC
LIMIT 1
`

type GetMostRecentBackingDonorRow struct {
	ID             pgtype.UUID        `json:"id"`
	ProjectID      pgtype.UUID        `json:"project_id"`
	BackerID       pgtype.UUID        `json:"backer_id"`
	Amount         int64              `json:"amount"`
	WordOfSupport  *string            `json:"word_of_support"`
	Status         BackingStatus      `json:"status"`
	CreatedAt      pgtype.Timestamptz `json:"created_at"`
	FirstName      *string            `json:"first_name"`
	LastName       *string            `json:"last_name"`
	ProfilePicture *string            `json:"profile_picture"`
}

func (q *Queries) GetMostRecentBackingDonor(ctx context.Context, projectID pgtype.UUID) (GetMostRecentBackingDonorRow, error) {
	row := q.db.QueryRow(ctx, getMostRecentBackingDonor, projectID)
	var i GetMostRecentBackingDonorRow
	err := row.Scan(
		&i.ID,
		&i.ProjectID,
		&i.BackerID,
		&i.Amount,
		&i.WordOfSupport,
		&i.Status,
		&i.CreatedAt,
		&i.FirstName,
		&i.LastName,
		&i.ProfilePicture,
	)
	return i, err
}

const getProjectByID = `-- name: GetProjectByID :one
SELECT id, user_id, category_id, title, description, cover_picture, goal_amount, current_amount, country, province, card_id, start_date, end_date, status FROM projects
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
		&i.CardID,
		&i.StartDate,
		&i.EndDate,
		&i.Status,
	)
	return i, err
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

const getTransactionByID = `-- name: GetTransactionByID :one
SELECT id, project_id, transaction_type, amount, initiator_card_id, recipient_card_id, status, created_at FROM transactions
WHERE id = $1
`

func (q *Queries) GetTransactionByID(ctx context.Context, id pgtype.UUID) (Transaction, error) {
	row := q.db.QueryRow(ctx, getTransactionByID, id)
	var i Transaction
	err := row.Scan(
		&i.ID,
		&i.ProjectID,
		&i.TransactionType,
		&i.Amount,
		&i.InitiatorCardID,
		&i.RecipientCardID,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const getTransactionsForProject = `-- name: GetTransactionsForProject :many
SELECT id, project_id, transaction_type, amount, initiator_card_id, recipient_card_id, status, created_at FROM transactions
WHERE project_id = $1 ORDER BY created_at ASC
`

func (q *Queries) GetTransactionsForProject(ctx context.Context, projectID pgtype.UUID) ([]Transaction, error) {
	rows, err := q.db.Query(ctx, getTransactionsForProject, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Transaction
	for rows.Next() {
		var i Transaction
		if err := rows.Scan(
			&i.ID,
			&i.ProjectID,
			&i.TransactionType,
			&i.Amount,
			&i.InitiatorCardID,
			&i.RecipientCardID,
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
SELECT projects.id, projects.user_id, projects.category_id, projects.title, projects.description, projects.cover_picture, projects.goal_amount, projects.current_amount, projects.country, projects.province, projects.card_id, projects.start_date, projects.end_date, projects.status, COUNT(backings.project_id) as backing_count
FROM projects
LEFT JOIN backings ON projects.ID = backings.project_id
WHERE 
    to_tsvector('english', title || ' ' || description || ' ' || province || ' ' || country) @@ plainto_tsquery('english', $1::text)
GROUP BY projects.ID
ORDER BY backing_count DESC
LIMIT $3::integer OFFSET $2::integer
`

type SearchProjectsParams struct {
	SearchQuery string `json:"search_query"`
	TotalOffset int32  `json:"total_offset"`
	PageLimit   int32  `json:"page_limit"`
}

type SearchProjectsRow struct {
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
	CardID        pgtype.UUID        `json:"card_id"`
	StartDate     pgtype.Timestamptz `json:"start_date"`
	EndDate       pgtype.Timestamptz `json:"end_date"`
	Status        ProjectStatus      `json:"status"`
	BackingCount  int64              `json:"backing_count"`
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
			&i.CategoryID,
			&i.Title,
			&i.Description,
			&i.CoverPicture,
			&i.GoalAmount,
			&i.CurrentAmount,
			&i.Country,
			&i.Province,
			&i.CardID,
			&i.StartDate,
			&i.EndDate,
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

const updateEscrowCard = `-- name: UpdateEscrowCard :exec
UPDATE escrow_users SET card_id = $2
WHERE id = $1
`

type UpdateEscrowCardParams struct {
	ID     pgtype.UUID `json:"id"`
	CardID pgtype.UUID `json:"card_id"`
}

func (q *Queries) UpdateEscrowCard(ctx context.Context, arg UpdateEscrowCardParams) error {
	_, err := q.db.Exec(ctx, updateEscrowCard, arg.ID, arg.CardID)
	return err
}

const updateEscrowUserByID = `-- name: UpdateEscrowUserByID :exec
UPDATE escrow_users SET email = $2, hashed_password = $3
WHERE id = $1
`

type UpdateEscrowUserByIDParams struct {
	ID             pgtype.UUID `json:"id"`
	Email          string      `json:"email"`
	HashedPassword string      `json:"hashed_password"`
}

func (q *Queries) UpdateEscrowUserByID(ctx context.Context, arg UpdateEscrowUserByIDParams) error {
	_, err := q.db.Exec(ctx, updateEscrowUserByID, arg.ID, arg.Email, arg.HashedPassword)
	return err
}

const updateFinishedProjectsStatus = `-- name: UpdateFinishedProjectsStatus :exec
UPDATE projects SET status = 'ended'
WHERE end_date <= NOW() AND status = 'ongoing'
`

func (q *Queries) UpdateFinishedProjectsStatus(ctx context.Context) error {
	_, err := q.db.Exec(ctx, updateFinishedProjectsStatus)
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

const updateProjectCard = `-- name: UpdateProjectCard :exec
UPDATE projects SET card_id = $2
WHERE id = $1
`

type UpdateProjectCardParams struct {
	ID     pgtype.UUID `json:"id"`
	CardID pgtype.UUID `json:"card_id"`
}

func (q *Queries) UpdateProjectCard(ctx context.Context, arg UpdateProjectCardParams) error {
	_, err := q.db.Exec(ctx, updateProjectCard, arg.ID, arg.CardID)
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

const updateProjectStatus = `-- name: UpdateProjectStatus :exec
UPDATE projects SET status = $2
WHERE id = $1
`

type UpdateProjectStatusParams struct {
	ID     pgtype.UUID   `json:"id"`
	Status ProjectStatus `json:"status"`
}

func (q *Queries) UpdateProjectStatus(ctx context.Context, arg UpdateProjectStatusParams) error {
	_, err := q.db.Exec(ctx, updateProjectStatus, arg.ID, arg.Status)
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
SET email = $2, first_name = $3, last_name = $4, profile_picture = $5
WHERE id = $1
`

type UpdateUserByIDParams struct {
	ID             pgtype.UUID `json:"id"`
	Email          string      `json:"email"`
	FirstName      string      `json:"first_name"`
	LastName       string      `json:"last_name"`
	ProfilePicture *string     `json:"profile_picture"`
}

func (q *Queries) UpdateUserByID(ctx context.Context, arg UpdateUserByIDParams) error {
	_, err := q.db.Exec(ctx, updateUserByID,
		arg.ID,
		arg.Email,
		arg.FirstName,
		arg.LastName,
		arg.ProfilePicture,
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
