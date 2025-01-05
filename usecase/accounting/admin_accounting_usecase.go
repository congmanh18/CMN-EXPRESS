package accounting

import (
	"context"
	"express_be/repository/accounting"
	"express_be/usecase"
)

type AdminUsecase interface {
	CreateAccounting(ctx context.Context, user *accounting.Accounting) *usecase.Error
}

type adminUsecaseImpl struct {
	repo accounting.Repo
}

func NewAdminUsecase(repo accounting.Repo) AdminUsecase {
	return &adminUsecaseImpl{repo: repo}
}

func (a *adminUsecaseImpl) CreateAccounting(ctx context.Context, user *accounting.Accounting) *usecase.Error {
	err := a.repo.CreateAccounting(ctx, user)
	if err != nil {
		return &usecase.Error{
			Code:    500,
			Message: err.Error(),
			Err:     err,
		}
	}
	return nil
}
