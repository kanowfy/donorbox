// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"database/sql/driver"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
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
	BackingStatus BackingStatus `json:"backing_status"`
	Valid         bool          `json:"valid"` // Valid is true if BackingStatus is not NULL
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
	TransactionStatus TransactionStatus `json:"transaction_status"`
	Valid             bool              `json:"valid"` // Valid is true if TransactionStatus is not NULL
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
	TransactionType TransactionType `json:"transaction_type"`
	Valid           bool            `json:"valid"` // Valid is true if TransactionType is not NULL
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
	UserType UserType `json:"user_type"`
	Valid    bool     `json:"valid"` // Valid is true if UserType is not NULL
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
	ID          pgtype.UUID        `json:"id"`
	ProjectID   pgtype.UUID        `json:"project_id"`
	BackerID    pgtype.UUID        `json:"backer_id"`
	Amount      int64              `json:"amount"`
	BackingDate pgtype.Timestamptz `json:"backing_date"`
	Status      BackingStatus      `json:"status"`
}

type Category struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

type EscrowUser struct {
	ID             pgtype.UUID        `json:"id"`
	Email          string             `json:"email"`
	HashedPassword string             `json:"hashed_password"`
	UserType       UserType           `json:"user_type"`
	PaymentID      pgtype.Text        `json:"payment_id"`
	CreatedAt      pgtype.Timestamptz `json:"created_at"`
}

type Project struct {
	ID            pgtype.UUID        `json:"id"`
	UserID        pgtype.UUID        `json:"user_id"`
	CategoryID    int32              `json:"category_id"`
	Title         string             `json:"title"`
	Description   string             `json:"description"`
	CoverPicture  string             `json:"cover_picture"`
	GoalAmount    int64              `json:"goal_amount"`
	CurrentAmount int64              `json:"current_amount"`
	Country       string             `json:"country"`
	Province      string             `json:"province"`
	StartDate     pgtype.Timestamptz `json:"start_date"`
	EndDate       pgtype.Timestamptz `json:"end_date"`
	PaymentID     pgtype.Text        `json:"payment_id"`
	IsActive      bool               `json:"is_active"`
}

type ProjectComment struct {
	ID              pgtype.UUID        `json:"id"`
	ProjectID       pgtype.UUID        `json:"project_id"`
	BackerID        pgtype.UUID        `json:"backer_id"`
	ParentCommentID pgtype.UUID        `json:"parent_comment_id"`
	Content         string             `json:"content"`
	CommentedAt     pgtype.Timestamptz `json:"commented_at"`
}

type ProjectUpdate struct {
	ID          pgtype.UUID        `json:"id"`
	ProjectID   pgtype.UUID        `json:"project_id"`
	Description string             `json:"description"`
	UpdateDate  pgtype.Timestamptz `json:"update_date"`
}

type Transaction struct {
	ID              pgtype.UUID        `json:"id"`
	ProjectID       pgtype.UUID        `json:"project_id"`
	TransactionType TransactionType    `json:"transaction_type"`
	Amount          int64              `json:"amount"`
	InitiatorID     pgtype.UUID        `json:"initiator_id"`
	RecipientID     pgtype.UUID        `json:"recipient_id"`
	Status          TransactionStatus  `json:"status"`
	CreateAt        pgtype.Timestamptz `json:"create_at"`
}

type User struct {
	ID             pgtype.UUID        `json:"id"`
	Email          string             `json:"email"`
	HashedPassword string             `json:"hashed_password"`
	FirstName      string             `json:"first_name"`
	LastName       string             `json:"last_name"`
	ProfilePicture pgtype.Text        `json:"profile_picture"`
	Activated      bool               `json:"activated"`
	UserType       UserType           `json:"user_type"`
	CreatedAt      pgtype.Timestamptz `json:"created_at"`
}
