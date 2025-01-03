package delivery

import (
	"context"
	"express_be/repository/delivery/entity"
	"fmt"
)

func (d *deliveryImpl) UpdateDeliveryPerson(ctx context.Context, id *string, deliveryPerson *entity.DeliveryPerson) error {
	err := d.DB.Executor.WithContext(ctx).
		Model(&entity.DeliveryPerson{}).
		Where("id =?", id).
		Updates(deliveryPerson).Error

	if err != nil {
		return fmt.Errorf("failed to update delivery person (ID: %d): %w", id, err)
	}

	return nil

}
