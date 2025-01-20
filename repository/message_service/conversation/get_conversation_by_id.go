package conversation

import (
	"context"
	"errors"
	"express_be/entity"
	"fmt"

	"gorm.io/gorm"
)

// GetConversationByID implements Repo.
func (c *conversationImpl) GetConversationByID(ctx context.Context, id *string) (*entity.Conversation, error) {
	conversation := entity.Conversation{}
	err := c.DB.Executor.
		WithContext(ctx).
		Debug().
		Where("id = ?", id).
		First(&conversation).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("conversation with id %s not found", *id)
		}
		return nil, err
	}
	return &conversation, nil
}
