package delivery

import (
	"context"
	"express_be/repository/delivery/entity"
	"fmt"
)

// Sử dụng giao dịch để đảm bảo tính toàn vẹn dữ liệu
func (d *deliveryImpl) CreateDeliveryPerson(ctx context.Context, delivery *entity.DeliveryPerson) error {
	tx := d.DB.Executor.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Kiểm tra xem số điện thoại đã tồn tại chưa
	var count int64
	if err := tx.WithContext(ctx).
		Model(&entity.DeliveryPerson{}).
		Where("phone = ?", delivery.Phone).
		Count(&count).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to check existing phone: %w", err)
	}

	if count > 0 {
		tx.Rollback()
		return fmt.Errorf("phone existing: %s", *delivery.Phone)
	}

	// Tạo mới DeliveryPerson
	if err := tx.WithContext(ctx).
		Create(&delivery).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to create delivery person: %w", err)
	}

	// Commit giao dịch
	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
