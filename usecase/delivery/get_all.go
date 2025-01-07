package delivery

import (
	"context"
	"express_be/repository/user/entity"
	"express_be/usecase"
)

func (c *deliveryPersonUsecaseImpl) GetAllDeliveryPersons(ctx context.Context, page *int, pageSize *int) ([]entity.DeliveryPersonDetails, *usecase.Error) {
	deliveryPersons, err := c.userRepo.FetchAllDeliveryPerson(ctx, page, pageSize)
	if err != nil {
		return nil, &usecase.Error{
			Code:    500,
			Message: err.Error(),
			Err:     err,
		}
	}

	return deliveryPersons, nil
}
