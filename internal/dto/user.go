package dto

import "time"

type UserRegisterRequest struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8,max=50"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
}

type EscrowRegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=50"`
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
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" validate:"required"`
	NewPassword string `json:"new_password" validate:"min=8,max=50"`
}

type ResetPasswordRequest struct {
	ResetToken  string `json:"reset_token"`
	NewPassword string `json:"new_password" validate:"min=8,max=50"`
}

type PendingUserVerificationResponse struct {
	ID          int64     `json:"id"`
	Email       string    `json:"email"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	DocumentUrl string    `json:"document_url"`
	CreatedAt   time.Time `json:"created_at"`
}

type VerificationApprovalRequest struct {
	UserID       int64   `json:"user_id" validate:"required,number"`
	Approved     *bool   `json:"approved,omitempty" validate:"omitnil"`
	RejectReason *string `json:"reject_reason,omitempty" validate:"omitnil"`
}
