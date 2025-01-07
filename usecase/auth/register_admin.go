package auth

import (
	"context"
	adminEntity "express_be/repository/admin/entity"
	userEntity "express_be/repository/user/entity"

	"express_be/usecase"
)

func (c *authUsecaseImpl) CreateAdmin(ctx context.Context, user *userEntity.User, admin *adminEntity.Admin) *usecase.Error {
	// Thêm user trước
	err := c.userRepo.Create(ctx, user)
	if err != nil {
		return &usecase.Error{
			Code:    500,
			Message: "Failed to create user. Please try again later. " + err.Error(),
			Err:     err,
		}
	}

	// Sau đó thêm admin
	err = c.adminRepo.Create(ctx, admin)
	if err != nil {
		return &usecase.Error{
			Code:    500,
			Message: "Failed to create admin. Please try again later. " + err.Error(),
			Err:     err,
		}
	}

	return nil
}
