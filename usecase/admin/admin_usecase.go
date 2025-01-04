package admin

import (
	"context"
	"express_be/repository/admin"

	"express_be/usecase"
)

type AdminUsecase interface {
	CreateAdmin(ctx context.Context, user *admin.Admin) *usecase.Error
}

type adminUsecaseImpl struct {
	repo admin.Repo
}

func NewAdminUsecase(repo admin.Repo) AdminUsecase {
	return &adminUsecaseImpl{
		repo: repo,
	}
}

func (a *adminUsecaseImpl) CreateAdmin(ctx context.Context, user *admin.Admin) *usecase.Error {
	err := a.repo.CreateAdmin(ctx, user)
	if err != nil {
		return &usecase.Error{
			Code:    500,
			Message: err.Error(),
			Err:     err,
		}
	}
	return nil
}
