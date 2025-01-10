package delivery

import (
	"context"
	deliveryPersonEntity "express_be/repository/delivery/entity"
	"fmt"
)

func (r *deliveryImpl) Create(ctx context.Context, deliveryPerson *deliveryPersonEntity.DeliveryPerson) error {
	if err := r.DB.Executor.WithContext(ctx).Create(deliveryPerson).Error; err != nil {
		return fmt.Errorf("failed to create delivery person: %w", err)
	}
	return nil
}
