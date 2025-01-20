package auth

import (
	"context"
	error "express_be/core/err"
	"express_be/entity"
)

func (c *authUsecaseImpl) CreateAdmin(ctx context.Context, user *entity.User, admin *entity.Admin) *error.Err {
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
