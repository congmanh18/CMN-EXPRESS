package record

import (
	"time"

	"gorm.io/gorm"
)

type BaseEntity struct {
	ID        *string         `gorm:"primaryKey" json:"id,omitempty"`
	CreatedAt *time.Time      `gorm:"autoCreateTime" json:"created_at,omitempty"`
	UpdatedAt *time.Time      `gorm:"autoCreateTime" json:"updated_at,omitempty"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at,omitempty"`
}
