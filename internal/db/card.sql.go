// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: card.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createCard = `-- name: CreateCard :one
INSERT INTO cards (
    token, card_owner_name, last_four_digits, brand
) VALUES (
    $1, $2, $3, $4
)
RETURNING id, token, card_owner_name, last_four_digits, brand, created_at
`

type CreateCardParams struct {
	Token          string
	CardOwnerName  string
	LastFourDigits string
	Brand          CardBrand
}

func (q *Queries) CreateCard(ctx context.Context, arg CreateCardParams) (Card, error) {
	row := q.db.QueryRow(ctx, createCard,
		arg.Token,
		arg.CardOwnerName,
		arg.LastFourDigits,
		arg.Brand,
	)
	var i Card
	err := row.Scan(
		&i.ID,
		&i.Token,
		&i.CardOwnerName,
		&i.LastFourDigits,
		&i.Brand,
		&i.CreatedAt,
	)
	return i, err
}

const getCardByID = `-- name: GetCardByID :one
SELECT id, token, card_owner_name, last_four_digits, brand, created_at FROM cards
WHERE id = $1
`

func (q *Queries) GetCardByID(ctx context.Context, id uuid.UUID) (Card, error) {
	row := q.db.QueryRow(ctx, getCardByID, id)
	var i Card
	err := row.Scan(
		&i.ID,
		&i.Token,
		&i.CardOwnerName,
		&i.LastFourDigits,
		&i.Brand,
		&i.CreatedAt,
	)
	return i, err
}
