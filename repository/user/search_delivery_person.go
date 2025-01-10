package user

import (
	"context"
	"express_be/repository/user/entity"
)

// SearchDeliveryPersons implements Repo.
func (c *userImpl) SearchDeliveryPersons(ctx context.Context, phone *string, name *string, status *string, page *int, pageSize *int) ([]entity.DeliveryPersonDetails, int64, error) {
	var result []entity.DeliveryPersonDetails
	var total int64

	// Đặt giá trị mặc định cho phân trang
	defaultPage, defaultPageSize := 1, 10
	if page == nil || *page < 1 {
		page = &defaultPage
	}
	if pageSize == nil || *pageSize < 1 {
		pageSize = &defaultPageSize
	}

	offset := (*page - 1) * *pageSize

	// Xây dựng truy vấn cơ bản
	query := c.DB.Executor.WithContext(ctx).
		Table("users").
		Select(`
			users.*,
			delivery_persons.salary_rate
		`).
		Joins("JOIN delivery_persons ON delivery_persons.phone = users.phone").
		Where("users.role = ?", entity.DeliveryPerson)

	// Thêm điều kiện nếu `status` được truyền
	if status != nil && *status != "" {
		query = query.Where("users.status = ?", *status)
	}

	// Thêm điều kiện nếu `name` được truyền
	if name != nil && *name != "" {
		query = query.Where("users.first_name ILIKE ?", "%"+*name+"%")
	}

	if phone != nil && *phone != "" {
		query = query.Where("users.phone  ILIKE ?", "%"+*phone+"%")
	}

	// Lấy tổng số bản ghi (trước phân trang)
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Thêm sắp xếp, phân trang và lấy dữ liệu
	if err := query.
		Order("users.first_name ASC").
		Offset(offset).
		Limit(*pageSize).
		Find(&result).Error; err != nil {
		return nil, 0, err
	}

	return result, total, nil
}
