package customer

import (
	"context"
	"express_be/repository/user/entity"
	"express_be/usecase"
)

func (c *customerUsecaseImpl) GetPendingCustomers(ctx context.Context, page, pageSize *int) ([]entity.CustomerDetails, *usecase.Error) {
	customers, err := c.userRepo.FetchPendingStatusCustomers(ctx, page, pageSize)
	if err != nil {
		return nil, &usecase.Error{
			Code:    500,
			Message: err.Error(),
			Err:     err,
		}
	}

	return customers, nil
}
