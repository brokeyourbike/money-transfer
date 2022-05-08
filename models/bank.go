package models

import (
	"time"

	"github.com/google/uuid"
)

type Bank struct {
	ID          uuid.UUID
	Code        string
	SwiftCode   string
	CountryCode string
	Name        string
	ShortName   string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}
