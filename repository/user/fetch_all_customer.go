package user

import (
	"context"
	"express_be/repository/user/entity"
)

func (c *userImpl) FetchAllCustomer(ctx context.Context, page, pageSize *int) ([]entity.CustomerDetails, error) {
	var result []entity.CustomerDetails
	offset := (*page - 1) * *pageSize

	query := c.DB.Executor.
		WithContext(ctx).
		Table("users").
		Select(`
			users.*,
			customers.account_type,
			customers.latitude,
			customers.longtitude
		`).
		Joins("JOIN customers ON customers.phone = users.phone"). // Liên kết bảng
		Where("users.role = ?", entity.Customer).                 // Điều kiện lọc
		Order("users.full_name ASC").
		Offset(offset).
		Limit(*pageSize).
		Find(&result)

	if query.Error != nil {
		return nil, query.Error
	}

	return result, nil
}
