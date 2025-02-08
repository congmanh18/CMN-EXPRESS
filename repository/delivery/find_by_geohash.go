package delivery

import (
	"context"
	"errors"
	"express_be/entity"

	"gorm.io/gorm"
)

func (d *deliveryImpl) FindByGeohash(ctx context.Context, geohash *string) (*entity.DeliveryPerson, error) {
	var deliveryPerson entity.DeliveryPerson

	query := d.DB.Executor.
		WithContext(ctx).
		Where("geo_hash = ?", *geohash).
		First(&deliveryPerson)

	if query.Error != nil {
		if errors.Is(query.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("delivery person not found")
		}
		return nil, query.Error
	}

	return &deliveryPerson, nil
}
