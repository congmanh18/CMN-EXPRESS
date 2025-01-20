package message

import (
	"context"
	"errors"
	"express_be/entity"
	"fmt"

	"gorm.io/gorm"
)

// GetMessagesByConversationID implements Repo.
func (m *messageImpl) GetMessagesByConversationID(ctx context.Context, conversationID *string, page, pageSize *int) ([]entity.Message, error) {
	var messages []entity.Message
	offset := (*page - 1) * *pageSize

	err := m.DB.Executor.
		WithContext(ctx).
		Debug().
		Where("conversation_id = ?", conversationID).
		Limit(*pageSize).
		Offset(offset).
		Find(&messages).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("messages with conversation_id %s not found", *conversationID)
		}
		return nil, err
	}
	return messages, nil
}
