package delivery

import (
	"context"
	"express_be/entity"
	"fmt"
)

func (r *deliveryImpl) Create(ctx context.Context, deliveryPerson *entity.DeliveryPerson) error {
	if err := r.DB.Executor.WithContext(ctx).Create(deliveryPerson).Error; err != nil {
		return fmt.Errorf("failed to create delivery person: %w", err)
	}
	return nil
}
