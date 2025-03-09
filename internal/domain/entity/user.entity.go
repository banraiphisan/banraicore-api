package entity

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID           uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Username     string     `gorm:"type:varchar(50);unique;not null"`
	Email        string     `gorm:"type:varchar(100);unique;not null"`
	PasswordHash string     `gorm:"type:text;not null"`
	RoleID       *uuid.UUID `gorm:"type:uuid;constraint:OnDelete:SET NULL"`
	Role         *Role      `gorm:"foreignKey:RoleID"`
	CreatedAt    time.Time  `gorm:"default:now()"`
}
