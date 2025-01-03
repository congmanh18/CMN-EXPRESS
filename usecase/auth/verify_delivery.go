package auth

import (
	"context"
	"errors"
	deliveryEntity "express_be/repository/delivery/entity"
	"express_be/usecase"
)

func (d *authUsecaseImpl) VerifyDeliveryPersonUsecase(ctx context.Context, uid *string) *usecase.Error {
	exists, err := d.deliveryPersonRepo.FetchID(ctx, uid)
	if err != nil {
		return &usecase.Error{
			Code:    400,
			Message: "Delivery person not found",
			Err:     err,
		}
	}
	if exists != nil {
		return &usecase.Error{
			Code:    400,
			Message: "Delivery person already verified",
			Err:     errors.New("delivery person already verified"),
		}
	}

	status := deliveryEntity.Verified
	err = d.deliveryPersonRepo.UpdateStatus(ctx, uid, &status)
	if err != nil {
		return &usecase.Error{
			Code:    500,
			Message: "Failed to verify delivery person",
			Err:     err,
		}
	}

	return nil
}
