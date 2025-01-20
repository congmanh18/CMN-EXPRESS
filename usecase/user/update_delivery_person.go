package user

import (
	"context"
	error "express_be/core/err"
	"express_be/entity"
)

func (c *userUsecaseImpl) UpdateDeliveryPersonInfo(ctx context.Context, id *string, user *entity.User, deliveryPersonEntity *entity.DeliveryPerson) *error.Err {
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
