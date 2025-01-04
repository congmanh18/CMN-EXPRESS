package delivery

import (
	"context"
	"express_be/repository/delivery/entity"
	"fmt"
)

func (d *deliveryImpl) UpdateStatus(ctx context.Context, id *string, status *entity.Status) error {
	result := d.DB.Executor.WithContext(ctx).
		Model(&entity.DeliveryPerson{}).
		Where("id = ?", *id).
		Update("status", status)

	// Kiểm tra lỗi truy vấn
	if result.Error != nil {
		return result.Error // Trả lỗi nếu có vấn đề
	}
	// Kiểm tra nếu không có hàng nào bị ảnh hưởng
	if result.RowsAffected == 0 {
		return fmt.Errorf("no customer found with id: %s", *id)
	}

	return nil
}
