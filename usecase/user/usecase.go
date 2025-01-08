package user

import (
	"context"
	"express_be/repository/customer"
	"express_be/repository/delivery"
	"express_be/repository/user"
	"express_be/repository/user/entity"
	"express_be/usecase"
)

type UserUsecase interface {
	GetInfoUser(ctx context.Context, id *string) (*entity.CustomerDetails, *entity.DeliveryPersonDetails, *usecase.Error)
	GetUsers(ctx context.Context, status, role *string, page, pageSize *int) ([]entity.CustomerDetails, []entity.DeliveryPersonDetails, *usecase.Error)
	UpdateStatus(ctx context.Context, userID *string, status *string) *usecase.Error
}

type userUsecaseImpl struct {
	userRepo           user.Repo
	customerRepo       customer.Repo
	deliveryPersonRepo delivery.Repo
}

func NewUserUsecase(repo user.Repo, customerRepo customer.Repo, deliveryPersonRepo delivery.Repo) UserUsecase {
	return &userUsecaseImpl{
		userRepo:           repo,
		customerRepo:       customerRepo,
		deliveryPersonRepo: deliveryPersonRepo,
	}
}
