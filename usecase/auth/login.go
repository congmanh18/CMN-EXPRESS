package auth

import (
	"context"
	error "express_be/core/err"
	"express_be/core/jwt"
	"express_be/core/pointer"
	"express_be/core/record"
	"express_be/core/security"
	"express_be/entity"
	"fmt"

	"time"

	"github.com/google/uuid"
)

func (a *authUsecaseImpl) Login(ctx context.Context, phone, password *string) (*jwt.TokenPair, *error.Err) {
	user, err := a.userRepo.FindByPhone(ctx, phone)
	if err != nil || user == nil {
		return nil, error.ErrNotFound
	}

	if !security.VerifyPassword(*password, *user.Password) {
		return nil, error.ErrIncorrectPassword
	}

	accessTokenDuration := time.Hour * 8
	refreshTokenDuration := time.Hour * 24 * 14
	role := user.Role.String()
	refreshTokenExpiresAt := time.Now().Add(refreshTokenDuration)

	token, err := a.GenToken(*user.ID, role, accessTokenDuration, refreshTokenDuration)
	if err != nil {
		return nil, error.ErrGentokenFailure
	}

	// Lưu accessToken vào Redis
	redisKey := fmt.Sprintf("accessToken:%s", *user.ID)
	err = a.tokenRepo.CacheAccessToken(ctx, &redisKey, &token.AccessToken, int64(accessTokenDuration.Seconds()))
	if err != nil {
		return nil, error.ErrSaveTokenFailure
	}

	refreshToken := entity.RefreshToken{
		BaseEntity: record.BaseEntity{
			ID: pointer.String(uuid.New().String()),
		},
		UserID:    user.ID,
		Role:      &role,
		Token:     &token.RefreshToken,
		ExpiresAt: refreshTokenExpiresAt,
	}

	err = a.tokenRepo.SaveRefreshToken(ctx, &refreshToken)
	if err != nil {
		return nil, error.ErrSaveTokenFailure
	}
	return token, nil
}
