package customer

import (
	"context"
	"express_be/repository/customer/entity"
)

func (c *customerImpl) DeleteCustomer(ctx context.Context, id *string) error {
	return c.DB.Executor.WithContext(ctx).
		Where("id = ?", id).
		Delete(&entity.Customer{}).
		Error
}
