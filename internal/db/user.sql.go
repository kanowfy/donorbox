// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: user.sql

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

func (q *Queries) ActivateUser(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, activateUser, id)
	return err
}

const createSocialLoginUser = `-- name: CreateSocialLoginUser :one
INSERT INTO users (
    email, hashed_password, first_name, last_name, profile_picture, activated
) VALUES (
    $1, 'xxxxxxxx', $2, $3, $4, TRUE
)
RETURNING id, email, first_name, last_name, profile_picture, hashed_password, activated, verification_status, verification_document_url, created_at
`

type CreateSocialLoginUserParams struct {
	Email          string
	FirstName      string
	LastName       string
	ProfilePicture *string
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
		&i.FirstName,
		&i.LastName,
		&i.ProfilePicture,
		&i.HashedPassword,
		&i.Activated,
		&i.VerificationStatus,
		&i.VerificationDocumentUrl,
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
RETURNING id, email, first_name, last_name, profile_picture, hashed_password, activated, verification_status, verification_document_url, created_at
`

type CreateUserParams struct {
	Email          string
	HashedPassword string
	FirstName      string
	LastName       string
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
		&i.FirstName,
		&i.LastName,
		&i.ProfilePicture,
		&i.HashedPassword,
		&i.Activated,
		&i.VerificationStatus,
		&i.VerificationDocumentUrl,
		&i.CreatedAt,
	)
	return i, err
}

const getAllUsers = `-- name: GetAllUsers :many
SELECT id, email, first_name, last_name, profile_picture, activated, verification_status, created_at FROM users
`

type GetAllUsersRow struct {
	ID                 int64
	Email              string
	FirstName          string
	LastName           string
	ProfilePicture     *string
	Activated          bool
	VerificationStatus VerificationStatus
	CreatedAt          pgtype.Timestamptz
}

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
			&i.VerificationStatus,
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
SELECT id, email, first_name, last_name, profile_picture, hashed_password, activated, verification_status, verification_document_url, created_at FROM users
WHERE email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.FirstName,
		&i.LastName,
		&i.ProfilePicture,
		&i.HashedPassword,
		&i.Activated,
		&i.VerificationStatus,
		&i.VerificationDocumentUrl,
		&i.CreatedAt,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, email, first_name, last_name, profile_picture, hashed_password, activated, verification_status, verification_document_url, created_at FROM users
WHERE id = $1
`

func (q *Queries) GetUserByID(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRow(ctx, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.FirstName,
		&i.LastName,
		&i.ProfilePicture,
		&i.HashedPassword,
		&i.Activated,
		&i.VerificationStatus,
		&i.VerificationDocumentUrl,
		&i.CreatedAt,
	)
	return i, err
}

const updateUserByID = `-- name: UpdateUserByID :exec
UPDATE users
SET email = $2, first_name = $3, last_name = $4, profile_picture = $5
WHERE id = $1
`

type UpdateUserByIDParams struct {
	ID             int64
	Email          string
	FirstName      string
	LastName       string
	ProfilePicture *string
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
	ID             int64
	HashedPassword string
}

func (q *Queries) UpdateUserPassword(ctx context.Context, arg UpdateUserPasswordParams) error {
	_, err := q.db.Exec(ctx, updateUserPassword, arg.ID, arg.HashedPassword)
	return err
}

const updateVerificationStatus = `-- name: UpdateVerificationStatus :exec
UPDATE users
SET verification_status = $2, verification_document_url = $3
WHERE id = $1
`

type UpdateVerificationStatusParams struct {
	ID                      int64
	VerificationStatus      VerificationStatus
	VerificationDocumentUrl *string
}

func (q *Queries) UpdateVerificationStatus(ctx context.Context, arg UpdateVerificationStatusParams) error {
	_, err := q.db.Exec(ctx, updateVerificationStatus, arg.ID, arg.VerificationStatus, arg.VerificationDocumentUrl)
	return err
}
