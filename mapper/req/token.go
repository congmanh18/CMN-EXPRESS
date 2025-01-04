package req

import (
	"express_be/core/pointer"
	"express_be/core/record"
	"express_be/core/security"
	"express_be/repository/token"
	"time"

	"github.com/google/uuid"
)

func SecureTokenToTokenEntity(scToken *security.Token, userID *string, refreshTokenDuration time.Duration) *token.RefreshToken {
	expires_at := time.Now().Add(refreshTokenDuration)

	return &token.RefreshToken{
		BaseEntity: record.BaseEntity{
			ID: pointer.String(uuid.New().String()),
		},
		UserID:    userID,
		Token:     scToken.RefreshToken,
		ExpiresAt: expires_at,
	}
}
