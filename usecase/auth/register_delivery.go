package auth

import (
	"context"
	error "express_be/core/err"
	"express_be/entity"
)

func (c *authUsecaseImpl) CreateDeliveryPerson(ctx context.Context, user *entity.User, deliveryPerson *entity.DeliveryPerson) *error.Err {
	// Thêm user trước
	err := c.userRepo.Create(ctx, user)
	if err != nil {
		return error.ErrRegisterUser
	}

	// Sau đó thêm deliveryPerson
	err = c.deliveryPersonRepo.Create(ctx, deliveryPerson)
	if err != nil {
		return error.ErrRegisterUser
	}

	return nil
}
