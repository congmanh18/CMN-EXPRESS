package customer

import (
	"context"
	"express_be/repository/customer/entity"
)

func (c *customerImpl) FindByPhone(ctx context.Context, phone *string) (*entity.Customer, error) {
	var result *entity.Customer
	query := c.DB.Executor.WithContext(ctx).
		Model(&entity.Customer{}).
		Where("phone = ?", phone).
		First(&result)

	if query.Error != nil {
		return nil, query.Error
	}

	return result, nil
}
