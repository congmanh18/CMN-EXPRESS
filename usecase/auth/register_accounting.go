package auth

import (
	"context"
	error "express_be/core/err"
	accountingEntity "express_be/repository/accounting/entity"
	userEntity "express_be/repository/user/entity"
)

func (c *authUsecaseImpl) CreateAccounting(ctx context.Context, user *userEntity.User, accounting *accountingEntity.Accounting) *error.Err {
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
