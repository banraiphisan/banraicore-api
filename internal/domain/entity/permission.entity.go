package entity

import (
	"github.com/google/uuid"
)

type Permission struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name        string    `gorm:"type:varchar(100);unique;not null"`
	Description string    `gorm:"type:text"`
}
