package model

import (
	"time"

	"github.com/google/uuid"
)

type EscrowUser struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	UserType  UserType  `json:"user_type"`
	CardID    uuid.UUID `json:"card_id"`
	CreatedAt time.Time `json:"created_at"`
}
