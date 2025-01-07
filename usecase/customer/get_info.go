package customer

import (
	"context"
	"express_be/mapper"
	userEntity "express_be/repository/user/entity"
	"express_be/usecase"
)

// AdminGetInfoCustomer implements AdminUsecase.
func (c *customerUsecaseImpl) GetInfoCustomer(ctx context.Context, id *string) (*userEntity.CustomerDetails, *usecase.Error) {
	user, err := c.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, &usecase.Error{
			Code:    404,
			Message: "User information not found. Please check the provided ID. " + err.Error(),
			Err:     err,
		}
	}
	customer, err := c.customerRepo.FindByID(ctx, id)
	if err != nil {
		return nil, &usecase.Error{
			Code:    404,
			Message: "Customer information not found. Please check the provided ID. " + err.Error(),
			Err:     err,
		}
	}
	customerdetail := mapper.CustomerInfoToDetail(user, customer)

	return customerdetail, nil
}
