-- name: CreateReport :one
INSERT INTO reports (
    project_id, reporter_email, reporter_phone_number, reason, details
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;
