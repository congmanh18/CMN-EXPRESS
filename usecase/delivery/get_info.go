package delivery

import (
	"context"
	"express_be/mapper"
	userEntity "express_be/repository/user/entity"
	"express_be/usecase"
)

func (d *deliveryPersonUsecaseImpl) GetInfoDeliveryPerson(ctx context.Context, id *string) (*userEntity.DeliveryPersonDetails, *usecase.Error) {
	user, err := d.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, &usecase.Error{
			Code:    500,
			Message: "User information not found. Please check the provided ID. " + err.Error(),
			Err:     err,
		}
	}
	deliveryPerson, err := d.deliveryPersonRepo.FindByID(ctx, id)
	if err != nil {
		return nil, &usecase.Error{
			Code:    500,
			Message: "Customer information not found. Please check the provided ID. " + err.Error(),
			Err:     err,
		}
	}
	deliveryPersonDetail := mapper.DeliveryPersonInfoToDetail(user, deliveryPerson)
	return deliveryPersonDetail, nil
}
