package dto

import "time"

type BackingRequest struct {
	Amount        int64   `json:"amount" validate:"required,number"`
	WordOfSupport *string `json:"word_of_support,omitempty"`
	UserID        *int64  `json:"user_id,omitempty"`
}

type UserBackingResponse struct {
	BackingID           int64     `json:"backing_id"`
	ProjectID           int64     `json:"project_id"`
	Amount              int64     `json:"amount"`
	WordOfSupport       *string   `json:"word_of_support,omitempty"`
	CreatedAt           time.Time `json:"created_at"`
	ProjectTitle        string    `json:"title"`
	ProjectCoverPicture string    `json:"cover_picture"`
}
