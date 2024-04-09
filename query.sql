-- name: GetAllUsers :many
SELECT * FROM users;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1;

-- name: GetUserByUsername :one
SELECT * FROM users
WHERE username = $1;

-- name: ActivateUser :exec
UPDATE users SET activated = TRUE
WHERE id = $1;

-- name: UpdateUserProfilePicture :exec
UPDATE users SET profile_picture = $2
WHERE id = $1;

-- name: UpdateUserPassword :exec
UPDATE users SET hashed_password = $2
WHERE id = $1;

-- name: DeleteUserByID :exec
DELETE FROM users WHERE id = $1;

-- name: CreateUser :one
INSERT INTO users (
    username, hashed_password, email, first_name, last_name, profile_picture
) VALUES (
    $1, $2, $3, $4, $5, $6
)
RETURNING *;

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
UPDATE projects SET current_amount = current_amount + $2
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
