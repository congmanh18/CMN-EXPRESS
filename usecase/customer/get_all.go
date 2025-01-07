package customer

import (
	"context"
	"express_be/repository/user/entity"
	"express_be/usecase"
)

func (c *customerUsecaseImpl) GetAllCustomers(ctx context.Context, page *int, pageSize *int) ([]entity.CustomerDetails, *usecase.Error) {
	customers, err := c.userRepo.FetchAllCustomer(ctx, page, pageSize)
	if err != nil {
		return nil, &usecase.Error{
			Code:    500,
			Message: err.Error(),
			Err:     err,
		}
	}
	return customers, nil
}
