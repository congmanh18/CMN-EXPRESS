package user

import (
	"context"
	error "express_be/core/err"
	"express_be/mapper"
	userEntity "express_be/repository/user/entity"
)

// AdminGetInfoCustomer implements AdminUsecase.
func (c *userUsecaseImpl) GetInfoUser(ctx context.Context, id *string) (*userEntity.CustomerDetails, *userEntity.DeliveryPersonDetails, *error.Err) {
	user, err := c.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, nil, error.ErrNotFound

	}
	switch user.Role {
	case userEntity.Customer:
		customer, err := c.customerRepo.FindByID(ctx, id)
		if err != nil {
			return nil, nil, error.ErrNotFound
		}
		customerdetail := mapper.CustomerInfoToDetail(user, customer)
		return customerdetail, nil, nil
	case userEntity.DeliveryPerson:
		deliveryPerson, err := c.deliveryPersonRepo.FindByID(ctx, id)
		if err != nil {
			return nil, nil, error.ErrNotFound
		}
		deliveryPersonDetail := mapper.DeliveryPersonInfoToDetail(user, deliveryPerson)
		return nil, deliveryPersonDetail, nil
	}

	return nil, nil, error.ErrInvalidRole
}
