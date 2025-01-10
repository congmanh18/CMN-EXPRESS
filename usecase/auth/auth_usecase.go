package auth

import (
	"context"
	"express_be/core/err"
	"express_be/core/security"
	"express_be/repository/accounting"
	accountingEntity "express_be/repository/accounting/entity"
	"express_be/repository/admin"
	adminEntity "express_be/repository/admin/entity"
	"express_be/repository/customer"
	customerEntity "express_be/repository/customer/entity"
	"express_be/repository/delivery"
	deliveryPersonEntity "express_be/repository/delivery/entity"
	"express_be/repository/token"
	"express_be/repository/user"
	userEntity "express_be/repository/user/entity"
)

type AuthUsecase interface {
	// Create accounts usecases
	CreateAdmin(ctx context.Context, user *userEntity.User, admin *adminEntity.Admin) *err.Err
	CreateAccounting(ctx context.Context, user *userEntity.User, accounting *accountingEntity.Accounting) *err.Err
	CreateCustomer(ctx context.Context, user *userEntity.User, customer *customerEntity.Customer) *err.Err
	CreateDeliveryPerson(ctx context.Context, user *userEntity.User, deliveryPerson *deliveryPersonEntity.DeliveryPerson) *err.Err
	// Login useaces
	Login(ctx context.Context, phone, password *string) (*security.Token, *err.Err)
	// ChangePassword useacses
	ChangePassword(ctx context.Context, phone *string, password *string) *err.Err
	// Validate Token
	ValidateRefreshToken(ctx context.Context, refreshToken *string) (*string, *err.Err)
}

type authUsecaseImpl struct {
	userRepo           user.Repo
	adminRepo          admin.Repo
	customerRepo       customer.Repo
	deliveryPersonRepo delivery.Repo
	accountingRepo     accounting.Repo
	tokenRepo          token.Repo
}

func NewAuthUsecase(
	userRepo user.Repo,
	adminRepo admin.Repo,
	customerRepo customer.Repo,
	deliveryPersonRepo delivery.Repo,
	accountingRepo accounting.Repo,
	tokenRepo token.Repo,
) AuthUsecase {
	return &authUsecaseImpl{
		userRepo:           userRepo,
		adminRepo:          adminRepo,
		customerRepo:       customerRepo,
		deliveryPersonRepo: deliveryPersonRepo,
		tokenRepo:          tokenRepo,
		accountingRepo:     accountingRepo,
	}
}
