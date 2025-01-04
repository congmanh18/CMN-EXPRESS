package customer

import (
	"context"
	"express_be/repository/customer/entity"
	"fmt"
)

func (c *customerImpl) UpdateStatus(ctx context.Context, id *string, status *entity.Status) error {
	result := c.DB.Executor.WithContext(ctx).
		Model(&entity.Customer{}).
		Where("id = ?", *id).
		Update("status", status)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no customer found with id: %s", *id)
	}

	return nil // Thành công
}
