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
AND project.status = 'ongoing'
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

-- name: GetFinishedProjects :many
SELECT projects.*, COUNT(backings.project_id) as backing_count
FROM projects
JOIN backings ON projects.ID = backings.project_id
WHERE projects.status = 'finished'
GROUP BY projects.ID
ORDER BY end_date DESC;

-- name: GetMilestoneByID :one
SELECT * FROM milestones
WHERE id = $1;

-- name: GetMilestoneForProject :many
SELECT * FROM milestones
WHERE project_id = $1;

-- name: UpdateProjectByID :exec
UPDATE projects
SET title = $2, description = $3, cover_picture = $4, receiver_number=$5, receiver_name=$6, address=$7, district=$8, city=$9, country = $10, end_date = $11
WHERE id = $1;

-- name: UpdateProjectStatus :exec
UPDATE projects SET status = $2
WHERE id = $1;

-- name: DeleteProjectByID :exec
DELETE FROM projects WHERE id = $1;

-- name: CreateProject :one
INSERT INTO projects (
    user_id, category_id, title, description, cover_picture, end_date, receiver_number, receiver_name, address, district, city, country
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
)
RETURNING *;

-- name: CreateMilestone :one
INSERT INTO milestones (
    project_id, title, description, fund_goal, bank_description
) VALUES (
    $1, $2, $3, $4, $5
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
