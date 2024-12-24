package delivery

import (
	"context"
	"express_be/core/db/postgresql"
	"express_be/repository/delivery/entity"
)

type Repo interface {
	CreateDeliveryPerson(ctx context.Context, delivery *entity.DeliveryPerson) error
	FindByPhone(ctx context.Context, phone *string) (*string, error) // Chưa phân trang
}

type deliveryImpl struct {
	DB *postgresql.Database
}

func NewRepo(db *postgresql.Database) Repo {
	return &deliveryImpl{
		DB: db,
	}
}
