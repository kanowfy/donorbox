-- name: GetAllProjects :many
SELECT projects.*, COUNT(backings.project_id) as backing_count
FROM projects
LEFT JOIN backings ON projects.ID = backings.project_id
WHERE category_id = 
    CASE WHEN @category::integer > 0 THEN @category::integer ELSE category_id END
AND projects.status = 'ongoing'
GROUP BY projects.ID
ORDER BY backing_count DESC
LIMIT @page_limit::integer OFFSET @total_offset::integer;

-- name: SearchProjects :many
SELECT projects.*, COUNT(backings.project_id) as backing_count
FROM projects
LEFT JOIN backings ON projects.ID = backings.project_id
WHERE 
    to_tsvector('english', title || ' ' || description || ' ' || province || ' ' || country) @@ plainto_tsquery('english', @search_query::text)
GROUP BY projects.ID
ORDER BY backing_count DESC
LIMIT @page_limit::integer OFFSET @total_offset::integer;

-- name: GetProjectByID :one
SELECT * FROM projects
WHERE id = $1;

-- name: GetProjectsForUser :many
SELECT * FROM projects
WHERE user_id = $1
ORDER BY start_date DESC;

-- name: GetEndedProjects :many
SELECT projects.*, COUNT(backings.project_id) as backing_count
FROM projects
JOIN backings ON projects.ID = backings.project_id
WHERE projects.status = 'ended'
GROUP BY projects.ID
ORDER BY end_date DESC;

-- name: UpdateProjectByID :exec
UPDATE projects
SET title = $2, description = $3, cover_picture = $4, goal_amount = $5, country = $6, province = $7, end_date = $8
WHERE id = $1;

-- name: UpdateProjectCard :exec
UPDATE projects SET card_id = $2
WHERE id = $1;

-- name: UpdateProjectFund :exec
UPDATE projects SET current_amount = current_amount + @backing_amount::bigint
WHERE id = $1;

-- name: UpdateProjectStatus :exec
UPDATE projects SET status = $2
WHERE id = $1;

-- name: DeleteProjectByID :exec
DELETE FROM projects WHERE id = $1;

-- name: CreateProject :one
INSERT INTO projects (
    user_id, category_id, title, description, cover_picture, goal_amount, country, province, end_date
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9
)
RETURNING *;

-- name: GetAllCategories :many
SELECT * FROM categories;

-- name: GetProjectUpdates :many
SELECT * FROM project_updates
WHERE project_id = $1
ORDER BY created_at DESC;

-- name: DeleteProjectUpdate :exec
DELETE FROM project_updates
WHERE id = $1;

-- name: CreateProjectUpdate :one
INSERT INTO project_updates (
    project_id, attachment_photo, description
) VALUES (
    $1, $2, $3
)
RETURNING *;

-- name: UpdateFinishedProjectsStatus :many
UPDATE projects SET status = 'ended'
WHERE end_date <= NOW() AND status = 'ongoing' RETURNING *;
