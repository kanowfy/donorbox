-- name: GetAllProjects :many
SELECT * FROM projects
WHERE category_id = 
    CASE WHEN @category::integer > 0 THEN @category::integer ELSE category_id END
ORDER BY
    CASE WHEN @end_date_asc::integer > 0 THEN end_date END ASC,
    CASE WHEN @end_date_desc::integer > 0 THEN end_date END DESC,
    CASE WHEN @current_amount_asc::integer > 0 THEN current_amount END ASC,
    CASE WHEN @current_amount_desc::integer > 0 THEN current_amount END DESC
LIMIT @page_limit::integer OFFSET @total_offset::integer;

-- name: GetProjectByID :one
SELECT * FROM projects
WHERE id = $1;

-- name: UpdateProjectByID :exec
UPDATE projects
SET title = $2, description = $3, cover_picture = $4, goal_amount = $5, country = $6, province = $7, end_date = $8
WHERE id = $1;

-- name: UpdateProjectFund :exec
UPDATE projects SET current_amount = current_amount + @backing_amount::bigint
WHERE id = $1;

-- name: UpdateProjectPaymentID :exec
UPDATE projects SET payment_id = $2
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
WHERE project_id = $1;

-- name: DeleteProjectUpdate :exec
DELETE FROM project_updates
WHERE id = $1;

-- name: CreateProjectUpdate :one
INSERT INTO project_updates (
    project_id, description
) VALUES (
    $1, $2
)
RETURNING *;

-- name: GetProjectComments :many
SELECT * FROM project_comments
WHERE project_id = $1;

-- name: DeleteProjectComment :exec
DELETE FROM project_comments
WHERE id = $1;

-- name: CreateProjectComment :one
INSERT INTO project_comments (
    project_id, backer_id, parent_comment_id, content
) VALUES (
    $1, $2, $3, $4
)
RETURNING *;

-- name: GetAllUsers :many
SELECT id, email, first_name, last_name, profile_picture, activated, user_type, created_at FROM users;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1;

-- name: UpdateUserByID :exec
UPDATE users
SET email = $2, first_name = $3, last_name = $4, profile_picture = $5, activated = $6
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

-- name: GetEscrowUserByID :one
SELECT * FROM escrow_users
WHERE id = $1;

-- name: GetEscrowUserByEmail :one
SELECT * FROM escrow_users
WHERE email = $1;

-- name: UpdateEscrowUserPaymentID :exec
UPDATE escrow_users
SET payment_id = $2
WHERE id = $1;

-- name: GetAllTransactions :many
SELECT * FROM transactions;

-- name: GetTransactionByID :one
SELECT * FROM transactions
WHERE id = $1;

-- name: CreateTransaction :one
INSERT INTO transactions (
    project_id, transaction_type, amount, initiator_id, recipient_id
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;

-- name: UpdateTransactionStatus :exec
UPDATE transactions
SET status = $2
WHERE id = $1;

-- name: GetBackingsForProject :many
SELECT * FROM backings
WHERE project_id = $1
ORDER BY backing_date DESC;

-- name: GetBackingByID :one
SELECT * FROM backings
WHERE id = $1;

-- name: GetBackingsForUser :many
SELECT * FROM backings
WHERE backer_id = $1;

-- name: CreateBacking :one
INSERT INTO backings (
    project_id, backer_id, amount
) VALUES (
    $1, $2, $3
) 
RETURNING *;

-- name: UpdateProjectBackingStatus :exec
UPDATE backings
SET status = $2
WHERE project_id = $1;
