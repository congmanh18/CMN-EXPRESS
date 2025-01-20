package delivery

import (
	"context"
	"express_be/core/db/postgresql"
	"express_be/entity"
)

type Repo interface {
	Create(ctx context.Context, deliveryPerson *entity.DeliveryPerson) error
	FetchID(ctx context.Context, id *string) (*string, error)
	FindByPhone(ctx context.Context, phone *string) (*entity.DeliveryPerson, error)
	FindByID(ctx context.Context, id *string) (*entity.DeliveryPerson, error)
	DeleteDeliveryPerson(ctx context.Context, id *string) error
	Update(ctx context.Context, id *string, deliveryPerson *entity.DeliveryPerson) error
	FetchAllDeliveryPersons(ctx context.Context, page, pageSize *int) ([]entity.DeliveryPerson, error)
	FetchPhone(ctx context.Context, phone *string) (*string, error)
}

type deliveryImpl struct {
	DB *postgresql.Database
}

func NewRepo(db *postgresql.Database) Repo {
	return &deliveryImpl{
		DB: db,
	}
}
