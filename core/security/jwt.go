package security

import (
	"fmt"
	"time"

	coreModel "express_be/core/model"

	"github.com/golang-jwt/jwt/v5"

	"github.com/spf13/viper"
)

type Token struct {
	AccessToken          *string        `json:"access_token"`
	RefreshToken         *string        `json:"refresh_token"`
	AccessTokenExpiresIn *time.Duration `json:"access_token_expires_in"`
	FreshTokenExpiresIn  *time.Duration `json:"fresh_token_expires_in"`
}

func GenToken(ID string, accessTokenDuration, refreshTokenDuration time.Duration) (*Token, error) {
	accessToken, err := GenAccessToken(ID, accessTokenDuration)
	if err != nil {
		return nil, fmt.Errorf("failed to generate access token: %w", err)
	}

	// Tạo RefreshToken
	refreshToken, err := GenRefreshToken(ID, refreshTokenDuration)
	if err != nil {
		return nil, fmt.Errorf("failed to generate refresh token: %w", err)
	}

	// Trả về cấu trúc Token
	return &Token{
		AccessToken:          accessToken,
		RefreshToken:         refreshToken,
		AccessTokenExpiresIn: &accessTokenDuration,
		FreshTokenExpiresIn:  &refreshTokenDuration,
	}, nil
}

func GenAccessToken(ID string, duration time.Duration) (*string, error) {
	claimsToken := &coreModel.JwtCustomClaims{
		ID: ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		},
	}

	secretKey := []byte(viper.Get("JWT_SECRET_KEY").(string))
	accessTokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsToken)

	accessTokenString, err := accessTokenObj.SignedString(secretKey)
	if err != nil {
		return nil, err
	}
	return &accessTokenString, nil
}

func GenRefreshToken(ID string, duration time.Duration) (*string, error) {
	var (
		claimsToken = &coreModel.JwtCustomClaims{
			ID: ID,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			},
		}
		secretKey       = []byte(viper.Get("JWT_SECRET_KEY").(string))
		refreshTokenObj = jwt.NewWithClaims(jwt.SigningMethodHS256, claimsToken)
	)

	var refreshTokenString, err = refreshTokenObj.SignedString(secretKey)
	if err != nil {
		return nil, err
	}
	return &refreshTokenString, nil
}

func ParseTokenFromString(tokenString string) (*coreModel.JwtCustomClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&coreModel.JwtCustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(viper.Get("JWT_SECRET_KEY").(string)), nil
		})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*coreModel.JwtCustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
