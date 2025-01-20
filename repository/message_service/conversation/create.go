package conversation

import (
	"context"
	"express_be/entity"
)

// CreateConversation implements Repo.
func (c *conversationImpl) CreateConversation(ctx context.Context, conversation *entity.Conversation) error {
	return c.DB.Executor.
		WithContext(ctx).
		Debug().
		Create(&conversation).
		Error
}
