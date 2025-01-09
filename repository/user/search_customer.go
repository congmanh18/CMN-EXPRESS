package user

import (
	"context"
	"express_be/repository/user/entity"
)

func (u *userImpl) SearchCustomers(ctx context.Context, phone, name, status *string, page, pageSize *int) ([]entity.CustomerDetails, int64, error) {
	var result []entity.CustomerDetails
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
	query := u.DB.Executor.WithContext(ctx).
		Table("users").
		Select(`
			users.*,
			customers.account_type,
			customers.latitude,
			customers.longtitude
		`).
		Joins("JOIN customers ON customers.phone = users.phone").
		Where("users.role = ?", entity.Customer)

	// Thêm điều kiện nếu `status` được truyền
	if status != nil && *status != "" {
		query = query.Where("users.status = ?", *status)
	}

	// Thêm điều kiện nếu `name` được truyền
	if name != nil && *name != "" {
		query = query.Where("users.name ILIKE ?", "%"+*name+"%")
	}

	if phone != nil && *phone != "" {
		query = query.Where("users.phone =?", *phone)
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
