package entity

import (
	"express_be/core/record"
	"time"
)

type MessageType string

const (
	Text  MessageType = "text"
	Image MessageType = "image"
	Video MessageType = "video"
	Audio MessageType = "audio"
)

type Message struct {
	record.BaseEntity
	ConversationID *string     `json:"conversation_id" gorm:"not null;index"`
	UserID         *string     `json:"user_id" gorm:"not null"`
	Content        *string     `json:"content"`
	MessageType    MessageType `json:"message_type"`
	ReadAt         *time.Time  `json:"read_at"`
}

func (m *Message) TableName() string {
	return "messages"
}
