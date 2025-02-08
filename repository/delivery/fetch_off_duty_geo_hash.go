package delivery

import (
	"context"
	"express_be/entity"
)

func (d *deliveryImpl) FetchOffDutyGeoHashes(ctx context.Context, status *string) ([]string, error) {
	var geoHashes []string

	query := d.DB.Executor.
		WithContext(ctx).
		Model(&entity.DeliveryPerson{}). // Chỉ lấy cột geo_hash
		Where("active_status = ?", *status).
		Pluck("geo_hash", &geoHashes) // Lấy danh sách geo_hash

	if query.Error != nil {
		return nil, query.Error
	}

	return geoHashes, nil
}
