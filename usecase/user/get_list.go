package user

import (
	"context"
	"express_be/repository/user/entity"
	"express_be/usecase"
	"net/http"
)

func (c *userUsecaseImpl) GetUsers(ctx context.Context, status, role *string, page, pageSize *int) ([]entity.CustomerDetails, []entity.DeliveryPersonDetails, *usecase.Error) {
	switch *role {
	case "customer":
		customers, err := c.repo.FetchCustomerUsers(ctx, status, page, pageSize)
		if err != nil {
			return nil, nil, &usecase.Error{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
				Err:     err,
			}
		}
		return customers, nil, nil

	case "delivery_person":
		deliveryPersons, err := c.repo.FetchDeliveryPersonUsers(ctx, status, page, pageSize)
		if err != nil {
			return nil, nil, &usecase.Error{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
				Err:     err,
			}
		}
		return nil, deliveryPersons, nil
	}

	return nil, nil, &usecase.Error{
		Code:    http.StatusBadRequest,
		Message: "Invalid role provided. Role should be either 'customer' or 'delivery_person'",
	}
}
