package entity

import (
	"express_be/core/record"
	"time"
)

type RefreshToken struct {
	record.BaseEntity
	UserID    *string   `gorm:"user_id;not null"`
	Role      *string   `gorm:"role"`
	Token     *string   `gorm:"unique;not null"`
	ExpiresAt time.Time `gorm:"expires_at;not null"`
}

func (r *RefreshToken) TableName() string {
	return "refresh_tokens"
}
