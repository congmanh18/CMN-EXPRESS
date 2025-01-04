package customer

import (
	"context"
	"express_be/repository/customer/entity"
	"fmt"
)

func (c *customerImpl) ChangePassword(ctx context.Context, phone *string, newPassword *string) error {
	result := c.DB.Executor.WithContext(ctx).
		Model(&entity.Customer{}).
		Where("phone =?", *phone).
		Update("password_hash", newPassword)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no customer found with phone: %s", *phone)
	}

	return nil
}
