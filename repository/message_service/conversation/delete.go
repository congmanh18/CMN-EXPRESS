package conversation

import (
	"context"
	"express_be/entity"
)

// DeleteConversation implements Repo.
func (c *conversationImpl) DeleteConversation(ctx context.Context, id *string) error {
	return c.DB.Executor.
		WithContext(ctx).
		Debug().
		Where("id = ?", id).
		Delete(&entity.Conversation{}).
		Error
}
