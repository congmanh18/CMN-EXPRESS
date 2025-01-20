package auth

import (
	"context"
	error "express_be/core/err"
	"express_be/entity"
)

func (c *authUsecaseImpl) CreateAccounting(ctx context.Context, user *entity.User, accounting *entity.Accounting) *error.Err {
	// Thêm user trước
	err := c.userRepo.Create(ctx, user)
	if err != nil {
		return error.ErrRegisterUser
	}
	// Sau đó thêm accounting
	err = c.accountingRepo.Create(ctx, accounting)
	if err != nil {
		return error.ErrRegisterUser
	}

	return nil
}
