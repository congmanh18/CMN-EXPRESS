package res

import "time"

type ConversationRes struct {
	ConversationID *string    `json:"conversation_id"`
	Name           *string    `json:"name"`
	UpdatedAt      *time.Time `json:"updated_at"`
}
