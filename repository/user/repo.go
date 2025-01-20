package user

import (
	"context"
	"express_be/core/db/postgresql"
	"express_be/entity"
)

type Repo interface {
	Create(ctx context.Context, user *entity.User) error
	FindByPhone(ctx context.Context, phone *string) (*entity.User, error)
	Update(ctx context.Context, id *string, user *entity.User) error
	UpdateStatus(ctx context.Context, id *string, approval_status *entity.ApprovalStatus, status *entity.AccountStatus) error // Dùng để verify tài khoản 2
	FetchCustomerUsers(ctx context.Context, status *string, page, pageSize *int) ([]entity.CustomerDetails, error)
	FetchDeliveryPersonUsers(ctx context.Context, status *string, page, pageSize *int) ([]entity.DeliveryPersonDetails, error)
	ChangePassword(ctx context.Context, phone *string, newPassword *string) error
	FindByID(ctx context.Context, id *string) (*entity.User, error)
	SearchCustomers(ctx context.Context, phone, name, status *string, page, pageSize *int) ([]entity.CustomerDetails, int64, error)
	SearchDeliveryPersons(ctx context.Context, phone, name, status *string, page, pageSize *int) ([]entity.DeliveryPersonDetails, int64, error)
}

type userImpl struct {
	DB *postgresql.Database
}

func NewRepo(db *postgresql.Database) Repo {
	return &userImpl{
		DB: db,
	}
}
