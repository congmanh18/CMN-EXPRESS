package delivery

import (
	"context"
	"express_be/repository/delivery/entity"
	"express_be/usecase"
)

func (c *adminUseCaseImpl) AdminGetAllDeliveryPersons(ctx context.Context, page *int, pageSize *int) ([]entity.DeliveryPerson, *usecase.Error) {
	deliveryPersons, err := c.deliveryPersonRepo.FetchAllDeliveryPersons(ctx, page, pageSize)
	if err != nil {
		return nil, &usecase.Error{
			Code:    500,
			Message: err.Error(),
			Err:     err,
		}
	}

	return deliveryPersons, nil
}
