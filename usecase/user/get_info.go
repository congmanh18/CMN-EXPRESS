package user

import (
	"context"
	"express_be/mapper"
	userEntity "express_be/repository/user/entity"
	"express_be/usecase"
)

// AdminGetInfoCustomer implements AdminUsecase.
func (c *userUsecaseImpl) GetInfoUser(ctx context.Context, id *string) (*userEntity.CustomerDetails, *userEntity.DeliveryPersonDetails, *usecase.Error) {
	user, err := c.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, nil, &usecase.Error{
			Code:    404,
			Message: "User information not found. Please check the provided ID. " + err.Error(),
			Err:     err,
		}
	}
	switch user.Role {
	case userEntity.Customer:
		customer, err := c.customerRepo.FindByID(ctx, id)
		if err != nil {
			return nil, nil, &usecase.Error{
				Code:    404,
				Message: "Customer information not found. Please check the provided ID. " + err.Error(),
				Err:     err,
			}
		}
		customerdetail := mapper.CustomerInfoToDetail(user, customer)
		return customerdetail, nil, nil
	case userEntity.DeliveryPerson:
		deliveryPerson, err := c.deliveryPersonRepo.FindByID(ctx, id)
		if err != nil {
			return nil, nil, &usecase.Error{
				Code:    404,
				Message: "Delivery person information not found. Please check the provided ID. " + err.Error(),
				Err:     err,
			}
		}
		deliveryPersonDetail := mapper.DeliveryPersonInfoToDetail(user, deliveryPerson)
		return nil, deliveryPersonDetail, nil
	}

	return nil, nil, &usecase.Error{
		Code:    400,
		Message: "Invalid role. Please provide 'customer' or 'delivery_person'.",
	}
}
