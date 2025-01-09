package user

import (
	"context"
	"express_be/repository/user/entity"
	"express_be/usecase"
)

// Search implements UserUsecase.
func (c *userUsecaseImpl) Search(ctx context.Context, role, phone *string, name *string, status *string, page *int, pageSize *int) ([]entity.CustomerDetails, []entity.DeliveryPersonDetails, *usecase.Error) {
	switch *role {
	case "customer":
		customers, _, err := c.userRepo.SearchCustomers(ctx, phone, name, status, page, pageSize)
		if err != nil {
			return nil, nil, &usecase.Error{
				Code:    500,
				Message: err.Error(),
				Err:     err,
			}
		}
		return customers, nil, nil

	case "delivery_person":
		deliveryPersons, _, err := c.userRepo.SearchDeliveryPersons(ctx, phone, name, status, page, pageSize)
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
