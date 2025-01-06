package auth

import (
	"context"
	"express_be/repository/admin"
	"express_be/usecase"
)

func (a *authUsecaseImpl) CreateAdmin(ctx context.Context, user *admin.Admin) *usecase.Error {
	err := a.adminRepo.CreateAdmin(ctx, user)
	if err != nil {
		return &usecase.Error{
			Code:    500,
			Message: err.Error(),
			Err:     err,
		}
	}
	return nil
}
