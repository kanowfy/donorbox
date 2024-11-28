// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: escrow.sql

package db

import (
	"context"
)

const createEscrowUser = `-- name: CreateEscrowUser :one
INSERT INTO escrow_users (email, hashed_password)
VALUES ($1, $2)
RETURNING id, email, hashed_password, created_at
`

type CreateEscrowUserParams struct {
	Email          string
	HashedPassword string
}

func (q *Queries) CreateEscrowUser(ctx context.Context, arg CreateEscrowUserParams) (EscrowUser, error) {
	row := q.db.QueryRow(ctx, createEscrowUser, arg.Email, arg.HashedPassword)
	var i EscrowUser
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.HashedPassword,
		&i.CreatedAt,
	)
	return i, err
}

const getEscrowUserByEmail = `-- name: GetEscrowUserByEmail :one
SELECT id, email, hashed_password, created_at FROM escrow_users
WHERE email = $1
`

func (q *Queries) GetEscrowUserByEmail(ctx context.Context, email string) (EscrowUser, error) {
	row := q.db.QueryRow(ctx, getEscrowUserByEmail, email)
	var i EscrowUser
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.HashedPassword,
		&i.CreatedAt,
	)
	return i, err
}

const getEscrowUserByID = `-- name: GetEscrowUserByID :one
SELECT id, email, hashed_password, created_at FROM escrow_users
WHERE id = $1
`

func (q *Queries) GetEscrowUserByID(ctx context.Context, id int64) (EscrowUser, error) {
	row := q.db.QueryRow(ctx, getEscrowUserByID, id)
	var i EscrowUser
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.HashedPassword,
		&i.CreatedAt,
	)
	return i, err
}

const updateEscrowUserByID = `-- name: UpdateEscrowUserByID :exec
UPDATE escrow_users SET email = $2, hashed_password = $3
WHERE id = $1
`

type UpdateEscrowUserByIDParams struct {
	ID             int64
	Email          string
	HashedPassword string
}

func (q *Queries) UpdateEscrowUserByID(ctx context.Context, arg UpdateEscrowUserByIDParams) error {
	_, err := q.db.Exec(ctx, updateEscrowUserByID, arg.ID, arg.Email, arg.HashedPassword)
	return err
}
