package dto

type BackingRequest struct {
	Amount        int64   `json:"amount" validate:"required,number"`
	WordOfSupport *string `json:"word_of_support,omitempty"`
	UserID        *int64  `json:"user_id,omitempty"`
}
