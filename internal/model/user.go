package model

import (
	"time"
)

type VerificationStatus string

const (
	VerificationStatusPending    VerificationStatus = "pending"
	VerificationStatusUnverified VerificationStatus = "unverified"
	VerificationStatusVerified   VerificationStatus = "verified"
)

type User struct {
	ID                 int64              `json:"id"`
	Email              string             `json:"email"`
	FirstName          string             `json:"first_name"`
	LastName           string             `json:"last_name"`
	ProfilePicture     *string            `json:"profile_picture"`
	Activated          bool               `json:"activated"`
	VerificationStatus VerificationStatus `json:"verification_status"`
	UserType           UserType           `json:"user_type"`
	CreatedAt          time.Time          `json:"created_at"`
}
