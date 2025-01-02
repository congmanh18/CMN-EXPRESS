package delivery

import (
	"context"
	"express_be/repository/delivery"
	"express_be/repository/delivery/entity"
	"express_be/usecase"
)

type AdminUsecase interface {
	AdminGetPendingDeliveryPersons(ctx context.Context, page, pageSize *int) ([]entity.DeliveryPerson, *usecase.Error)
}

type adminUseCaseImpl struct {
	deliveryPersonRepo delivery.Repo
}

func NewAdminUsecase(repo delivery.Repo) AdminUsecase {
	return &adminUseCaseImpl{
		deliveryPersonRepo: repo,
	}
}
