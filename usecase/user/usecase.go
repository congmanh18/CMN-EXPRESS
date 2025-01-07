package user

import (
	"context"
	"express_be/repository/user"
	"express_be/usecase"
)

type UserUsecase interface {
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
