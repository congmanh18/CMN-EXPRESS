package auth

import (
	"context"
	accountingEntity "express_be/repository/accounting/entity"
	userEntity "express_be/repository/user/entity"
	"express_be/usecase"
)

func (c *authUsecaseImpl) CreateAccounting(ctx context.Context, user *userEntity.User, accounting *accountingEntity.Accounting) *usecase.Error {
	// Thêm user trước
	err := c.userRepo.Create(ctx, user)
	if err != nil {
		return &usecase.Error{
			Code:    500,
			Message: "Failed to create user. Please try again later. " + err.Error(),
			Err:     err,
		}
	}

	// Sau đó thêm accounting
	err = c.accountingRepo.Create(ctx, accounting)
	if err != nil {
		return &usecase.Error{
			Code:    500,
			Message: "Failed to create accounting. Please try again later. " + err.Error(),
			Err:     err,
		}
	}

	return nil
}
