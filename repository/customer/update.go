package customer

import (
	"context"
	"express_be/repository/customer/entity"
	"fmt"
)

// UpdateCustomer implements Repo.
func (c *customerImpl) UpdateCustomer(ctx context.Context, id *string, customer *entity.Customer) error {
	err := c.DB.Executor.WithContext(ctx).
		Model(&entity.Customer{}).
		Where("id = ?", id).
		Updates(customer).Error

	if err != nil {
		return fmt.Errorf("failed to update customer (ID: %d): %w", id, err)
	}

	return nil

}
