package conversation

import (
	"context"
	"express_be/entity"
)

// GetAllConversations implements Repo.
func (c *conversationImpl) GetAllConversations(ctx context.Context, userID *string, page, pageSize *int) ([]entity.Conversation, error) {
	var conversations []entity.Conversation
	offset := (*page - 1) * *pageSize

	err := c.DB.Executor.
		WithContext(ctx).
		Debug().
		Table("conversations").
		Joins("JOIN participants ON participants.conversation_id = conversations.id").
		Where("participants.user_id = ?", *userID).
		Order("conversations.updated_at DESC").
		Limit(*pageSize).
		Offset(offset).
		Find(&conversations).
		Error

	return conversations, err
}
