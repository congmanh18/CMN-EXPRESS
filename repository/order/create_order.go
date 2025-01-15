package order

import (
	"context"
	"express_be/repository/order/entity"
)

func (r *orderImpl) Create(ctx context.Context, order *entity.Order) error {
	return r.DB.Executor.WithContext(ctx).Debug().Create(&order).Error
}
