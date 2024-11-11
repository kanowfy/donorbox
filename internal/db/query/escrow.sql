-- name: GetEscrowUserByID :one
SELECT * FROM escrow_users
WHERE id = $1;

-- name: GetEscrowUserByEmail :one
SELECT * FROM escrow_users
WHERE email = $1;

-- name: GetEscrowUser :one
SELECT * FROM escrow_users
LIMIT 1;

-- name: UpdateEscrowUserByID :exec
UPDATE escrow_users SET email = $2, hashed_password = $3
WHERE id = $1;
