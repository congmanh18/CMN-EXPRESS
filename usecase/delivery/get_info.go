package delivery

import (
	"context"
	"express_be/repository/delivery/entity"
	"express_be/usecase"
)

func (d *deliveryPersonUsecaseImpl) GetInfoDeliveryPerson(ctx context.Context, id *string) (*entity.DeliveryPerson, *usecase.Error) {
	deliveryPerson, err := d.deliveryPersonRepo.FindByID(ctx, id)
	if err != nil {
		return nil, &usecase.Error{
			Code:    500,
			Message: err.Error(),
			Err:     err,
		}
	}

	return deliveryPerson, nil
}
