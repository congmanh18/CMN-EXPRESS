package delivery

import (
	"context"
	"express_be/repository/delivery/entity"
)

func (d *deliveryImpl) CreateDeliveryPerson(ctx context.Context, delivery *entity.DeliveryPerson) error {
	return d.DB.Executor.WithContext(ctx).
		Debug().
		Create(&delivery).Error
}
