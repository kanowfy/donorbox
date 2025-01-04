-- name: GetBackingsForProject :many
SELECT backings.*, users.first_name, users.last_name, users.profile_picture
FROM backings
LEFT JOIN users
ON backings.user_id = users.id
WHERE project_id = $1
ORDER BY backings.created_at DESC;

-- name: GetBackingByID :one
SELECT * FROM backings
WHERE id = $1;

-- name: GetBackingsForUser :many
SELECT * FROM backings
WHERE user_id = $1;

-- name: GetMostBackingDonor :one
SELECT backings.*, users.first_name, users.last_name, users.profile_picture
FROM backings
LEFT JOIN users
ON backings.user_id = users.id
WHERE project_id = $1
ORDER BY backings.amount DESC
LIMIT 1;

-- name: GetFirstBackingDonor :one
SELECT backings.*, users.first_name, users.last_name, users.profile_picture
FROM backings
LEFT JOIN users
ON backings.user_id = users.id
WHERE project_id = $1
ORDER BY backings.created_at
LIMIT 1;

-- name: GetMostRecentBackingDonor :one
SELECT backings.*, users.first_name, users.last_name, users.profile_picture
FROM backings
LEFT JOIN users
ON backings.user_id = users.id
WHERE project_id = $1
ORDER BY backings.created_at DESC
LIMIT 1;

-- name: GetBackingCountForProject :one
SELECT COUNT(*) AS backing_count
FROM backings
WHERE project_id = $1;

-- name: CreateBacking :one
INSERT INTO backings (
    project_id, user_id, amount, word_of_support
) VALUES (
    $1, $2, $3, $4
) 
RETURNING *;

-- name: GetTotalBackingByMonth :many
WITH months AS (
    SELECT generate_series(
            date_trunc('month', (SELECT MIN(created_at) FROM backings)), 
            date_trunc('month', current_date), 
            interval '1 month') AS month
)
SELECT 
    to_char(months.month, 'YYYY-MM') AS month,
    COALESCE(SUM(backings.amount), 0)::bigint AS total_donated
FROM months
LEFT JOIN backings ON date_trunc('month', backings.created_at) = months.month
GROUP BY months.month
ORDER BY months.month;