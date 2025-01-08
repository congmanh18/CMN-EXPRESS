package user

import (
	"context"
	"express_be/repository/user"
	"express_be/repository/user/entity"
	"express_be/usecase"
)

type UserUsecase interface {
	GetUsers(ctx context.Context, status, role *string, page, pageSize *int) ([]entity.CustomerDetails, []entity.DeliveryPersonDetails, *usecase.Error)
	UpdateStatus(ctx context.Context, userID *string, status *string) *usecase.Error
}

type userUsecaseImpl struct {
	repo user.Repo
}

func NewUserUsecase(repo user.Repo) UserUsecase {
	return &userUsecaseImpl{
		repo: repo,
	}
}
