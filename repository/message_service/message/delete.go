package message

import (
	"context"
	"express_be/entity"
)

// DeleteMessage implements Repo.
func (m *messageImpl) DeleteMessage(ctx context.Context, messageID *string) error {
	return m.DB.Executor.
		WithContext(ctx).
		Debug().
		Where("id = ?", messageID).
		Delete(&entity.Message{}).
		Error
}
