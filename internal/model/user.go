package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID `json:"id"`
	Email          string    `json:"email"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	ProfilePicture *string   `json:"profile_picture"`
	Activated      bool      `json:"activated"`
	UserType       UserType  `json:"user_type"`
	CreatedAt      time.Time `json:"created_at"`
}
