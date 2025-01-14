package user

import (
	"context"
	error "express_be/core/err"
	deliveryPersonEntity "express_be/repository/delivery/entity"
	userEntity "express_be/repository/user/entity"
)

func (c *userUsecaseImpl) UpdateDeliveryPersonInfo(ctx context.Context, id *string, user *userEntity.User, deliveryPersonEntity *deliveryPersonEntity.DeliveryPerson) *error.Err {
	err := c.userRepo.Update(ctx, id, user)
	if err != nil {
		return error.ErrInternalServer
	}

	err = c.deliveryPersonRepo.Update(ctx, id, deliveryPersonEntity)
	if err != nil {
		return error.ErrInternalServer
	}

	return nil
}
