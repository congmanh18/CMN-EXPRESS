package order

import (
	"context"
	error "express_be/core/err"
	"express_be/entity"
)

var validStatuses = map[string]entity.Status{
	string(entity.Created):              entity.Created,
	string(entity.PickingUp):            entity.PickingUp,
	string(entity.PickedUp):             entity.PickedUp,
	string(entity.PickupCanceled):       entity.PickupCanceled,
	string(entity.ReturnToHub):          entity.ReturnToHub,
	string(entity.AtHub):                entity.AtHub,
	string(entity.OutForDelivery):       entity.OutForDelivery,
	string(entity.Delivered):            entity.Delivered,
	string(entity.Processing):           entity.Processing,
	string(entity.ReattemptDelivery):    entity.ReattemptDelivery,
	string(entity.Returning):            entity.Returning,
	string(entity.Transferred):          entity.Transferred,
	string(entity.Refunded):             entity.Refunded,
	string(entity.Canceled):             entity.Canceled,
	string(entity.CompensationChecking): entity.CompensationChecking,
	string(entity.Compensated):          entity.Compensated,
}

func (o *orderUsecaseImpl) UpdateOrderStatus(ctx context.Context, id *string, status *string) *error.Err {
	orderStatus, exists := validStatuses[*status]
	if !exists {
		return error.ErrInvalidFormat
	}

	err := o.orderRepo.UpdateOrderStatus(ctx, id, &orderStatus)
	if err != nil {
		return error.ErrInternalServer
	}

	return nil
}
