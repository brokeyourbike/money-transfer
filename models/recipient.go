package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BankAccountType string

const (
	BankAccountTypeCurrent BankAccountType = "current"
	BankAccountTypeSavings                 = "savings"
)

type Recipient struct {
	gorm.Model
	ID                   uuid.UUID
	FirstName            string
	LastName             string
	BirthDate            time.Time
	Gender               GenderType
	IdentificationType   IdentificationType
	IdentificationNumber string
	IdentificationExpiry time.Time
	CountryCode          string
	Street               string
	PostalCode           string
	City                 string
	Email                string
	PhoneNumber          string
	MobileProvider       string
	TransferReasonCode   string
	BankName             string
	BankCode             string
	BankAccountType      BankAccountType
	BankAccount          string
	SortCode             string
	IBAN                 string
	BIC                  string
	ExternalID           string
}
