package user

import (
	"context"
	"express_be/entity"
)

func (c *userImpl) FetchCustomerUsers(ctx context.Context, status *string, page, pageSize *int) ([]entity.CustomerDetails, error) {
	var result []entity.CustomerDetails

	defaultPage, defaultPageSize := 1, 10
	if page == nil || *page < 1 {
		page = &defaultPage
	}
	if pageSize == nil || *pageSize < 1 {
		pageSize = &defaultPageSize
	}

	offset := (*page - 1) * *pageSize

	// Bắt đầu xây dựng câu truy vấn
	query := c.DB.Executor.
		WithContext(ctx).
		Table("users").
		Select(`
			users.*,
			customers.account_type,
			customers.latitude,
			customers.longitude
		`).
		Joins("JOIN customers ON customers.phone = users.phone").
		Where("users.role = ?", entity.CustomerRole)

	// Thêm điều kiện nếu `status` được truyền
	if status != nil && *status != "" {
		query = query.Where("users.status = ?", *status)
	}
	// Thêm sắp xếp, phân trang
	query = query.
		Order("users.first_name ASC").
		Offset(offset).
		Limit(*pageSize).
		Find(&result)

	// Kiểm tra lỗi
	if query.Error != nil {
		return nil, query.Error
	}

	return result, nil
}
