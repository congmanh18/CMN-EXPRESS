package auth

import (
	"context"
	"errors"
	deliveryPersonEntity "express_be/repository/delivery/entity"
	"fmt"

	"express_be/usecase"

	"github.com/lib/pq"
)

func (r *authUsecaseImpl) CreateDeliveryPersonUsecase(ctx context.Context, deliveryPerson *deliveryPersonEntity.DeliveryPerson) *usecase.Error {
	err := r.deliveryPersonRepo.CreateDeliveryPerson(ctx, deliveryPerson)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) && pqErr.Code == "23505" { // SQLSTATE 23505: duplicate key
			return &usecase.Error{
				Code:    400,
				Message: fmt.Sprintf("Delivery person already exists: %s", *deliveryPerson.Phone),
				Err:     errors.New("delivery person already exists"),
			}
		}

		return &usecase.Error{
			Code:    500,
			Message: "Failed to create delivery person -" + err.Error(),
			Err:     err,
		}
	}

	return nil
}
