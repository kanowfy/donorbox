// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"database/sql/driver"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

type NotificationType string

const (
	NotificationTypeApprovedVerification NotificationType = "approved_verification"
	NotificationTypeRejectedVerification NotificationType = "rejected_verification"
	NotificationTypeApprovedProject      NotificationType = "approved_project"
	NotificationTypeRejectedProject      NotificationType = "rejected_project"
	NotificationTypeMilestoneCompletion  NotificationType = "milestone_completion"
)

func (e *NotificationType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = NotificationType(s)
	case string:
		*e = NotificationType(s)
	default:
		return fmt.Errorf("unsupported scan type for NotificationType: %T", src)
	}
	return nil
}

type NullNotificationType struct {
	NotificationType NotificationType
	Valid            bool // Valid is true if NotificationType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullNotificationType) Scan(value interface{}) error {
	if value == nil {
		ns.NotificationType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.NotificationType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullNotificationType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.NotificationType), nil
}

type ProjectStatus string

const (
	ProjectStatusPending  ProjectStatus = "pending"
	ProjectStatusOngoing  ProjectStatus = "ongoing"
	ProjectStatusRejected ProjectStatus = "rejected"
	ProjectStatusFinished ProjectStatus = "finished"
)

func (e *ProjectStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = ProjectStatus(s)
	case string:
		*e = ProjectStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for ProjectStatus: %T", src)
	}
	return nil
}

type NullProjectStatus struct {
	ProjectStatus ProjectStatus
	Valid         bool // Valid is true if ProjectStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullProjectStatus) Scan(value interface{}) error {
	if value == nil {
		ns.ProjectStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.ProjectStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullProjectStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.ProjectStatus), nil
}

type VerificationStatus string

const (
	VerificationStatusUnverified VerificationStatus = "unverified"
	VerificationStatusPending    VerificationStatus = "pending"
	VerificationStatusVerified   VerificationStatus = "verified"
)

func (e *VerificationStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = VerificationStatus(s)
	case string:
		*e = VerificationStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for VerificationStatus: %T", src)
	}
	return nil
}

type NullVerificationStatus struct {
	VerificationStatus VerificationStatus
	Valid              bool // Valid is true if VerificationStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullVerificationStatus) Scan(value interface{}) error {
	if value == nil {
		ns.VerificationStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.VerificationStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullVerificationStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.VerificationStatus), nil
}

type AuditTrail struct {
	ID            int64
	UserID        *int64
	EscrowID      *int64
	EntityType    string
	EntityID      *int64
	OperationType string
	FieldName     string
	OldValue      []byte
	NewValue      []byte
	CreatedAt     pgtype.Timestamptz
}

type Backing struct {
	ID            int64
	UserID        *int64
	ProjectID     int64
	Amount        int64
	WordOfSupport *string
	CreatedAt     pgtype.Timestamptz
}

type Category struct {
	ID           int32
	Name         string
	Description  string
	CoverPicture string
}

type EscrowUser struct {
	ID             int64
	Email          string
	HashedPassword string
	CreatedAt      pgtype.Timestamptz
}

type Milestone struct {
	ID              int64
	ProjectID       int64
	Title           string
	Description     *string
	FundGoal        int64
	CurrentFund     int64
	BankDescription string
	Completed       bool
	CreatedAt       pgtype.Timestamptz
}

type MilestoneCompletion struct {
	ID             int64
	MilestoneID    int64
	TransferAmount int64
	TransferNote   *string
	TransferImage  *string
	CompletedAt    pgtype.Timestamptz
}

type Notification struct {
	ID               int64
	UserID           int64
	NotificationType NotificationType
	Message          string
	ProjectID        *int64
	IsRead           bool
	CreatedAt        pgtype.Timestamptz
}

type Project struct {
	ID             int64
	UserID         int64
	Title          string
	Description    string
	CoverPicture   string
	CategoryID     int32
	EndDate        pgtype.Timestamptz
	ReceiverNumber string
	ReceiverName   string
	Address        string
	District       string
	City           string
	Country        string
	Status         ProjectStatus
	CreatedAt      pgtype.Timestamptz
}

type ProjectUpdate struct {
	ID              int64
	ProjectID       int64
	AttachmentPhoto *string
	Description     string
	CreatedAt       pgtype.Timestamptz
}

type User struct {
	ID                      int64
	Email                   string
	FirstName               string
	LastName                string
	ProfilePicture          *string
	HashedPassword          string
	Activated               bool
	VerificationStatus      VerificationStatus
	VerificationDocumentUrl *string
	CreatedAt               pgtype.Timestamptz
}
