package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	ID    uuid.UUID
	Name  string
	Email string
}
