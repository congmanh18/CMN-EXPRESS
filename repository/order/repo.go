package order

import (
	"context"
	"express_be/core/db/postgresql"
	"express_be/repository/order/entity"
)

type Repo interface {
	Create(ctx context.Context, order *entity.Order) error
}

type orderImpl struct {
	DB *postgresql.Database
}

func NewRepo(db *postgresql.Database) Repo {
	return &orderImpl{
		DB: db,
	}
}

func (r *orderImpl) Create(ctx context.Context, order *entity.Order) error {
	return r.DB.Executor.WithContext(ctx).Debug().Create(&order).Error
}
