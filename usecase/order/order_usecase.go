package order

import (
	"context"
	error "express_be/core/err"
	model "express_be/model/req"
	"express_be/repository/order"
)

type OrderUsecase interface {
	CreateOrder(ctx context.Context, req model.CreateOrderReq) *error.Err
}

type orderUsecaseImpl struct {
	orderRepo order.Repo
}

func NewOrderUsecase(orderRepo order.Repo) OrderUsecase {
	return &orderUsecaseImpl{orderRepo: orderRepo}
}
