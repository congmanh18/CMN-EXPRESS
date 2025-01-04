package customer

import (
	"context"
	"express_be/repository/customer"
	"express_be/repository/customer/entity"
	"express_be/usecase"
)

type CustomerUsecase interface {
	GetInfoCustomer(ctx context.Context, id *string) (*entity.Customer, *usecase.Error)
	UpdateInfoCustomer(ctx context.Context, customer *entity.Customer, id *string) *usecase.Error
	SoftDeleteCustomer(ctx context.Context, id *string) *usecase.Error
	ChangePassword(ctx context.Context, phone, password *string) *usecase.Error
}

type customerUsecaseImpl struct {
	customerRepo customer.Repo
}

func NewCustomerUsecase(repo customer.Repo) CustomerUsecase {
	return &customerUsecaseImpl{
		customerRepo: repo,
	}
}
