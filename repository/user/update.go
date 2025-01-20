package user

import (
	"context"
	"express_be/core/gorm"
	"express_be/entity"
	"fmt"
	"reflect"
)

func (c *userImpl) Update(ctx context.Context, id *string, user *entity.User) error {
	omitFields := gorm.OmitFields(user, func(fieldValue reflect.Value) bool {
		return fieldValue.IsZero()
	})

	err := c.DB.Executor.WithContext(ctx).
		Model(&entity.User{}).
		Where("id = ?", *id).
		Omit(omitFields...).
		Updates(user).Error

	if err != nil {
		return fmt.Errorf("failed to update user (ID: %d): %w", id, err)
	}

	return nil
}
