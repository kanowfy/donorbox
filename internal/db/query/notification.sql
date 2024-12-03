-- name: CreateNotification :one
INSERT INTO notifications (
    user_id, notification_type, message, project_id
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetNotificationsForUser :many
SELECT * FROM notifications
WHERE user_id = $1
ORDER BY created_at DESC;

-- name: UpdateReadNotification :exec
UPDATE notifications
SET is_read = TRUE
WHERE id = $1;