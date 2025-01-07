package auth

import (
	"context"
	deliveryPersonEntity "express_be/repository/delivery/entity"
	user "express_be/repository/user/entity"

	"express_be/usecase"
)

func (c *authUsecaseImpl) CreateDeliveryPerson(ctx context.Context, user *user.User, deliveryPerson *deliveryPersonEntity.DeliveryPerson) *usecase.Error {
	err := c.userRepo.Create(ctx, user)
	if err != nil {
		return &usecase.Error{
			Code:    500,
			Message: "Failed to create user account. Please try again later." + err.Error(),
			Err:     err,
		}
	}

	err = c.deliveryPersonRepo.CreateDeliveryPerson(ctx, deliveryPerson)
	if err != nil {
		return &usecase.Error{
			Code:    500,
			Message: "Failed to create accounting record. Please try again later. " + err.Error(),
			Err:     err,
		}
	}

	return nil
}
