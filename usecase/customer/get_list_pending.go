package customer

import (
	"context"
	"express_be/repository/customer/entity"
	"express_be/usecase"
)

func (c *adminUsecaseImpl) AdminGetPendingCustomers(ctx context.Context, page, pageSize *int) ([]entity.Customer, *usecase.Error) {
	customers, err := c.customerRepo.FetchPendingStatusCustomers(ctx, page, pageSize)
	if err != nil {
		return nil, &usecase.Error{
			Code:    500,
			Message: err.Error(),
			Err:     err,
		}
	}

	return customers, nil
}
