-- name: GetBackingsForProject :many
SELECT backings.*, users.first_name, users.last_name, users.profile_picture
FROM backings
LEFT JOIN users
ON backings.backer_id = users.id
WHERE project_id = $1
ORDER BY backings.created_at DESC;

-- name: GetBackingByID :one
SELECT * FROM backings
WHERE id = $1;

-- name: GetBackingsForUser :many
SELECT * FROM backings
WHERE backer_id = $1;

-- name: GetMostBackingDonor :one
SELECT backings.*, users.first_name, users.last_name, users.profile_picture
FROM backings
LEFT JOIN users
ON backings.backer_id = users.id
WHERE project_id = $1
ORDER BY backings.amount DESC
LIMIT 1;

-- name: GetFirstBackingDonor :one
SELECT backings.*, users.first_name, users.last_name, users.profile_picture
FROM backings
LEFT JOIN users
ON backings.backer_id = users.id
WHERE project_id = $1
ORDER BY backings.created_at
LIMIT 1;

-- name: GetMostRecentBackingDonor :one
SELECT backings.*, users.first_name, users.last_name, users.profile_picture
FROM backings
LEFT JOIN users
ON backings.backer_id = users.id
WHERE project_id = $1
ORDER BY backings.created_at DESC
LIMIT 1;

-- name: GetBackingCountForProject :one
SELECT COUNT(*) AS backing_count
FROM backings
WHERE project_id = $1;

-- name: CreateBacking :one
INSERT INTO backings (
    project_id, backer_id, amount, word_of_support
) VALUES (
    $1, $2, $3, $4
) 
RETURNING *;

-- name: UpdateProjectBackingStatus :exec
UPDATE backings
SET status = $2
WHERE project_id = $1;
