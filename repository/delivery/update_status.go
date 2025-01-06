package delivery

import (
	"context"
	"express_be/repository/delivery/entity"
	"fmt"
)

func (d *deliveryImpl) UpdateStatus(ctx context.Context, id *string, approval_status *entity.ApprovalStatus, status *entity.Status) error {
	result := d.DB.Executor.WithContext(ctx).
		Model(&entity.DeliveryPerson{}).
		Where("id = ?", *id).
		Updates(map[string]interface{}{
			"approval_status": approval_status,
			"status":          status,
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no customer found with id: %s", *id)
	}

	return nil
}
