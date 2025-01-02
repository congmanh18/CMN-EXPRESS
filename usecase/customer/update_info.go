package customer

import (
	"context"
	"express_be/repository/customer/entity"
	"express_be/usecase"
)

// UpdateInfoCustomer implements CustomerUsecase.
func (c *customerUsecaseImpl) UpdateInfoCustomer(ctx context.Context, customer *entity.Customer, id *string) *usecase.Error {
	err := c.customerRepo.UpdateCustomer(ctx, id, customer)
	if err != nil {
		return &usecase.Error{
			Code:    500,
			Message: err.Error(),
			Err:     err,
		}
	}

	return nil
}
