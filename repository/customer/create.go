package customer

import (
	"context"
	"express_be/repository/customer/entity"
	"fmt"
)

func (c *customerImpl) CreateCustomer(ctx context.Context, customer *entity.Customer) error {
	tx := c.DB.Executor.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var count int64
	if err := tx.WithContext(ctx).
		Model(&entity.Customer{}).
		Where("phone = ?", customer.Phone).
		Count(&count).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to check existing phone: %w", err)
	}

	if count > 0 {
		tx.Rollback()
		return fmt.Errorf("phone existing: %s", *customer.Phone)
	}

	if err := tx.WithContext(ctx).
		Create(customer).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to create customer: %w", err)
	}

	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
