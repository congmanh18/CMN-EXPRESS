package order

import (
	"context"
	error "express_be/core/err"
	mapper "express_be/mapper/req"
	model "express_be/model/req"
)

// CreateOrder implements OrderUsecase.
func (o *orderUsecaseImpl) CreateOrder(ctx context.Context, req model.CreateOrderReq, customerID *string) *error.Err {
	order := mapper.CreateOrderReqToOrder(req, customerID)
	err := o.orderRepo.Create(ctx, order)
	if err != nil {
		return error.ErrInternalServer
	}

	return nil
}
