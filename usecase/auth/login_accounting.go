package auth

import (
	"context"
	"express_be/core/security"
	mapper "express_be/mapper/req"
	"express_be/repository/accounting"

	"express_be/usecase"
	"time"
)

func (a *authUsecaseImpl) LoginAccounting(ctx context.Context, phone, password *string) (*security.Token, *accounting.Accounting, *usecase.Error) {
	accounting, err := a.accountingRepo.FindByPhone(ctx, phone)
	if err != nil {
		return nil, nil, &usecase.Error{
			Code:    404,
			Message: "customer not found",
			Err:     err,
		}
	}
	if !security.VerifyPassword(*password, *accounting.Password) {
		return nil, nil, &usecase.Error{
			Code:    401,
			Message: "invalid password",
			Err:     nil,
		}
	}
	accessTokenDuration := time.Hour * 8
	refreshTokenDuration := time.Hour * 24 * 14
	scToken, err := security.GenToken(*accounting.ID, accessTokenDuration, refreshTokenDuration)
	if err != nil {
		return nil, nil, &usecase.Error{
			Code:    500,
			Message: "internal server error",
			Err:     err,
		}
	}
	token := mapper.SecureTokenToTokenEntity(scToken, accounting.ID, refreshTokenDuration)
	err = a.tokenRepo.SaveToken(ctx, token)
	if err != nil {
		return nil, nil, &usecase.Error{
			Code:    500,
			Message: "failed to save token" + err.Error(),
			Err:     err,
		}
	}

	return scToken, accounting, nil

}
