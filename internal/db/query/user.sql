-- name: GetAllUsers :many
SELECT id, email, first_name, last_name, profile_picture, created_at FROM users;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1;

-- name: UpdateUserByID :exec
UPDATE users
SET email = $2, first_name = $3, last_name = $4, profile_picture = $5
WHERE id = $1;

-- name: UpdateUserPassword :exec
UPDATE users
SET hashed_password = $2
WHERE id = $1;

-- name: ActivateUser :exec
UPDATE users
SET activated = TRUE
WHERE id = $1;
    
-- name: CreateUser :one
INSERT INTO users (
    email, hashed_password, first_name, last_name 
) VALUES (
    $1, $2, $3, $4
)
RETURNING *;

-- name: CreateSocialLoginUser :one
INSERT INTO users (
    email, hashed_password, first_name, last_name, profile_picture, activated
) VALUES (
    $1, 'xxxxxxxx', $2, $3, $4, TRUE
)
RETURNING *;
