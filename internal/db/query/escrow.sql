-- name: GetEscrowUserByID :one
SELECT * FROM escrow_users
WHERE id = $1;

-- name: GetEscrowUserByEmail :one
SELECT * FROM escrow_users
WHERE email = $1;

-- name: CreateEscrowUser :one
INSERT INTO escrow_users (email, hashed_password)
VALUES ($1, $2)
RETURNING *;

-- name: UpdateEscrowUserByID :exec
UPDATE escrow_users SET email = $2, hashed_password = $3
WHERE id = $1;
