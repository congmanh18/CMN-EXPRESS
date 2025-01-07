package delivery

import (
	"context"
	"express_be/repository/delivery/entity"
	userEntity "express_be/repository/user/entity"
	"express_be/usecase"
)

func (d *deliveryPersonUsecaseImpl) UpdateInfoDeliveryPerson(ctx context.Context, userEntity *userEntity.User, deliveryPersonEntity *entity.DeliveryPerson, id *string) *usecase.Error {
	err := d.userRepo.Update(ctx, id, userEntity)
	if err != nil {
		return &usecase.Error{
			Code:    500,
			Message: "Failed to update user information. Please try again later. " + err.Error(),
			Err:     err,
		}
	}
	err = d.deliveryPersonRepo.UpdateDeliveryPerson(ctx, id, deliveryPersonEntity)
	if err != nil {
		return &usecase.Error{
			Code:    500,
			Message: "Failed to update delivery person information. Please try again later. " + err.Error(),
			Err:     err,
		}
	}

	return nil
}
