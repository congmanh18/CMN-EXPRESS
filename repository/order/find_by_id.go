package order

import (
	"context"
	"express_be/repository/order/entity"
)

func (r *orderImpl) FindByID(ctx context.Context, id *string) (*entity.OrderDetail, error) {
	var result entity.OrderDetail

	// Thực hiện truy vấn
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
		Where("orders.id = ?", id).
		First(&result)

	if query.Error != nil {
		return nil, query.Error
	}

	return &result, nil
}
