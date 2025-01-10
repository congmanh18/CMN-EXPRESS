package delivery

import (
	"context"
	"express_be/repository/delivery/entity"
)

func (d *deliveryImpl) FetchPhone(ctx context.Context, phone *string) (*string, error) {
	var result string
	query := d.DB.Executor.WithContext(ctx).
		Model(&entity.DeliveryPerson{}).
		Select("phone").
		Where("phone = ?", phone).
		First(&result)

	if query.Error != nil {
		return nil, query.Error
	}

	if result == "" {
		return nil, nil
	}

	return &result, nil
}
