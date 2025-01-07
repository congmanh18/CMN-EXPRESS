package auth

import (
	"context"
	"express_be/core/security"
	mapper "express_be/mapper/req"
	"express_be/usecase"

	"time"
)

func (a *authUsecaseImpl) LoginCustomer(ctx context.Context, phone, password *string) (*security.Token, *usecase.Error) {
	user, err := a.userRepo.FindByPhone(ctx, phone)
	if err != nil || user == nil {
		return nil, &usecase.Error{
			Code:    401,
			Message: "The phone number is not registered or invalid. " + err.Error(),
			Err:     err,
		}
	}

	customer, err := a.customerRepo.FindByPhone(ctx, phone)
	if err != nil || customer == nil {
		return nil, &usecase.Error{
			Code:    401,
			Message: "Customer information not found. Please check your details. " + err.Error(),
			Err:     err,
		}
	}
	if !security.VerifyPassword(*password, *user.Password) {
		return nil, &usecase.Error{
			Code:    401,
			Message: "Incorrect password. Please try again.",
			Err:     nil,
		}
	}

	accessTokenDuration := time.Hour * 8
	refreshTokenDuration := time.Hour * 24 * 14
	scToken, err := security.GenToken(*customer.ID, accessTokenDuration, refreshTokenDuration)
	if err != nil {
		return nil, &usecase.Error{
			Code:    500,
			Message: "Unable to generate access token. Please try again later. " + err.Error(),
			Err:     err,
		}
	}

	token := mapper.SecureTokenToTokenEntity(scToken, customer.ID, refreshTokenDuration)
	err = a.tokenRepo.SaveToken(ctx, token)
	if err != nil {
		return nil, &usecase.Error{
			Code:    500,
			Message: "Failed to save token. Please contact support if this issue persists. " + err.Error(),
			Err:     err,
		}
	}

	return scToken, nil
}
