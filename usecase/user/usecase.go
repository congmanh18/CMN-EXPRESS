package user

import (
	"context"
	error "express_be/core/err"
	"express_be/repository/customer"
	"express_be/repository/delivery"
	"express_be/repository/user"
	"express_be/repository/user/entity"
)

type UserUsecase interface {
	GetInfoUser(ctx context.Context, id *string) (*entity.CustomerDetails, *entity.DeliveryPersonDetails, *error.Err)
	GetUsers(ctx context.Context, status, role *string, page, pageSize *int) ([]entity.CustomerDetails, []entity.DeliveryPersonDetails, *error.Err)
	UpdateStatus(ctx context.Context, userID *string, status *string) *error.Err
	Search(ctx context.Context, role, phone, name, status *string, page, pageSize *int) ([]entity.CustomerDetails, []entity.DeliveryPersonDetails, *error.Err)
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
