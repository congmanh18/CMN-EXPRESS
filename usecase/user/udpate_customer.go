package user

import (
	"context"
	error "express_be/core/err"
	"express_be/entity"
)

func (c *userUsecaseImpl) UpdateCustomerInfo(ctx context.Context, id *string, user *entity.User, customer *entity.Customer) *error.Err {
	err := c.userRepo.Update(ctx, id, user)
	if err != nil {
		return error.ErrInternalServer
	}

	err = c.customerRepo.Update(ctx, id, customer)
	if err != nil {
		return error.ErrInternalServer
	}

	return nil
}
