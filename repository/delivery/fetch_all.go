package delivery

import (
	"context"
	"express_be/entity"
)

func (d *deliveryImpl) FetchAllDeliveryPersons(ctx context.Context, page, pageSize *int) ([]entity.DeliveryPerson, error) {
	var result []entity.DeliveryPerson
	offset := (*page - 1) * *pageSize

	query := d.DB.Executor.
		WithContext(ctx).
		Order("created_at DESC").
		Offset(offset).
		Limit(*pageSize).
		Find(&result)

	if query.Error != nil {
		return nil, query.Error
	}

	return result, nil
}
