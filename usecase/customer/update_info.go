package customer

import (
	"context"
	customerEntity "express_be/repository/customer/entity"
	userEntity "express_be/repository/user/entity"

	"express_be/usecase"
)

// UpdateInfoCustomer implements CustomerUsecase.
func (c *customerUsecaseImpl) UpdateInfoCustomer(ctx context.Context, user *userEntity.User, customer *customerEntity.Customer, id *string) *usecase.Error {
	err := c.userRepo.Update(ctx, id, user)
	if err != nil {
		return &usecase.Error{
			Code:    500,
			Message: "Failed to update user information. Please try again later. " + err.Error(),
			Err:     err,
		}
	}

	err = c.customerRepo.UpdateCustomer(ctx, id, customer)
	if err != nil {
		return &usecase.Error{
			Code:    500,
			Message: "Failed to update customer information. Please try again later. " + err.Error(),
			Err:     err,
		}
	}

	return nil
}
