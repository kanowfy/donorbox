-- name: GetAllUsers :many
SELECT id, email, first_name, last_name, profile_picture, activated, verification_status, created_at FROM users;

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

-- name: UpdateVerificationStatus :exec
UPDATE users
SET verification_status = $2, verification_document_url = $3
WHERE id = $1;

-- name: GetPendingVerificationUsers :many
SELECT id, email, first_name, last_name, verification_document_url, created_at
FROM users
WHERE verification_status = 'pending'
ORDER BY created_at DESC;