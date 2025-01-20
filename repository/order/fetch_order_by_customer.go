package order

import (
	"context"
	"express_be/entity"
)

// FetchOrderByCustomer implements Repo.
func (r *orderImpl) FetchOrderByCustomer(ctx context.Context, sender_id *string, page *int, pageSize *int) ([]entity.OrderDetail, error) {
	var result []entity.OrderDetail

	defaultPage, defaultPageSize, maxPageSize := 1, 10, 100
	if page == nil || *page < 1 {
		page = &defaultPage
	}

	if pageSize == nil || *pageSize < 1 || *pageSize > maxPageSize {
		pageSize = &defaultPageSize
	}

	offset := (*page - 1) * *pageSize
	query := r.DB.Executor.WithContext(ctx).
		Table("orders").
		Select(`
			orders.*,
			picker.phone AS picker_phone,
			CONCAT(picker.first_name, ' ', picker.middle_name, ' ', picker.last_name) AS picker_fullname,
			deliverer.phone AS deliverer_phone,
			CONCAT(deliverer.first_name, ' ', deliverer.middle_name, ' ', deliverer.last_name) AS deliverer_fullname
		`).
		Joins(`LEFT JOIN users AS picker ON orders.picker_id = picker.id`).
		Joins(`LEFT JOIN users AS deliverer ON orders.deliverer_id = deliverer.id`).
		Where("orders.sender_id = ?", sender_id).
		Order("orders.created_at DESC").
		Offset(offset).
		Limit(*pageSize).
		Find(&result)

	if query.Error != nil {
		return nil, query.Error
	}

	return result, nil
}
