-- name: CreateAuditLog :one
INSERT INTO audit_trails (
    user_id, escrow_id, entity_type, entity_id, operation_type, field_name, old_value, new_value
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING *;

-- name: GetAuditHistory :many
SELECT * FROM audit_trails
ORDER BY created_at DESC;