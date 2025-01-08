package user

import (
	"context"
	"express_be/repository/user/entity"
	"express_be/usecase"
)

func (c *userUsecaseImpl) GetUsers(ctx context.Context, status, role *string, page, pageSize *int) ([]entity.CustomerDetails, []entity.DeliveryPersonDetails, *usecase.Error) {
	switch *role {
	case "customer":
		customers, err := c.userRepo.FetchCustomerUsers(ctx, status, page, pageSize)
		if err != nil {
			return nil, nil, &usecase.Error{
				Code:    500,
				Message: err.Error(),
				Err:     err,
			}
		}
		return customers, nil, nil

	case "delivery_person":
		deliveryPersons, err := c.userRepo.FetchDeliveryPersonUsers(ctx, status, page, pageSize)
		if err != nil {
			return nil, nil, &usecase.Error{
				Code:    500,
				Message: err.Error(),
				Err:     err,
			}
		}
		return nil, deliveryPersons, nil
	}

	return nil, nil, &usecase.Error{
		Code:    400,
		Message: "Invalid role. Please provide 'customer' or 'delivery_person'.",
	}
}
