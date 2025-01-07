package auth

import (
	"context"
	"express_be/repository/admin"
	user "express_be/repository/user/entity"

	"express_be/usecase"
)

func (a *authUsecaseImpl) CreateAdmin(ctx context.Context, user *user.User, admin *admin.Admin) *usecase.Error {
	err := a.userRepo.Create(ctx, user)
	if err != nil {
		return &usecase.Error{
			Code:    500,
			Message: "Failed to create user account. Please try again later." + err.Error(),
			Err:     err,
		}
	}

	err = a.adminRepo.CreateAdmin(ctx, admin)
	if err != nil {
		return &usecase.Error{
			Code:    500,
			Message: "Failed to create accounting record. Please try again later. " + err.Error(),
			Err:     err,
		}
	}
	return nil
}
