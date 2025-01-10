package auth

import (
	"context"
	error "express_be/core/err"
	"express_be/core/security"
)

// ChangePasswordAccounting implements AuthUsecase.
func (c *authUsecaseImpl) ChangePassword(ctx context.Context, phone *string, password *string) *error.Err {
	hashedPassword, err := security.HashPassword(*password)
	if err != nil {
		return error.ErrInternalServer
	}
	err = c.userRepo.ChangePassword(ctx, phone, &hashedPassword)
	if err != nil {
		return error.ErrChangePassword
	}

	return nil
}
