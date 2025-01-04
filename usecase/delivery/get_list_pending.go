package delivery

import (
	"context"
	"express_be/repository/delivery/entity"
	"express_be/usecase"
	"net/http"
)

func (c *adminUseCaseImpl) AdminGetPendingDeliveryPersons(ctx context.Context, page, pageSize *int) ([]entity.DeliveryPerson, *usecase.Error) {
	deliveryPersons, err := c.deliveryPersonRepo.FetchPendingStatusDeliveryPerson(ctx, page, pageSize)
	if err != nil {
		return nil, &usecase.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Err:     err,
		}
	}

	return deliveryPersons, nil
}
