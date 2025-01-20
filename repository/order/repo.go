package order

import (
	"context"
	"express_be/core/db/postgresql"
	"express_be/entity"
)

type Repo interface {
	Create(ctx context.Context, order *entity.Order) error
	UpdateOrderStatus(ctx context.Context, id *string, status *entity.Status) error
	FetchOrderByCustomer(ctx context.Context, id *string, page, pageSize *int) ([]entity.OrderDetail, error)
	FetchOrderByDeliveryPerson(ctx context.Context, id *string, page, pageSize *int) ([]entity.OrderDetail, error)
	FetchAllOrders(ctx context.Context, page, pageSize *int) ([]entity.OrderDetail, error)
	FindByID(ctx context.Context, id *string) (*entity.OrderDetail, error)
}

type orderImpl struct {
	DB *postgresql.Database
}

func NewRepo(db *postgresql.Database) Repo {
	return &orderImpl{
		DB: db,
	}
}
