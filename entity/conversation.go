package entity

import (
	"express_be/core/record"
)

type ConversationType string

const (
	AdminConversationType   ConversationType = "admin"
	UserConversationType    ConversationType = "user"
	SupportConversationType ConversationType = "support"
)

type Conversation struct {
	record.BaseEntity
	Name         *string       `json:"name"`
	Messages     []Message     `gorm:"foreignKey:ConversationID"`
	Participants []Participant `gorm:"foreignKey:ConversationID"`
}

func (c *Conversation) TableName() string {
	return "conversations"
}
