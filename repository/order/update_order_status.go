package order

import (
	"context"
	"express_be/repository/order/entity"
	"fmt"
)

// UpdateOrderStatus implements Repo.
func (r *orderImpl) UpdateOrderStatus(ctx context.Context, id *string, status *entity.Status) error {
	result := r.DB.Executor.WithContext(ctx).
		Model(&entity.Order{}).
		Where("id = ?", *id).
		Update("status", status)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no order found with id: %s", *id)
	}

	return nil // Thành công
}
