package auth

import (
	"context"
	"errors"
	"express_be/usecase"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// ValidateRefreshToken implements AuthUsecase.
func (a *authUsecaseImpl) ValidateRefreshToken(ctx context.Context, refreshToken *string) (*string, *usecase.Error) {
	// 1. Truy vấn token từ repository
	token, err := a.tokenRepo.ValidateToken(ctx, refreshToken)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &usecase.Error{
				Code:    401,
				Message: "Invalid or expired refresh token",
				Err:     err,
			}
		}
		return nil, &usecase.Error{
			Code:    500,
			Message: "Internal server error",
			Err:     err,
		}
	}

	// 2. Kiểm tra thời gian hết hạn
	if time.Now().After(token.ExpiresAt) {
		return nil, &usecase.Error{
			Code:    401,
			Message: "Refresh token has expired",
			Err:     fmt.Errorf("token expired at %s", token.ExpiresAt),
		}
	}

	// 3. Trả về UserID
	return token.UserID, nil
}
