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

-- name: GetStatsAggregation :one
SELECT COALESCE(
	(SELECT COUNT(*)
	FROM projects GROUP BY status HAVING status = 'pending'), 0
)::bigint AS projects_pending, COALESCE(
	(SELECT COUNT(*)
	FROM projects GROUP BY status HAVING status = 'ongoing'), 0
)::bigint AS projects_ongoing, COALESCE(
	(SELECT COUNT(*)
	FROM projects GROUP BY status HAVING status = 'finished'), 0
)::bigint AS projects_finished, COALESCE(
	(SELECT COUNT(*)
	FROM projects GROUP BY status HAVING status = 'rejected'), 0
)::bigint AS projects_rejected, COALESCE(
        (SELECT SUM(amount)
	FROM backings), 0
)::bigint AS total_fund, (
        SELECT COUNT(*)
	FROM backings
) AS backing_count, COALESCE(
        (SELECT COUNT(*)
        FROM users GROUP BY verification_status HAVING verification_status='pending'), 0
)::bigint AS verification_count;

-- name: GetCategoriesCount :many
SELECT c.id, c.name, COALESCE(COUNT(*), 0)::bigint AS count
FROM categories c
JOIN projects p ON c.id = p.category_id
GROUP BY c.id
ORDER BY c.id;