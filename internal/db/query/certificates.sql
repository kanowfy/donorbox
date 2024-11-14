-- name: GetAllVerifiedCertificates :many
SELECT * FROM certificates
WHERE verified
ORDER BY created_at DESC;

-- name: GetCerificateByID :one
SELECT * FROM certificates
WHERE id = $1;

-- name: CreateCertificate :one
INSERT INTO certificates (
    escrow_user_id, user_id, milestone_id
) VALUES (
    $1, $2, $3
)
RETURNING *;

-- name: UpdateVerifyingCertificate :exec
UPDATE certificates
SET verified = TRUE, verified_at = $2
WHERE milestone_id = $1;
