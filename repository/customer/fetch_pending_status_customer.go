package customer

import (
	"context"
	"express_be/repository/customer/entity"
)

func (c *customerImpl) FetchPendingStatusCustomers(ctx context.Context, page, pageSize *int) ([]entity.Customer, error) {
	var result []entity.Customer
	offset := (*page - 1) * *pageSize
	query := c.DB.Executor.WithContext(ctx).
		Order("created_at DESC").
		Where("status = ?", entity.Pending).
		Offset(offset).
		Limit(*pageSize).
		Find(&result)

	if query.Error != nil {
		return nil, query.Error
	}

	return result, nil
}
