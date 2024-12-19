-- name: GetAllProjects :many
WITH aggregated_backings AS (
    SELECT project_id, COUNT(*) AS backing_count
    FROM backings
    GROUP BY project_id
),
aggregated_milestones AS (
    SELECT project_id, 
           SUM(current_fund) AS total_fund, 
           SUM(fund_goal) AS fund_goal
    FROM milestones
    GROUP BY project_id
)
SELECT p.*, 
       COALESCE(m.total_fund, 0) AS total_fund,
       COALESCE(m.fund_goal, 0) AS fund_goal,
       COALESCE(b.backing_count, 0) AS backing_count
FROM projects p
LEFT JOIN aggregated_backings b ON p.id = b.project_id
LEFT JOIN aggregated_milestones m ON p.id = m.project_id
WHERE p.category_id =
    CASE WHEN @category::integer > 0 THEN @category::integer ELSE p.category_id END
ORDER BY backing_count DESC;

-- name: SearchProjects :many
WITH aggregated_backings AS (
    SELECT project_id, COUNT(*) AS backing_count
    FROM backings
    GROUP BY project_id
),
aggregated_milestones AS (
    SELECT project_id, 
           SUM(current_fund) AS total_fund, 
           SUM(fund_goal) AS fund_goal
    FROM milestones
    GROUP BY project_id
)
SELECT p.*, 
       COALESCE(m.total_fund, 0) AS total_fund,
       COALESCE(m.fund_goal, 0) AS fund_goal,
       COALESCE(b.backing_count, 0) AS backing_count
FROM projects p
LEFT JOIN aggregated_backings b ON p.id = b.project_id
LEFT JOIN aggregated_milestones m ON p.id = m.project_id
WHERE
    to_tsvector('english', p.title || ' ' || p.description || ' ' || city || ' ' || country) @@ plainto_tsquery('english', @search_query::text)
ORDER BY backing_count DESC;

-- name: GetProjectByID :one
SELECT projects.*, SUM(milestones.current_fund) AS total_fund, SUM(milestones.fund_goal) AS fund_goal
FROM projects
LEFT JOIN milestones ON projects.ID = milestones.project_id
WHERE projects.ID = $1
GROUP BY projects.ID;

-- name: GetProjectsForUser :many
SELECT projects.*, SUM(milestones.current_fund) AS total_fund, SUM(milestones.fund_goal) AS fund_goal
FROM projects
LEFT JOIN milestones ON projects.ID = milestones.project_id
WHERE user_id = $1
GROUP BY projects.ID
ORDER BY projects.created_at DESC;

-- name: GetFinishedProjects :many
SELECT projects.*, SUM(milestones.current_fund) AS total_fund,
SUM(milestones.fund_goal) AS fund_goal, COUNT(backings.project_id) as backing_count
FROM projects
JOIN backings ON projects.ID = backings.project_id
LEFT JOIN milestones ON projects.ID = milestones.project_id
WHERE projects.status = 'finished'
GROUP BY projects.ID
ORDER BY end_date DESC;

-- name: GetPendingProjects :many
SELECT projects.*, SUM(milestones.fund_goal) AS fund_goal
FROM projects
LEFT JOIN milestones ON projects.ID = milestones.project_id
WHERE projects.status = 'pending'
GROUP BY projects.ID
ORDER BY projects.created_at DESC;

-- name: GetMilestoneForProject :many
SELECT m.*, c.transfer_amount, c.transfer_note AS fund_released_note, c.transfer_image AS fund_released_image,
c.created_at AS fund_released_at
FROM milestones m
LEFT JOIN escrow_milestone_completions c ON m.id = c.milestone_id
WHERE m.project_id = $1
ORDER BY m.id;

-- name: UpdateProjectByID :one
UPDATE projects
SET title = $2, description = $3, cover_picture = $4, receiver_number=$5, receiver_name=$6, address=$7, district=$8, city=$9, country = $10, end_date = $11
WHERE id = $1
RETURNING *;

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

-- name: GetCategoryByName :one
SELECT * FROM categories
WHERE name = $1;

-- name: GetMilestoneByID :one
SELECT m.*, c.transfer_amount, c.transfer_note AS fund_released_note, c.transfer_image AS fund_released_image,
c.created_at AS fund_released_at
FROM milestones m
LEFT JOIN escrow_milestone_completions c ON m.id = c.milestone_id
WHERE m.id = $1;

-- name: UpdateMilestoneFund :exec
UPDATE milestones
SET current_fund = current_fund + @amount::bigint
WHERE id = $1;

-- name: UpdateMilestoneStatus :exec
UPDATE milestones
SET status = $2
WHERE id = $1;

-- name: GetFundedMilestones :many
SELECT m.*, c.transfer_amount, c.transfer_note AS fund_released_note, c.transfer_image AS fund_released_image, c.created_at AS fund_released_at,
p.address, p.district, p.city, p.country, p.receiver_name, p.receiver_number
FROM milestones m
JOIN projects p ON m.project_id = p.id
LEFT JOIN escrow_milestone_completions c ON m.id = c.milestone_id
WHERE current_fund >= fund_goal
ORDER BY m.id;

-- name: GetAllMilestones :many
SELECT m.*, c.transfer_amount, c.transfer_note AS fund_released_note, c.transfer_image AS fund_released_image, c.created_at AS fund_released_at,
p.address, p.district, p.city, p.country, p.receiver_name, p.receiver_number
FROM milestones m
LEFT JOIN escrow_milestone_completions c ON m.id = c.milestone_id
JOIN projects p ON m.project_id = p.id
ORDER BY m.id;

-- name: CreateMilestoneCompletion :one
INSERT INTO escrow_milestone_completions (
    milestone_id, transfer_amount, transfer_note, transfer_image
) VALUES (
    $1, $2, $3, $4
)
RETURNING *;

-- name: GetMilestoneCompletionByMilestoneID :one
SELECT * FROM escrow_milestone_completions
WHERE milestone_id = $1;

-- name: CreateSpendingProof :one
INSERT INTO user_spending_proofs (
    milestone_id, transfer_image, proof_media_url, description
) VALUES (
    $1, $2, $3, $4
) 
RETURNING *;

-- name: UpdateSpendingProofStatus :exec
UPDATE user_spending_proofs
SET status = $2, rejected_cause = $3
WHERE id = $1;

-- name: GetSpendingProofByID :one
SELECT * FROM user_spending_proofs
WHERE id = $1;

-- name: GetSpendingProofsForMilestone :many
SELECT * FROM user_spending_proofs
WHERE milestone_id = $1
ORDER BY created_at DESC;

-- name: GetMilestoneAndProofs :many
SELECT p.*, m.title AS milestone_title, m.description AS milestone_description, m.fund_goal, m.current_fund, 
m.bank_description, m.status AS milestone_status, m.created_at AS milestone_created_at 
FROM user_spending_proofs p
JOIN milestones m ON m.id = p.milestone_id
ORDER BY p.created_at;

-- name: GetDisputedProjects :many
SELECT p.*, SUM(m.current_fund) AS total_fund,
SUM(m.fund_goal) AS fund_goal, COUNT(b.project_id) as backing_count
FROM projects p
LEFT JOIN backings b ON p.ID = b.project_id
LEFT JOIN milestones m ON p.ID = m.project_id
WHERE p.status = 'disputed'
GROUP BY p.ID
ORDER BY p.created_at;