package user

import (
	"context"
	error "express_be/core/err"
)

func (c *userUsecaseImpl) SoftDeleteCustomer(ctx context.Context, id *string) *error.Err {
	err := c.customerRepo.DeleteCustomer(ctx, id)
	if err != nil {
		return error.ErrInternalServer
	}

	return nil
}
