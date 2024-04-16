package models

type RegisterAccountRequest struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8,max=50"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UpdateAccountRequest struct {
	Email          *string `json:"email,omitempty" validate:"omitnil,email"`
	FirstName      *string `json:"first_name,omitempty"`
	LastName       *string `json:"last_name,omitempty"`
	ProfilePicture *string `json:"profile_picture,omitempty" validate:"omitnil,http_url"`
	Activated      *bool   `json:"activated,omitempty"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" validate:"required"`
	NewPassword string `json:"new_password" validate:"min=8,max=50"`
}
