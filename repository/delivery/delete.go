package delivery

import (
	"context"
	"express_be/entity"
)

func (d *deliveryImpl) DeleteDeliveryPerson(ctx context.Context, id *string) error {
	return d.DB.Executor.WithContext(ctx).
		Where("id = ?", id).
		Delete(&entity.DeliveryPerson{}).
		Error
}
