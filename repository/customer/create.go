package customer

import (
	"context"
	"express_be/repository/customer/entity"
)

func (c *customerImpl) CreateCustomer(ctx context.Context, customer *entity.Customer) error {
	return c.DB.Executor.WithContext(ctx).
		Debug().
		Create(&customer).Error
}
