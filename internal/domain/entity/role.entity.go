package entity

import "github.com/google/uuid"

type Role struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name        string    `gorm:"type:varchar(50);unique;not null"`
	Description string    `gorm:"type:text"`
	Code        string    `gorm:"type:varchar(50);unique;not null"`
}
