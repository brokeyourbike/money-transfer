package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Bank struct {
	gorm.Model
	ID          uuid.UUID
	Code        string
	SwiftCode   string
	CountryCode string
	Name        string
	ShortName   string
}
