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
WHERE project_id = $1 AND transaction_type = 'backing';

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

-- name: GetStatistics :one
SELECT (
	SELECT COUNT(*) AS status_aggregation
	FROM projects GROUP BY status HAVING status = 'ended'
) AS ended, (
	SELECT COUNT(*) AS status_aggregation
	FROM projects GROUP BY status HAVING status = 'ongoing'
) AS ongoing, (
	SELECT COUNT(*) AS status_aggregation
	FROM projects GROUP BY status HAVING status = 'completed_payout'
) AS completed_payout, (
	SELECT COUNT(*) AS status_aggregation
	FROM projects GROUP BY status HAVING status = 'completed_refund'
) AS completed_refund, ((
	SELECT SUM(amount) AS transaction_amount
	FROM transactions GROUP BY transaction_type HAVING transaction_type = 'backing'
) - (
	SELECT SUM(amount) AS transaction_amount
	FROM transactions GROUP BY transaction_type HAVING transaction_type = 'payout'
) - (
	SELECT SUM(amount) AS transaction_amount
	FROM transactions GROUP BY transaction_type HAVING transaction_type = 'refund'
)) AS balance;

-- name: GetTransactionsStatsByWeek :many
SELECT
    DATE_TRUNC('week', created_at)::date AS week,
    COUNT(*) FILTER (WHERE transaction_type = 'backing') AS backings,
	COUNT(*) FILTER (WHERE transaction_type = 'payout') AS payouts,
	COUNT(*) FILTER (WHERE transaction_type = 'refund') AS refunds
FROM
    transactions
GROUP BY
    week
ORDER BY
    week;
