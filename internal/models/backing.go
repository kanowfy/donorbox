package models

// mock
type CardInformation struct {
	Number          string `json:"card_number" validate:"required,credit_card"`
	ExpirationMonth int    `json:"expiration_month" validate:"required,min=1,max=12"`
	ExpirationYear  int    `json:"expiration_year" validate:"required,min=2024,max=2500"`
	CVV             string `json:"cvv" validate:"required,number,min=3,max=3"`
}

type BackingRequest struct {
	Amount          string          `json:"amount" validate:"required,number"`
	CardInformation CardInformation `json:"card_info" validate:"required"`
}
