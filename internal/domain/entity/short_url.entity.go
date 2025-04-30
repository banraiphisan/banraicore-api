package entity

import "time"

type ShortUrl struct {
	ID        int64      `gorm:"primaryKey"`
	Code      string     `gorm:"column:short_code"`
	TargetUrl string     `gorm:"column:original_url"`
	CreatedAt time.Time  `gorm:"column:created_at"`
	ExpiresAt *time.Time `gorm:"column:expire_at"`
}
