package customer

import (
	"context"
	"express_be/core/gorm"
	"express_be/repository/customer/entity"
	"fmt"
	"reflect"
)

// UpdateCustomer implements Repo.
func (c *customerImpl) UpdateCustomer(ctx context.Context, id *string, customer *entity.Customer) error {
	omitFields := gorm.OmitFields(customer, func(fieldValue reflect.Value) bool {
		return fieldValue.IsZero()
	})

	err := c.DB.Executor.WithContext(ctx).
		Model(&entity.Customer{}).
		Where("id = ?", *id).
		Omit(omitFields...).
		Updates(customer).Error

	if err != nil {
		return fmt.Errorf("failed to update customer (ID: %d): %w", id, err)
	}

	return nil

}
