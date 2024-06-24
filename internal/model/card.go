package model

import (
	"time"

	"github.com/google/uuid"
)

type CardBrand string

const (
	VISA       CardBrand = "VISA"
	MASTERCARD CardBrand = "MASTERCARD"
)

type Card struct {
	ID             uuid.UUID `json:"id"`
	Token          string    `json:"token"`
	CardOwnerName  string    `json:"card_owner_name"`
	LastFourDigits string    `json:"last_four_digits"`
	Brand          CardBrand `json:"brand"`
	CreatedAt      time.Time `json:"created_at"`
}
