package delivery

import (
	"context"
	"express_be/repository/delivery/entity"
)

func (d *deliveryImpl) FindByPhone(ctx context.Context, phone *string) (*entity.DeliveryPerson, error) {
	var result *entity.DeliveryPerson
	query := d.DB.Executor.WithContext(ctx).
		Model(&entity.DeliveryPerson{}).
		Where("phone = ?", phone).
		First(&result)

	if query.Error != nil {
		return nil, query.Error
	}

	return result, nil
}
