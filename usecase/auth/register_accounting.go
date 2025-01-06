package auth

import (
	"context"
	"express_be/repository/accounting"
	"express_be/usecase"
)

func (a *authUsecaseImpl) CreateAccounting(ctx context.Context, user *accounting.Accounting) *usecase.Error {
	err := a.accountingRepo.CreateAccounting(ctx, user)
	if err != nil {
		return &usecase.Error{
			Code:    500,
			Message: err.Error(),
			Err:     err,
		}
	}
	return nil
}
