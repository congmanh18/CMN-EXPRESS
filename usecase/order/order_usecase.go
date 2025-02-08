package order

import (
	"context"
	error "express_be/core/err"
	"express_be/model/req"
	"express_be/model/res"
	"express_be/repository/delivery"
	"express_be/repository/order"
)

type OrderUsecase interface {
	CreateOrder(ctx context.Context, req req.CreateOrderReq, id *string) *error.Err
	UpdateOrderStatus(ctx context.Context, id, status *string) *error.Err
	GetOrderList(ctx context.Context, user_id, role *string, page, limit *int) ([]res.OrderRes, *error.Err)
	GetOrderDetail(ctx context.Context, id *string) (*res.OrderRes, *error.Err)
	FindNearestDeliveryPerson(ctx context.Context, lat, lon *float64, status *string) (*res.BaseDeliveryPersonRes, *error.Err)
}

type orderUsecaseImpl struct {
	orderRepo    order.Repo
	deliveryRepo delivery.Repo
}

func NewOrderUsecase(orderRepo order.Repo, deliveryPersonRepo delivery.Repo) OrderUsecase {
	return &orderUsecaseImpl{orderRepo: orderRepo}
}
