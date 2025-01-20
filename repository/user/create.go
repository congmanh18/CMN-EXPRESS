package user

import (
	"context"
	"express_be/entity"
	"fmt"
)

func (r *userImpl) Create(ctx context.Context, user *entity.User) error {
	if err := r.DB.Executor.WithContext(ctx).Create(user).Error; err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}
