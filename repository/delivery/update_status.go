package delivery

import (
	"context"
	"errors"
	"express_be/repository/delivery/entity"
)

func (d *deliveryImpl) UpdateStatus(ctx context.Context, id *string, status *entity.Status) error {
	if id == nil {
		return errors.New("id cannot be nil")
	}

	err := d.DB.Executor.WithContext(ctx).
		Model(&entity.DeliveryPerson{}).
		Where("id = ?", id).
		Update("status", status).Error

	if err != nil {
		return err
	}

	return nil
}
