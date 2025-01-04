package delivery

import (
	"context"
	"express_be/repository/delivery"
	"express_be/repository/delivery/entity"
	"express_be/usecase"
)

type DeliveryPersonUsecase interface {
	GetInfoDeliveryPerson(ctx context.Context, id *string) (*entity.DeliveryPerson, *usecase.Error)
	UpdateInfoDeliveryPerson(ctx context.Context, deliveryPerson *entity.DeliveryPerson, id *string) *usecase.Error
	SoftDeleteDeliveryPerson(ctx context.Context, id *string) *usecase.Error
	ChangePassword(ctx context.Context, phone, password *string) *usecase.Error
}

type deliveryPersonUsecaseImpl struct {
	deliveryPersonRepo delivery.Repo
}

func NewDeliveryPersonUsecase(repo delivery.Repo) DeliveryPersonUsecase {
	return &deliveryPersonUsecaseImpl{
		deliveryPersonRepo: repo,
	}
}
