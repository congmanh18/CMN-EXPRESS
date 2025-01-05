package auth

import (
	"context"
	"express_be/core/security"
	mapper "express_be/mapper/req"
	deliveryPersonEntity "express_be/repository/delivery/entity"
	"express_be/usecase"
	"time"
)

func (u *authUsecaseImpl) LoginDeliveryPerson(ctx context.Context, phone, password *string) (*security.Token, *deliveryPersonEntity.DeliveryPerson, *usecase.Error) {
	deliveryPerson, err := u.deliveryPersonRepo.FindByPhone(ctx, phone)
	if deliveryPerson == nil {
		return nil, nil, &usecase.Error{
			Code:    401,
			Message: "Invalid phone or password",
		}
	}
	if err != nil {
		return nil, nil, &usecase.Error{
			Code:    404,
			Message: "Invalid phone number or password",
			Err:     err,
		}
	}

	if !security.VerifyPassword(*password, *deliveryPerson.PasswordHash) {
		return nil, nil, &usecase.Error{
			Code:    401,
			Message: "Invalid phone number or password",
			Err:     err,
		}
	}

	accessTokenDuration := time.Hour * 8
	refreshTokenDuration := time.Hour * 24 * 14
	scToken, err := security.GenToken(*deliveryPerson.ID, accessTokenDuration, refreshTokenDuration)
	if err != nil {
		return nil, nil, &usecase.Error{
			Code:    500,
			Message: "internal server error",
			Err:     err,
		}
	}

	token := mapper.SecureTokenToTokenEntity(scToken, deliveryPerson.ID, refreshTokenDuration)
	err = u.tokenRepo.SaveToken(ctx, token)
	if err != nil {
		return nil, nil, &usecase.Error{
			Code:    500,
			Message: "failed to save token" + err.Error(),
			Err:     err,
		}
	}

	return scToken, deliveryPerson, nil
}
