package message

import (
	"context"
	"express_be/entity"
)

// SaveMessage implements Repo.
func (m *messageImpl) SaveMessage(ctx context.Context, message *entity.Message) error {
	return m.DB.Executor.
		WithContext(ctx).
		Debug().
		Create(&message).
		Error
}
