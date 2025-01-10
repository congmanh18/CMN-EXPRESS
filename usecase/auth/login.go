package auth

import (
	"context"
	error "express_be/core/err"
	"express_be/core/security"
	mapper "express_be/mapper/req"
	"time"
)

func (a *authUsecaseImpl) Login(ctx context.Context, phone, password *string) (*security.Token, *error.Err) {
	user, err := a.userRepo.FindByPhone(ctx, phone)
	if err != nil || user == nil {
		return nil, error.ErrNotFound
	}

	if !security.VerifyPassword(*password, *user.Password) {
		return nil, error.ErrIncorrectPassword
	}

	accessTokenDuration := time.Hour * 8
	refreshTokenDuration := time.Hour * 24 * 14
	// jwt := "secret_key"
	scToken, err := security.GenToken(*user.ID, accessTokenDuration, refreshTokenDuration)
	if err != nil {
		return nil, error.ErrGentokenFailure
	}

	token := mapper.SecureTokenToTokenEntity(scToken, user.ID, refreshTokenDuration)
	err = a.tokenRepo.SaveToken(ctx, token)
	if err != nil {
		return nil, error.ErrSaveTokenFailure
	}
	return scToken, nil
}
