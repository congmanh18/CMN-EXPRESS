package delivery

import (
	"context"
	"express_be/core/gorm"
	"express_be/repository/delivery/entity"
	"fmt"
	"reflect"
)

func (d *deliveryImpl) UpdateDeliveryPerson(ctx context.Context, id *string, deliveryPerson *entity.DeliveryPerson) error {
	omitFields := gorm.OmitFields(deliveryPerson, func(fieldValue reflect.Value) bool {
		return fieldValue.IsZero()
	})

	err := d.DB.Executor.WithContext(ctx).
		Model(&entity.DeliveryPerson{}).
		Where("id =?", *id).
		Omit(omitFields...).
		Updates(deliveryPerson).Error

	if err != nil {
		return fmt.Errorf("failed to update delivery person (ID: %d): %w", id, err)
	}

	return nil

}
