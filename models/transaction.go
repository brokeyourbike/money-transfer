package models

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TransactionType string

const (
	TransactionTypeBank   TransactionType = "bank"
	TransactionTypeCash                   = "cash"
	TransactionTypeMobile                 = "mobile"
	TransactionTypeCard                   = "card"
	TransactionTypeTopup                  = "topup"
)

type Transaction struct {
	gorm.Model
	ID                  uuid.UUID
	CustomerID          uuid.UUID
	SenderID            uuid.UUID
	RecipientID         uuid.UUID
	ReferenceID         uint32
	Type                TransactionType
	InputCurrencyCode   string
	OutputCurrencyCode  string
	OutputAmountInCents string
	ExternalID          string
	ExpiresAt           sql.NullTime
	InvalidAt           sql.NullTime
}

func (t *Transaction) GetReferenceForHumans() string {
	return fmt.Sprintf("%010d", t.ReferenceID)
}
