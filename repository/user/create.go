package user

import (
	"context"
	userEntity "express_be/repository/user/entity"
	"fmt"
)

func (r *userImpl) Create(ctx context.Context, user *userEntity.User) error {
	if err := r.DB.Executor.WithContext(ctx).Create(user).Error; err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}
