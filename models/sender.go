package models

import (
	"time"

	"github.com/google/uuid"
)

type SenderType string

const (
	SenderTypePerson   SenderType = "person"
	SenderTypeBusiness            = "business"
)

type GenderType string

const (
	GenderTypeMale   GenderType = "male"
	GenderTypeFemale            = "female"
)

type IdentificationType string

const (
	IdentificationTypePassport       IdentificationType = "passport"
	IdentificationTypeChequeBook                        = "cheque_book"
	IdentificationTypeDrivingLicense                    = "driving_license"
	IdentificationTypeNationalId                        = "national_id"
	IdentificationTypeUtilityBill                       = "utility_bill"
	IdentificationTypeVotingCard                        = "voting_card"
	IdentificationTypeOther                             = "other"
)

type Sender struct {
	ID                   uuid.UUID
	Type                 SenderType
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
	IPAddress            string
	ExternalID           string
	CreatedAt            time.Time
	UpdatedAt            time.Time
	DeletedAt            time.Time
}
