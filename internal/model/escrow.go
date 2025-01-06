package model

import (
	"time"
)

type EscrowUser struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	UserType  UserType  `json:"user_type"`
	CreatedAt time.Time `json:"created_at"`
}
