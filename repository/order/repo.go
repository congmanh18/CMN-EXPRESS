package order

import (
	"context"
	"express_be/core/db/postgresql"
	"express_be/repository/order/entity"
)

type Repo interface {
	Create(ctx context.Context, order *entity.Order) error
	UpdateOrderStatus(ctx context.Context, id *string, status *entity.Status) error
}

type orderImpl struct {
	DB *postgresql.Database
}

func NewRepo(db *postgresql.Database) Repo {
	return &orderImpl{
		DB: db,
	}
}
