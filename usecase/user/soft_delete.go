package user

import (
	"context"
	"express_be/usecase"
)

func (c *userUsecaseImpl) SoftDeleteCustomer(ctx context.Context, id *string) *usecase.Error {
	err := c.customerRepo.DeleteCustomer(ctx, id)
	if err != nil {
		return &usecase.Error{
			Code:    500,
			Message: "Failed to delete customer" + err.Error(),
			Err:     err,
		}
	}

	return nil
}
