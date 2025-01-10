package auth

import (
	"context"
	error "express_be/core/err"
	deliveryEntity "express_be/repository/delivery/entity"
	user "express_be/repository/user/entity"
)

func (c *authUsecaseImpl) CreateDeliveryPerson(ctx context.Context, user *user.User, deliveryPerson *deliveryEntity.DeliveryPerson) *error.Err {
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
