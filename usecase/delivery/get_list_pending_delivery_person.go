package delivery

import (
	"context"
	"express_be/repository/user/entity"
	"express_be/usecase"
	"net/http"
)

func (c *deliveryPersonUsecaseImpl) GetPendingDeliveryPersons(ctx context.Context, page, pageSize *int) ([]entity.DeliveryPersonDetails, *usecase.Error) {
	deliveryPersons, err := c.userRepo.FetchPendingStatusDeliveryPerson(ctx, page, pageSize)
	if err != nil {
		return nil, &usecase.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Err:     err,
		}
	}

	return deliveryPersons, nil
}
