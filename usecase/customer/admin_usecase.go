package customer

import (
	"context"
	"express_be/repository/customer"
	"express_be/repository/customer/entity"
	"express_be/usecase"
)

type AdminUsecase interface {
	AdminGetPendingCustomers(ctx context.Context, page, pageSize *int) ([]entity.Customer, *usecase.Error)
	AdminGetAllCustomers(ctx context.Context, page, pageSize *int) ([]entity.Customer, *usecase.Error)
	AdminUpdateStatusCustomer(ctx context.Context, customerID *string, status *string) *usecase.Error
}

type adminUsecaseImpl struct {
	customerRepo customer.Repo
}

func NewAdminUsecase(repo customer.Repo) AdminUsecase {
	return &adminUsecaseImpl{
		customerRepo: repo,
	}
}
