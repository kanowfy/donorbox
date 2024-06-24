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
