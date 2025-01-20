package user

import (
	"context"
	error "express_be/core/err"
	"express_be/entity"
)

// Search implements UserUsecase.
func (c *userUsecaseImpl) Search(ctx context.Context, role, phone *string, name *string, status *string, page *int, pageSize *int) ([]entity.CustomerDetails, []entity.DeliveryPersonDetails, *error.Err) {
	switch *role {
	case "customer":
		customers, _, err := c.userRepo.SearchCustomers(ctx, phone, name, status, page, pageSize)
		if err != nil {
			return nil, nil, error.ErrInternalServer
		}
		return customers, nil, nil

	case "delivery_person":
		deliveryPersons, _, err := c.userRepo.SearchDeliveryPersons(ctx, phone, name, status, page, pageSize)
		if err != nil {
			return nil, nil, error.ErrInternalServer
		}
		return nil, deliveryPersons, nil
	}

	return nil, nil, error.ErrInvalidRole
}
