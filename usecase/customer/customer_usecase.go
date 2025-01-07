package customer

import (
	"context"
	"express_be/repository/customer"
	customerEntity "express_be/repository/customer/entity"
	"express_be/repository/user"
	userEntity "express_be/repository/user/entity"
	"express_be/usecase"
)

type CustomerUsecase interface {
	GetInfoCustomer(ctx context.Context, id *string) (*userEntity.CustomerDetails, *usecase.Error)
	UpdateInfoCustomer(ctx context.Context, user *userEntity.User, customer *customerEntity.Customer, id *string) *usecase.Error
	SoftDeleteCustomer(ctx context.Context, id *string) *usecase.Error
	GetAllCustomers(ctx context.Context, page *int, pageSize *int) ([]userEntity.CustomerDetails, *usecase.Error)
	GetPendingCustomers(ctx context.Context, page, pageSize *int) ([]userEntity.CustomerDetails, *usecase.Error)
}

type customerUsecaseImpl struct {
	userRepo     user.Repo
	customerRepo customer.Repo
}

func NewCustomerUsecase(userRepo user.Repo, customerRepo customer.Repo) CustomerUsecase {
	return &customerUsecaseImpl{
		userRepo:     userRepo,
		customerRepo: customerRepo,
	}
}
