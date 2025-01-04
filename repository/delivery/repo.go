package delivery

import (
	"context"
	"express_be/core/db/postgresql"
	"express_be/repository/delivery/entity"
)

type Repo interface {
	CreateDeliveryPerson(ctx context.Context, delivery *entity.DeliveryPerson) error
	FetchID(ctx context.Context, id *string) (*string, error)
	FindByPhone(ctx context.Context, phone *string) (*entity.DeliveryPerson, error)
	FindByID(ctx context.Context, id *string) (*entity.DeliveryPerson, error)
	UpdateStatus(ctx context.Context, id *string, status *entity.Status) error
	DeleteDeliveryPerson(ctx context.Context, id *string) error
	UpdateDeliveryPerson(ctx context.Context, id *string, deliveryPerson *entity.DeliveryPerson) error
	FetchAllDeliveryPersons(ctx context.Context, page, pageSize *int) ([]entity.DeliveryPerson, error)
	FetchPendingStatusDeliveryPerson(ctx context.Context, page, pageSize *int) ([]entity.DeliveryPerson, error)
	FetchPhone(ctx context.Context, phone *string) (*string, error)
	ChangePassword(ctx context.Context, phone, newPassword *string) error
}

type deliveryImpl struct {
	DB *postgresql.Database
}

func NewRepo(db *postgresql.Database) Repo {
	return &deliveryImpl{
		DB: db,
	}
}
