package model

import (
	"time"

	"github.com/google/uuid"
)

type EscrowUser struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	UserType  UserType  `json:"user_type"`
	CreatedAt time.Time `json:"created_at"`
}
