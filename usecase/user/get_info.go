package user

import (
	"context"
	error "express_be/core/err"
	"express_be/entity"
	mapper "express_be/mapper/res"
	"express_be/model/res"
)

// AdminGetInfoCustomer implements AdminUsecase.
func (c *userUsecaseImpl) GetInfoUser(ctx context.Context, id *string) (*res.CustomerRes, *res.DeliveryPersonRes, *error.Err) {
	user, err := c.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, nil, error.ErrNotFound

	}
	switch user.Role {
	case entity.CustomerRole:
		customer, err := c.customerRepo.FindByID(ctx, id)
		if err != nil {
			return nil, nil, error.ErrNotFound
		}
		customerdetail := mapper.CustomerInfoToDetail(user, customer)
		customers := mapper.CustomerToRes(customerdetail)
		return &customers, nil, nil
	case entity.DeliveryPersonRole:
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
