package models

type BackingRequest struct {
	Amount          string          `json:"amount" validate:"required,number"`
	CardInformation CardInformation `json:"card_info" validate:"required"`
	WordOfSupport   *string         `json:"word_of_support,omitempty"`
}
