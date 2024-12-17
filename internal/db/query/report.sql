-- name: CreateProjectReport :one
INSERT INTO project_reports (
    project_id, email, full_name, phone_number, relation, reason, details
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
)
RETURNING *;

-- name: GetAllProjectReports :many
SELECT * FROM project_reports
ORDER BY created_at;

-- name: UpdateProjectReportStatus :exec
UPDATE project_reports
SET status = $2
WHERE id = $1;

-- name: GetProjectReportByID :one
SELECT * FROM project_reports
WHERE id = $1;

-- name: GetResolvedProjectReportsForProject :many
SELECT * FROM project_reports
WHERE project_id = $1 AND status = 'resolved';