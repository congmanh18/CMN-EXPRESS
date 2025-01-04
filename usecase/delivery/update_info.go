package delivery

import (
	"context"
	"express_be/repository/delivery/entity"
	"express_be/usecase"
)

func (d *deliveryPersonUsecaseImpl) UpdateInfoDeliveryPerson(ctx context.Context, deliveryPerson *entity.DeliveryPerson, id *string) *usecase.Error {
	err := d.deliveryPersonRepo.UpdateDeliveryPerson(ctx, id, deliveryPerson)
	if err != nil {
		return &usecase.Error{
			Code:    500,
			Message: err.Error(),
			Err:     err,
		}
	}

	return nil
}
