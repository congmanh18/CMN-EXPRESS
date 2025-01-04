package auth

import (
	"context"
	"errors"
	"express_be/core/security"
	"express_be/repository/customer"
	customerEntity "express_be/repository/customer/entity"
	"express_be/repository/delivery"
	deliveryPersonEntity "express_be/repository/delivery/entity"
	"express_be/repository/token"
	"express_be/usecase"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type AuthUsecase interface {
	CreateCustomerUsecase(ctx context.Context, customer *customerEntity.Customer) *usecase.Error
	CreateDeliveryPersonUsecase(ctx context.Context, deliveryPerson *deliveryPersonEntity.DeliveryPerson) *usecase.Error
	LoginCustomer(ctx context.Context, phone, password *string) (*security.Token, *customerEntity.Customer, *usecase.Error)
	LoginDeliveryPerson(ctx context.Context, phone, password *string) (*security.Token, *deliveryPersonEntity.DeliveryPerson, *usecase.Error)
	ValidateRefreshToken(ctx context.Context, refreshToken *string) (*string, *usecase.Error)
}

type authUsecaseImpl struct {
	customerRepo       customer.Repo
	deliveryPersonRepo delivery.Repo
	tokenRepo          token.Repo
}

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

func NewAuthUsecase(customerRepo customer.Repo, deliveryPersonRepo delivery.Repo, tokenRepo token.Repo) AuthUsecase {
	return &authUsecaseImpl{
		customerRepo:       customerRepo,
		deliveryPersonRepo: deliveryPersonRepo,
		tokenRepo:          tokenRepo,
	}
}
