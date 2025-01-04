package delivery

import (
	"context"
	"express_be/repository/delivery/entity"
	"fmt"
)

func (c *deliveryImpl) ChangePassword(ctx context.Context, phone *string, newPassword *string) error {
	result := c.DB.Executor.WithContext(ctx).
		Model(&entity.DeliveryPerson{}).
		Where("phone =?", *phone).
		Update("password_hash", newPassword)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no delivery person found with phone: %s", *phone)
	}

	return nil
}
