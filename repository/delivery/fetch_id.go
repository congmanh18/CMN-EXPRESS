package delivery

import (
	"context"
	"express_be/entity"
)

func (d *deliveryImpl) FetchID(ctx context.Context, id *string) (*string, error) {
	var result string
	query := d.DB.Executor.WithContext(ctx).
		Model(&entity.DeliveryPerson{}).
		Select("id").
		Where("id = ?", id).
		First(&result)

	if query.Error != nil {
		return nil, query.Error
	}

	return &result, nil
}
