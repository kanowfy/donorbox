// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type BackingStatus string

const (
	BackingStatusPending  BackingStatus = "pending"
	BackingStatusReleased BackingStatus = "released"
	BackingStatusRefunded BackingStatus = "refunded"
)

func (e *BackingStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = BackingStatus(s)
	case string:
		*e = BackingStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for BackingStatus: %T", src)
	}
	return nil
}

type NullBackingStatus struct {
	BackingStatus BackingStatus
	Valid         bool // Valid is true if BackingStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullBackingStatus) Scan(value interface{}) error {
	if value == nil {
		ns.BackingStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.BackingStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullBackingStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.BackingStatus), nil
}

type CardBrand string

const (
	CardBrandVISA       CardBrand = "VISA"
	CardBrandMASTERCARD CardBrand = "MASTERCARD"
)

func (e *CardBrand) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = CardBrand(s)
	case string:
		*e = CardBrand(s)
	default:
		return fmt.Errorf("unsupported scan type for CardBrand: %T", src)
	}
	return nil
}

type NullCardBrand struct {
	CardBrand CardBrand
	Valid     bool // Valid is true if CardBrand is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullCardBrand) Scan(value interface{}) error {
	if value == nil {
		ns.CardBrand, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.CardBrand.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullCardBrand) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.CardBrand), nil
}

type ProjectStatus string

const (
	ProjectStatusOngoing         ProjectStatus = "ongoing"
	ProjectStatusEnded           ProjectStatus = "ended"
	ProjectStatusCompletedPayout ProjectStatus = "completed_payout"
	ProjectStatusCompletedRefund ProjectStatus = "completed_refund"
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

type ReportStatus string

const (
	ReportStatusPending   ReportStatus = "pending"
	ReportStatusResolved  ReportStatus = "resolved"
	ReportStatusDismissed ReportStatus = "dismissed"
)

func (e *ReportStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = ReportStatus(s)
	case string:
		*e = ReportStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for ReportStatus: %T", src)
	}
	return nil
}

type NullReportStatus struct {
	ReportStatus ReportStatus
	Valid        bool // Valid is true if ReportStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullReportStatus) Scan(value interface{}) error {
	if value == nil {
		ns.ReportStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.ReportStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullReportStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.ReportStatus), nil
}

type TransactionStatus string

const (
	TransactionStatusPending   TransactionStatus = "pending"
	TransactionStatusCompleted TransactionStatus = "completed"
	TransactionStatusFailed    TransactionStatus = "failed"
)

func (e *TransactionStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = TransactionStatus(s)
	case string:
		*e = TransactionStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for TransactionStatus: %T", src)
	}
	return nil
}

type NullTransactionStatus struct {
	TransactionStatus TransactionStatus
	Valid             bool // Valid is true if TransactionStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullTransactionStatus) Scan(value interface{}) error {
	if value == nil {
		ns.TransactionStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.TransactionStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullTransactionStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.TransactionStatus), nil
}

type TransactionType string

const (
	TransactionTypeBacking TransactionType = "backing"
	TransactionTypePayout  TransactionType = "payout"
	TransactionTypeRefund  TransactionType = "refund"
)

func (e *TransactionType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = TransactionType(s)
	case string:
		*e = TransactionType(s)
	default:
		return fmt.Errorf("unsupported scan type for TransactionType: %T", src)
	}
	return nil
}

type NullTransactionType struct {
	TransactionType TransactionType
	Valid           bool // Valid is true if TransactionType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullTransactionType) Scan(value interface{}) error {
	if value == nil {
		ns.TransactionType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.TransactionType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullTransactionType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.TransactionType), nil
}

type UserType string

const (
	UserTypeRegular UserType = "regular"
	UserTypeEscrow  UserType = "escrow"
)

func (e *UserType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = UserType(s)
	case string:
		*e = UserType(s)
	default:
		return fmt.Errorf("unsupported scan type for UserType: %T", src)
	}
	return nil
}

type NullUserType struct {
	UserType UserType
	Valid    bool // Valid is true if UserType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullUserType) Scan(value interface{}) error {
	if value == nil {
		ns.UserType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.UserType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullUserType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.UserType), nil
}

type Backing struct {
	ID            uuid.UUID
	ProjectID     uuid.UUID
	BackerID      uuid.UUID
	Amount        int64
	WordOfSupport *string
	Status        BackingStatus
	CreatedAt     time.Time
}

type Card struct {
	ID             uuid.UUID
	Token          string
	CardOwnerName  string
	LastFourDigits string
	Brand          CardBrand
	CreatedAt      time.Time
}

type Category struct {
	ID           int32
	Name         string
	Description  string
	CoverPicture string
}

type EscrowUser struct {
	ID             uuid.UUID
	Email          string
	HashedPassword string
	UserType       UserType
	CardID         uuid.UUID
	CreatedAt      time.Time
}

type Project struct {
	ID            uuid.UUID
	UserID        uuid.UUID
	CategoryID    int32
	Title         string
	Description   string
	CoverPicture  string
	GoalAmount    int64
	CurrentAmount int64
	Country       string
	Province      string
	CardID        uuid.UUID
	StartDate     time.Time
	EndDate       time.Time
	Status        ProjectStatus
}

type ProjectUpdate struct {
	ID              uuid.UUID
	ProjectID       uuid.UUID
	AttachmentPhoto *string
	Description     string
	CreatedAt       time.Time
}

type Report struct {
	ID                  uuid.UUID
	ProjectID           uuid.UUID
	ReporterEmail       string
	ReporterPhoneNumber string
	ReporterFullname    string
	Reason              string
	Details             string
	Status              ReportStatus
	CreatedAt           time.Time
}

type Transaction struct {
	ID              uuid.UUID
	ProjectID       uuid.UUID
	TransactionType TransactionType
	Amount          int64
	InitiatorCardID uuid.UUID
	RecipientCardID uuid.UUID
	Status          TransactionStatus
	CreatedAt       time.Time
}

type User struct {
	ID             uuid.UUID
	Email          string
	HashedPassword string
	FirstName      string
	LastName       string
	ProfilePicture *string
	Activated      bool
	UserType       UserType
	CreatedAt      time.Time
}
