package user

import (
	"context"
	"express_be/core/db/postgresql"
	user "express_be/repository/user/entity"
)

type Repo interface {
	Create(ctx context.Context, user *user.User) error
	FindByPhone(ctx context.Context, phone *string) (*user.User, error)
	Update(ctx context.Context, id *string, user *user.User) error
	UpdateStatus(ctx context.Context, id *string, approval_status *user.ApprovalStatus, status *user.Status) error // Dùng để verify tài khoản 2
	FetchCustomerUsers(ctx context.Context, status *string, page, pageSize *int) ([]user.CustomerDetails, error)
	FetchDeliveryPersonUsers(ctx context.Context, status *string, page, pageSize *int) ([]user.DeliveryPersonDetails, error)
	ChangePassword(ctx context.Context, phone *string, newPassword *string) error
	FindByID(ctx context.Context, id *string) (*user.User, error)
}

type userImpl struct {
	DB *postgresql.Database
}

func NewRepo(db *postgresql.Database) Repo {
	return &userImpl{
		DB: db,
	}
}
