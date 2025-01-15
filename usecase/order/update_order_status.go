package order

import (
	"context"
	error "express_be/core/err"
	"express_be/repository/order/entity"
)

var validStatuses = map[string]entity.Status{
	string(entity.Created):          entity.Created,
	string(entity.PickingUp):        entity.PickingUp,
	string(entity.PickedUp):         entity.PickedUp,
	string(entity.PickupCanceled):   entity.PickupCanceled,
	string(entity.ReturnToHub):      entity.ReturnToHub,
	string(entity.AtHub):            entity.AtHub,
	string(entity.OutForDelivery):   entity.OutForDelivery,
	string(entity.Delivered):        entity.Delivered,
	string(entity.Returned):         entity.Returned,
	string(entity.DeliveryCanceled): entity.DeliveryCanceled,
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
