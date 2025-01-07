package auth

import (
	"context"
	"express_be/core/security"
	mapper "express_be/mapper/req"
	"express_be/usecase"
	"time"
)

func (u *authUsecaseImpl) LoginDeliveryPerson(ctx context.Context, phone, password *string) (*security.Token, *usecase.Error) {
	user, err := u.userRepo.FindByPhone(ctx, phone)
	if err != nil || user == nil {
		return nil, &usecase.Error{
			Code:    401,
			Message: "The phone number is not registered or invalid. " + err.Error(),
			Err:     err,
		}
	}

	deliveryPerson, err := u.deliveryPersonRepo.FindByPhone(ctx, phone)
	if err != nil || deliveryPerson == nil {
		return nil, &usecase.Error{
			Code:    401,
			Message: "Invalid phone or password: " + err.Error(),
			Err:     err,
		}
	}

	if !security.VerifyPassword(*password, *user.Password) {
		return nil, &usecase.Error{
			Code:    401,
			Message: "Incorrect password. Please try again.",
			Err:     err,
		}
	}

	accessTokenDuration := time.Hour * 8
	refreshTokenDuration := time.Hour * 24 * 14
	scToken, err := security.GenToken(*deliveryPerson.ID, accessTokenDuration, refreshTokenDuration)
	if err != nil {
		return nil, &usecase.Error{
			Code:    500,
			Message: "Unable to generate access token. Please try again later. " + err.Error(),
			Err:     err,
		}
	}

	token := mapper.SecureTokenToTokenEntity(scToken, deliveryPerson.ID, refreshTokenDuration)
	err = u.tokenRepo.SaveToken(ctx, token)
	if err != nil {
		return nil, &usecase.Error{
			Code:    500,
			Message: "Failed to save token. Please contact support if this issue persists. " + err.Error(),
			Err:     err,
		}
	}

	return scToken, nil
}
