package user

import (
	"context"
	error "express_be/core/err"
	mapper "express_be/mapper/res"
	"express_be/model/res"
	userEntity "express_be/repository/user/entity"
)

// AdminGetInfoCustomer implements AdminUsecase.
func (c *userUsecaseImpl) GetInfoUser(ctx context.Context, id *string) (*res.CustomerRes, *res.DeliveryPersonRes, *error.Err) {
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
		customers := mapper.CustomerToRes(customerdetail)
		return &customers, nil, nil
	case userEntity.DeliveryPerson:
		deliveryPerson, err := c.deliveryPersonRepo.FindByID(ctx, id)
		if err != nil {
			return nil, nil, error.ErrNotFound
		}
		deliveryPersonDetail := mapper.DeliveryPersonInfoToDetail(user, deliveryPerson)
		deliveryPersons := mapper.DeliveryPersonToRes(deliveryPersonDetail)
		return nil, &deliveryPersons, nil
	}

	return nil, nil, error.ErrInvalidRole
}
