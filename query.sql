--:::::::::: PROJECT ::::::::::--

-- name: GetAllProjects :many
SELECT projects.*, COUNT(backings.project_id) as backing_count
FROM projects
LEFT JOIN backings ON projects.ID = backings.project_id
WHERE category_id = 
    CASE WHEN @category::integer > 0 THEN @category::integer ELSE category_id END
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

-- name: UpdateFinishedProjectsStatus :exec
UPDATE projects SET status = 'ended'
WHERE end_date <= NOW() AND status = 'ongoing';

--:::::::::: PROJECT UPDATE ::::::::::--

-- name: GetProjectUpdates :many
SELECT * FROM project_updates
WHERE project_id = $1
ORDER BY created_at DESC;

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

--:::::::::: USER ::::::::::--

-- name: GetAllUsers :many
SELECT id, email, first_name, last_name, profile_picture, created_at FROM users;

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

--:::::::::: ESCROW USER ::::::::::--
-- name: GetEscrowUserByID :one
SELECT * FROM escrow_users
WHERE id = $1;

-- name: GetEscrowUserByEmail :one
SELECT * FROM escrow_users
WHERE email = $1;

-- name: GetEscrowUser :one
SELECT * FROM escrow_users
LIMIT 1;

-- name: UpdateEscrowUserByID :exec
UPDATE escrow_users SET email = $2, hashed_password = $3
WHERE id = $1;

-- name: UpdateEscrowCard :exec
UPDATE escrow_users SET card_id = $2
WHERE id = $1;

--:::::::::: BACKING ::::::::::--

-- name: GetBackingsForProject :many
SELECT users.id AS user_id, users.first_name, users.last_name, users.profile_picture, backings.id AS backing_id, backings.amount, backings.created_at FROM users
JOIN backings ON backings.backer_id = users.id
WHERE project_id = $1
ORDER BY backings.created_at DESC;

-- name: GetBackingByID :one
SELECT * FROM backings
WHERE id = $1;

-- name: GetBackingsForUser :many
SELECT * FROM backings
WHERE backer_id = $1;

-- name: GetMostBackingDonor :one
SELECT users.id AS user_id, users.first_name, users.last_name, users.profile_picture, backings.id AS backing_id, backings.amount, backings.created_at FROM users
JOIN backings ON backings.backer_id = users.id
WHERE project_id = $1
ORDER BY backings.amount DESC
LIMIT 1;

-- name: GetFirstBackingDonor :one
SELECT users.id AS user_id, users.first_name, users.last_name, users.profile_picture, backings.id AS backing_id, backings.amount, backings.created_at FROM users
JOIN backings ON backings.backer_id = users.id
WHERE project_id = $1
ORDER BY backings.created_at
LIMIT 1;

-- name: GetMostRecentBackingDonor :one
SELECT users.id AS user_id, users.first_name, users.last_name, users.profile_picture, backings.id AS backing_id, backings.amount, backings.created_at FROM users
JOIN backings ON backings.backer_id = users.id
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

--:::::::::: TRANSACTION ::::::::::--

-- name: GetAllTransactions :many
SELECT * FROM transactions
ORDER BY created_at DESC;

-- name: GetTransactionByID :one
SELECT * FROM transactions
WHERE id = $1;

-- name: GetTransactionsForProject :many
SELECT * FROM transactions
WHERE project_id = $1 ORDER BY created_at ASC;

-- name: GetBackingTransactionsForProject :many
SELECT * FROM transactions
WHERE project_id = $1 AND transaction_type == 'backing';

-- name: CreateTransaction :one
INSERT INTO transactions (
    project_id, transaction_type, amount, initiator_card_id, recipient_card_id
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;

-- name: UpdateTransactionStatus :exec
UPDATE transactions
SET status = $2
WHERE id = $1;

--:::::::::: CARD ::::::::::--

-- name: GetCardByID :one
SELECT * FROM cards
WHERE id = $1;

-- name: CreateCard :one
INSERT INTO cards (
    token, card_owner_name, last_four_digits, brand
) VALUES (
    $1, $2, $3, $4
)
RETURNING *;
