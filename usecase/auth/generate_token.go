package auth

import (
	core "express_be/core/jwt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func (a *authUsecaseImpl) GenAccessToken(userID, role string, duration time.Duration) (string, error) {
	claims := &core.JwtCustomClaims{
		ID:   userID,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(a.jwtSecret))
}

func (a *authUsecaseImpl) GenRefreshToken(userID string, duration time.Duration) (string, error) {
	claims := &core.JwtCustomClaims{
		ID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(a.jwtSecret))
}

func (a *authUsecaseImpl) GenToken(userID, role string, accessTokenDuration, refreshTokenDuration time.Duration) (*core.TokenPair, error) {
	accessToken, err := a.GenAccessToken(userID, role, accessTokenDuration)
	if err != nil {
		return nil, err
	}

	refreshToken, err := a.GenRefreshToken(userID, refreshTokenDuration)
	if err != nil {
		return nil, err
	}

	return &core.TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
