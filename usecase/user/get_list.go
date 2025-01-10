package user

import (
	"context"
	error "express_be/core/err"
	"express_be/repository/user/entity"
)

func (c *userUsecaseImpl) GetUsers(ctx context.Context, status, role *string, page, pageSize *int) ([]entity.CustomerDetails, []entity.DeliveryPersonDetails, *error.Err) {
	switch *role {
	case "customer":
		customers, err := c.userRepo.FetchCustomerUsers(ctx, status, page, pageSize)
		if err != nil {
			return nil, nil, error.ErrInternalServer
		}
		return customers, nil, nil

	case "delivery_person":
		deliveryPersons, err := c.userRepo.FetchDeliveryPersonUsers(ctx, status, page, pageSize)
		if err != nil {
			return nil, nil, error.ErrInternalServer
		}
		return nil, deliveryPersons, nil
	}

	return nil, nil, error.ErrInvalidRole
}
