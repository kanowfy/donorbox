package model

import (
	"time"
)

type User struct {
	ID             int64     `json:"id"`
	Email          string    `json:"email"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	ProfilePicture *string   `json:"profile_picture"`
	Activated      bool      `json:"activated"`
	UserType       UserType  `json:"user_type"`
	CreatedAt      time.Time `json:"created_at"`
}
