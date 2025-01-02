package auth

import (
	"context"
	"express_be/core/security"
	"express_be/usecase"
)

func (a *authUsecaseImpl) LoginCustomer(ctx context.Context, phone, password *string) (*string, *usecase.Error) {
	customer, err := a.customerRepo.FindByPhone(ctx, phone)
	if err != nil {
		return nil, &usecase.Error{
			Code:    404,
			Message: "customer not found",
			Err:     err,
		}
	}
	if !security.VerifyPassword(*password, *customer.PasswordHash) {
		return nil, &usecase.Error{
			Code:    401,
			Message: "invalid password",
			Err:     err,
		}
	}
	return customer.ID, nil
}
