package dto

type BackingRequest struct {
	Amount          string          `json:"amount" validate:"required,number"`
	CardInformation CardInformation `json:"card_info" validate:"required"`
	WordOfSupport   *string         `json:"word_of_support,omitempty"`
	UserID          *string         `json:"user_id,omitempty"`
}

type CardInformation struct {
	Number          string `json:"card_number" validate:"required,credit_card"`
	ExpirationMonth int    `json:"expiration_month" validate:"required,min=1,max=12"`
	ExpirationYear  int    `json:"expiration_year" validate:"required,min=2024,max=2050"`
	CVV             string `json:"cvv" validate:"required,number,min=3,max=3"`
	Brand           string `json:"brand" validate:"required"`
	OwnerName       string `json:"owner_name" validate:"required,max=64"`
}
