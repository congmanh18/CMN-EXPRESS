package auth

import (
	"context"
	error "express_be/core/err"

	"express_be/entity"
)

func (c *authUsecaseImpl) CreateCustomer(ctx context.Context, user *entity.User, customer *entity.Customer) *error.Err {
	// Thêm user trước
	err := c.userRepo.Create(ctx, user)
	if err != nil {
		return error.ErrRegisterUser
	}

	// Sau đó thêm customer
	err = c.customerRepo.Create(ctx, customer)
	if err != nil {
		return error.ErrRegisterUser
	}

	return nil
}
