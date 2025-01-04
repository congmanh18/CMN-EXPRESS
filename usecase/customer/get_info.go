package customer

import (
	"context"
	"express_be/repository/customer/entity"
	"express_be/usecase"
)

// AdminGetInfoCustomer implements AdminUsecase.
func (c *customerUsecaseImpl) GetInfoCustomer(ctx context.Context, id *string) (*entity.Customer, *usecase.Error) {
	customer, err := c.customerRepo.FindByID(ctx, id)
	if err != nil {
		return nil, &usecase.Error{
			Code:    500,
			Message: err.Error(),
			Err:     err,
		}
	}

	return customer, nil
}
