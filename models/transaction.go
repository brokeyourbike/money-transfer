package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
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
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           time.Time
	ExpiresAt           time.Time
	InvalidAt           time.Time
}

func (t *Transaction) GetReferenceForHumans() string {
	return fmt.Sprintf("%010d", t.ReferenceID)
}
