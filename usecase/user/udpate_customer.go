package user

import (
	"context"
	error "express_be/core/err"
	customerEntity "express_be/repository/customer/entity"
	userEntity "express_be/repository/user/entity"
)

func (c *userUsecaseImpl) UpdateCustomerInfo(ctx context.Context, id *string, user *userEntity.User, customer *customerEntity.Customer) *error.Err {
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
