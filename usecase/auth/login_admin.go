package auth

import (
	"context"
	"express_be/core/security"
	mapper "express_be/mapper/req"
	"express_be/usecase"
	"time"
)

func (a *authUsecaseImpl) LoginAdmin(ctx context.Context, phone, password *string) (*security.Token, *usecase.Error) {
	admin, err := a.adminRepo.FindByPhone(ctx, phone)
	if err != nil {
		return nil, &usecase.Error{
			Code:    401,
			Message: "Invalid phone or password",
			Err:     err,
		}
	}
	if !security.VerifyPassword(*password, *admin.Password) {
		return nil, &usecase.Error{
			Code:    401,
			Message: "Invalid phone or password",
			Err:     err,
		}
	}

	accessTokenDuration := time.Hour * 8
	refreshTokenDuration := time.Hour * 24 * 14
	// jwt := "secret_key"
	scToken, err := security.GenToken(*admin.ID, accessTokenDuration, refreshTokenDuration)
	if err != nil {
		return nil, &usecase.Error{
			Code:    500,
			Message: "internal server error",
			Err:     err,
		}
	}

	token := mapper.SecureTokenToTokenEntity(scToken, admin.ID, refreshTokenDuration)
	err = a.tokenRepo.SaveToken(ctx, token)
	if err != nil {
		return nil, &usecase.Error{
			Code:    500,
			Message: "failed to save token" + err.Error(),
			Err:     err,
		}
	}
	return scToken, nil
}
