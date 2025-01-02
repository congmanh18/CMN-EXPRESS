package auth

import (
	"context"
	"errors"
	customerEntity "express_be/repository/customer/entity"
	"express_be/usecase"
)

func (c *authUsecaseImpl) VerifyCustomerUsecase(ctx context.Context, uid *string) *usecase.Error {
	exists, err := c.customerRepo.FetchID(ctx, uid)
	if err != nil {
		return &usecase.Error{
			Code:    400,
			Message: "Customer not found",
			Err:     err,
		}
	}
	if exists != nil {
		return &usecase.Error{
			Code:    400,
			Message: "Customer already verified",
			Err:     errors.New("customer already verified"),
		}
	}

	// Convert the Status constant to a pointer
	status := customerEntity.Verified
	err = c.customerRepo.UpdateStatus(ctx, uid, &status)
	if err != nil {
		return &usecase.Error{
			Code:    500,
			Message: "Failed to verify customer",
			Err:     err,
		}
	}

	return nil
}
