package entity

import (
	"github.com/google/uuid"
)

type RolePermission struct {
	RoleID       uuid.UUID  `gorm:"type:uuid;primaryKey"`
	PermissionID uuid.UUID  `gorm:"type:uuid;primaryKey"`
	Role         Role       `gorm:"foreignKey:RoleID;constraint:OnDelete:CASCADE"`
	Permission   Permission `gorm:"foreignKey:PermissionID;constraint:OnDelete:CASCADE"`
}
