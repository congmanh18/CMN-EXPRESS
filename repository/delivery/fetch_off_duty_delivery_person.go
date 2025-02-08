package delivery

import (
	"context"
	"express_be/entity"
)

func (d *deliveryImpl) FetchOffDutyDeliveryPersons(ctx context.Context) ([]entity.DeliveryPerson, error) {
	var result []entity.DeliveryPerson

	query := d.DB.Executor.
		WithContext(ctx).
		Where("active_status = ?", "off_duty"). // Lọc shipper có trạng thái off_duty
		Order("created_at DESC").
		Find(&result)

	if query.Error != nil {
		return nil, query.Error
	}

	return result, nil
}
