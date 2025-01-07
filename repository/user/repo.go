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
	FetchPendingStatusCustomers(ctx context.Context, page, pageSize *int) ([]user.CustomerDetails, error)
	FetchPendingStatusDeliveryPerson(ctx context.Context, page, pageSize *int) ([]user.DeliveryPersonDetails, error)
	ChangePassword(ctx context.Context, phone *string, newPassword *string) error
	FindByID(ctx context.Context, id *string) (*user.User, error)
	FetchAllCustomer(ctx context.Context, page, pageSize *int) ([]user.CustomerDetails, error)
	FetchAllDeliveryPerson(ctx context.Context, page, pageSize *int) ([]user.DeliveryPersonDetails, error)
}

type userImpl struct {
	DB *postgresql.Database
}

func NewRepo(db *postgresql.Database) Repo {
	return &userImpl{
		DB: db,
	}
}
