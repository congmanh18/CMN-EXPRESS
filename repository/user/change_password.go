package user

import (
	"context"
	"express_be/entity"
	"fmt"
)

func (c *userImpl) ChangePassword(ctx context.Context, phone *string, newPassword *string) error {
	result := c.DB.Executor.WithContext(ctx).
		Model(&entity.User{}).
		Where("phone = ?", *phone).
		Update("password", newPassword)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no delivery person found with phone: %s", *phone)
	}

	return nil
}
