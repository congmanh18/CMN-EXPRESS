package delivery

import (
	"context"
	"express_be/repository/delivery"
	deliveryPersonEntity "express_be/repository/delivery/entity"
	"express_be/repository/user"
	userEntity "express_be/repository/user/entity"

	"express_be/usecase"
)

type DeliveryPersonUsecase interface {
	GetInfoDeliveryPerson(ctx context.Context, id *string) (*userEntity.DeliveryPersonDetails, *usecase.Error)
	UpdateInfoDeliveryPerson(ctx context.Context, userEntity *userEntity.User, deliveryPerson *deliveryPersonEntity.DeliveryPerson, id *string) *usecase.Error
	SoftDeleteDeliveryPerson(ctx context.Context, id *string) *usecase.Error
	GetAllDeliveryPersons(ctx context.Context, page *int, pageSize *int) ([]userEntity.DeliveryPersonDetails, *usecase.Error)
	GetPendingDeliveryPersons(ctx context.Context, page, pageSize *int) ([]userEntity.DeliveryPersonDetails, *usecase.Error)
}

type deliveryPersonUsecaseImpl struct {
	userRepo           user.Repo
	deliveryPersonRepo delivery.Repo
}

func NewDeliveryPersonUsecase(userRepo user.Repo, deliveryPersonRepo delivery.Repo) DeliveryPersonUsecase {
	return &deliveryPersonUsecaseImpl{
		userRepo:           userRepo,
		deliveryPersonRepo: deliveryPersonRepo,
	}
}
