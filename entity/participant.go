package entity

import (
	"express_be/core/record"

	"time"
)

type Participant struct {
	record.BaseEntity
	ConversationID *string    `json:"conversation_id" gorm:"not null;index"`
	UserID         *string    `json:"user_id" gorm:"not null;index"`
	JoinedAt       *time.Time `json:"joined_at"`
	User           User       `gorm:"foreignKey:UserID;references:ID"`
}

func (p *Participant) TableName() string {
	return "participants"
}
