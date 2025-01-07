package auth

import (
	"context"
	customerEntity "express_be/repository/customer/entity"
	user "express_be/repository/user/entity"

	"express_be/usecase"
)

func (c *authUsecaseImpl) CreateCustomer(ctx context.Context, user *user.User, customer *customerEntity.Customer) *usecase.Error {
	err := c.userRepo.Create(ctx, user)
	if err != nil {
		return &usecase.Error{
			Code:    500,
			Message: "Failed to create user account. Please try again later. " + err.Error(),
			Err:     err,
		}
	}

	err = c.customerRepo.CreateCustomer(ctx, customer)
	if err != nil {
		return &usecase.Error{
			Code:    500,
			Message: "Failed to create accounting record. Please try again later. " + err.Error(),
			Err:     err,
		}
	}
	return nil
}
