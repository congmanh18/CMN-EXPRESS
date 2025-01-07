package auth

import (
	"context"
	deliveryEntity "express_be/repository/delivery/entity"
	user "express_be/repository/user/entity"

	"express_be/usecase"
)

func (c *authUsecaseImpl) CreateDeliveryPerson(ctx context.Context, user *user.User, deliveryPerson *deliveryEntity.DeliveryPerson) *usecase.Error {
	// Thêm user trước
	err := c.userRepo.Create(ctx, user)
	if err != nil {
		return &usecase.Error{
			Code:    500,
			Message: "Failed to create user. Please try again later. " + err.Error(),
			Err:     err,
		}
	}

	// Sau đó thêm deliveryPerson
	err = c.deliveryPersonRepo.Create(ctx, deliveryPerson)
	if err != nil {
		return &usecase.Error{
			Code:    500,
			Message: "Failed to create delivery person. Please try again later. " + err.Error(),
			Err:     err,
		}
	}

	return nil
}
