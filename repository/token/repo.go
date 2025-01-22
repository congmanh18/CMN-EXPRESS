package token

import (
	"context"
	"errors"
	"express_be/core/db/postgresql"
	"express_be/entity"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Repo interface {
	CacheAccessToken(ctx context.Context, redisKey *string, token *string, duration int64) error
	ValidateAccessToken(ctx context.Context, redisKey *string, token *string) (bool, error)

	SaveRefreshToken(ctx context.Context, refreshToken *entity.RefreshToken) error
	ValidateToken(ctx context.Context, refreshToken *string) (*entity.RefreshToken, error)
}

type tokenImpl struct {
	DB    *postgresql.Database
	Cache *redis.Client
}

// ValidateAccessToken implements Repo.
func (t *tokenImpl) ValidateAccessToken(ctx context.Context, redisKey *string, token *string) (bool, error) {
	if redisKey == nil || token == nil {
		return false, errors.New("redis key or token is nil")
	}

	// Lấy token từ Redis
	cachedToken, err := t.Cache.Get(ctx, *redisKey).Result()
	if err == redis.Nil {
		return false, fmt.Errorf("token not found in cache")
	} else if err != nil {
		return false, fmt.Errorf("failed to get token from cache: %w", err)
	}

	// So sánh token
	if cachedToken != *token {
		return false, fmt.Errorf("invalid token")
	}

	return true, nil
}

// CacheAccessToken implements Repo.
func (t *tokenImpl) CacheAccessToken(ctx context.Context, redisKey *string, token *string, duration int64) error {
	if redisKey == nil || token == nil {
		return errors.New("redis key or token is nil")
	}

	// Lưu token vào Redis với thời gian hết hạn
	err := t.Cache.Set(ctx, *redisKey, *token, time.Duration(duration)*time.Second).Err()
	if err != nil {
		return fmt.Errorf("failed to cache access token: %w", err)
	}

	return nil
}

func (t *tokenImpl) SaveRefreshToken(ctx context.Context, refreshToken *entity.RefreshToken) error {
	return t.DB.Executor.WithContext(ctx).Save(refreshToken).Error
}

func (t *tokenImpl) ValidateToken(ctx context.Context, refreshToken *string) (*entity.RefreshToken, error) {
	var result *entity.RefreshToken
	err := t.DB.Executor.WithContext(ctx).
		Model(&entity.RefreshToken{}).
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
	return &tokenImpl{
		DB: db,
	}
}
