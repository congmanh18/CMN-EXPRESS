package user

import (
	"context"
	"express_be/repository/user/entity"
	"fmt"
)

// Create implements Repo.
func (c *userImpl) Create(ctx context.Context, user *entity.User) error {
	tx := c.DB.Executor.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var count int64
	if err := tx.WithContext(ctx).
		Model(&entity.User{}).
		Where("phone = ?", user.Phone).
		Count(&count).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to check existing phone: %w", err)
	}

	if count > 0 {
		tx.Rollback()
		return fmt.Errorf("phone existing: %s", *user.Phone)
	}

	if err := tx.WithContext(ctx).
		Create(user).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to create user: %w", err)
	}

	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil

}
