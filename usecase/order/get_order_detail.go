package order

import (
	"context"
	error "express_be/core/err"
	mapper "express_be/mapper/res"

	"express_be/model/res"
)

func (o *orderUsecaseImpl) GetOrderDetail(ctx context.Context, id *string) (*res.OrderRes, *error.Err) {
	order, err := o.orderRepo.FindByID(ctx, id)
	if err != nil {
		return nil, error.ErrInternalServer
	}
	orderResponse := mapper.OrderDetailToRes(order)

	return &orderResponse, nil
}
