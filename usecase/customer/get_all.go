package customer

import (
	"context"
	"express_be/repository/customer/entity"
	"express_be/usecase"
)

func (c *adminUsecaseImpl) AdminGetAllCustomers(ctx context.Context, page *int, pageSize *int) ([]entity.Customer, *usecase.Error) {
	customers, err := c.customerRepo.FetchAllCustomer(ctx, page, pageSize)
	if err != nil {
		return nil, &usecase.Error{
			Code:    500,
			Message: err.Error(),
			Err:     err,
		}
	}
	return customers, nil
}
