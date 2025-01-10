package token

import (
	"context"
	"errors"
	"express_be/core/db/postgresql"
	"fmt"

	"gorm.io/gorm"
)

type Repo interface {
	ValidateToken(ctx context.Context, refreshToken *string) (*RefreshToken, error)
	SaveToken(ctx context.Context, refreshToken *RefreshToken) error
}

type tokenImpl struct {
	DB *postgresql.Database
}

func (t *tokenImpl) SaveToken(ctx context.Context, refreshToken *RefreshToken) error {
	return t.DB.Executor.WithContext(ctx).Save(refreshToken).Error
}

func (t *tokenImpl) ValidateToken(ctx context.Context, refreshToken *string) (*RefreshToken, error) {
	var result *RefreshToken
	err := t.DB.Executor.WithContext(ctx).
		Model(&RefreshToken{}).
		Where("token = ?", refreshToken).
		First(&result).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("refresh token not found or invalid")
		}
		return nil, fmt.Errorf("database error: %w", err)
	}
	return result, nil
}

func NewRepo(db *postgresql.Database) Repo {
	return &tokenImpl{DB: db}
}
