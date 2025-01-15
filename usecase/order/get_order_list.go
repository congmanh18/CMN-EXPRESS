package order

import (
	"context"
	error "express_be/core/err"
	mapper "express_be/mapper/res"
	model "express_be/model/res"
)

// GetOrderList implements OrderUsecase.
func (o *orderUsecaseImpl) GetOrderList(ctx context.Context, user_id, role *string, page *int, limit *int) ([]model.OrderRes, *error.Err) {
	switch *role {
	case "admin":
		orders, err := o.orderRepo.FetchAllOrders(ctx, page, limit)
		if err != nil {
			return nil, error.ErrInternalServer
		}

		var orderResponses []model.OrderRes
		for _, order := range orders {
			res := mapper.OrderDetailToRes(&order)
			orderResponses = append(orderResponses, res)
		}

		return orderResponses, nil
	case "customer":
		orders, err := o.orderRepo.FetchOrderByCustomer(ctx, user_id, page, limit)
		if err != nil {
			return nil, error.ErrInternalServer
		}

		var orderResponses []model.OrderRes
		for _, order := range orders {
			res := mapper.OrderDetailToRes(&order)
			orderResponses = append(orderResponses, res)
		}

		return orderResponses, nil
	case "delivery_person":
		orders, err := o.orderRepo.FetchOrderByDeliveryPerson(ctx, user_id, page, limit)
		if err != nil {
			return nil, error.ErrInternalServer
		}

		var orderResponses []model.OrderRes
		for _, order := range orders {
			res := mapper.OrderDetailToRes(&order)
			orderResponses = append(orderResponses, res)
		}

		return orderResponses, nil
	}

	return nil, error.ErrInvalidRole
}
