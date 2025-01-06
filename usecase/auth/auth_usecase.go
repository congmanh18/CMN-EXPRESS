package auth

import (
	"context"
	"express_be/core/security"
	"express_be/repository/accounting"
	"express_be/repository/admin"
	"express_be/repository/customer"
	customerEntity "express_be/repository/customer/entity"
	"express_be/repository/delivery"
	deliveryPersonEntity "express_be/repository/delivery/entity"
	"express_be/repository/token"
	"express_be/usecase"
)

type AuthUsecase interface {
	CreateAdmin(ctx context.Context, user *admin.Admin) *usecase.Error
	CreateAccounting(ctx context.Context, user *accounting.Accounting) *usecase.Error
	CreateCustomer(ctx context.Context, customer *customerEntity.Customer) *usecase.Error
	CreateDeliveryPerson(ctx context.Context, deliveryPerson *deliveryPersonEntity.DeliveryPerson) *usecase.Error
	LoginAdmin(ctx context.Context, phone, password *string) (*security.Token, *admin.Admin, *usecase.Error)
	LoginCustomer(ctx context.Context, phone, password *string) (*security.Token, *customerEntity.Customer, *usecase.Error)
	LoginDeliveryPerson(ctx context.Context, phone, password *string) (*security.Token, *deliveryPersonEntity.DeliveryPerson, *usecase.Error)
	LoginAccounting(ctx context.Context, phone, password *string) (*security.Token, *accounting.Accounting, *usecase.Error)
	ValidateRefreshToken(ctx context.Context, refreshToken *string) (*string, *usecase.Error)
}

type authUsecaseImpl struct {
	adminRepo          admin.Repo
	customerRepo       customer.Repo
	deliveryPersonRepo delivery.Repo
	accountingRepo     accounting.Repo
	tokenRepo          token.Repo
}

func NewAuthUsecase(
	adminRepo admin.Repo,
	customerRepo customer.Repo,
	deliveryPersonRepo delivery.Repo,
	accountingRepo accounting.Repo,
	tokenRepo token.Repo,
) AuthUsecase {
	return &authUsecaseImpl{
		adminRepo:          adminRepo,
		customerRepo:       customerRepo,
		deliveryPersonRepo: deliveryPersonRepo,
		tokenRepo:          tokenRepo,
		accountingRepo:     accountingRepo,
	}
}
