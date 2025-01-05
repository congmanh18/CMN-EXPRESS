package auth

import (
	"context"
	"express_be/core/security"
	mapper "express_be/mapper/req"
	customerEntity "express_be/repository/customer/entity"
	"express_be/usecase"

	"time"
)

func (a *authUsecaseImpl) LoginCustomer(ctx context.Context, phone, password *string) (*security.Token, *customerEntity.Customer, *usecase.Error) {
	customer, err := a.customerRepo.FindByPhone(ctx, phone)
	if customer == nil {
		return nil, nil, &usecase.Error{
			Code:    401,
			Message: "Invalid phone or password",
		}
	}
	if err != nil {
		return nil, nil, &usecase.Error{
			Code:    404,
			Message: "customer not found",
			Err:     err,
		}
	}
	if !security.VerifyPassword(*password, *customer.PasswordHash) {
		return nil, nil, &usecase.Error{
			Code:    401,
			Message: "invalid password",
			Err:     nil,
		}
	}

	accessTokenDuration := time.Hour * 8
	refreshTokenDuration := time.Hour * 24 * 14
	scToken, err := security.GenToken(*customer.ID, accessTokenDuration, refreshTokenDuration)
	if err != nil {
		return nil, nil, &usecase.Error{
			Code:    500,
			Message: "internal server error",
			Err:     err,
		}
	}

	token := mapper.SecureTokenToTokenEntity(scToken, customer.ID, refreshTokenDuration)
	err = a.tokenRepo.SaveToken(ctx, token)
	if err != nil {
		return nil, nil, &usecase.Error{
			Code:    500,
			Message: "failed to save token" + err.Error(),
			Err:     err,
		}
	}

	return scToken, customer, nil
}
