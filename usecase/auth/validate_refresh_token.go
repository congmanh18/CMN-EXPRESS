package auth

import (
	"context"
	"errors"
	error "express_be/core/err"
	"time"

	"gorm.io/gorm"
)

// ValidateRefreshToken implements AuthUsecase.
func (a *authUsecaseImpl) ValidateRefreshToken(ctx context.Context, refreshToken *string) (*string, *error.Err) {
	// 1. Truy vấn token từ repository
	token, err := a.tokenRepo.ValidateToken(ctx, refreshToken)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, error.ErrInvalidToken
		}
		return nil, error.ErrInternalServer
	}
	// 2. Kiểm tra thời gian hết hạn
	if time.Now().After(token.ExpiresAt) {
		return nil, error.ErrTokenExpired
	}
	// 3. Trả về UserID
	return token.UserID, nil
}
