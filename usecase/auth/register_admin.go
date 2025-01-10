package auth

import (
	"context"
	error "express_be/core/err"
	adminEntity "express_be/repository/admin/entity"
	userEntity "express_be/repository/user/entity"
)

func (c *authUsecaseImpl) CreateAdmin(ctx context.Context, user *userEntity.User, admin *adminEntity.Admin) *error.Err {
	// Thêm user trước
	err := c.userRepo.Create(ctx, user)
	if err != nil {
		return error.ErrRegisterUser
	}

	// Sau đó thêm admin
	err = c.adminRepo.Create(ctx, admin)
	if err != nil {
		return error.ErrRegisterUser
	}

	return nil
}
