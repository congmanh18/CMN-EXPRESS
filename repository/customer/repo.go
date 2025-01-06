package customer

import (
	"context"
	"express_be/core/db/postgresql"
	"express_be/repository/customer/entity"
)

type Repo interface {
	CreateCustomer(ctx context.Context, customer *entity.Customer) error                                               // OKEY Dùng để đăng ký tài khoản
	FetchID(ctx context.Context, id *string) (*string, error)                                                          // Dùng để verify tài khoản 1
	FindByPhone(ctx context.Context, phone *string) (*entity.Customer, error)                                          // Dùng để Login
	FindByID(ctx context.Context, id *string) (*entity.Customer, error)                                                // OKEY Admin xem thông tin chi tiết tài khoản
	UpdateStatus(ctx context.Context, id *string, approval_status *entity.ApprovalStatus, status *entity.Status) error // Dùng để verify tài khoản 2
	DeleteCustomer(ctx context.Context, id *string) error                                                              // OKEY Dùng để xóa tài khoản
	UpdateCustomer(ctx context.Context, id *string, customer *entity.Customer) error                                   // OKEY Dùng để cập nhật thông tin tài khoản
	FetchAllCustomer(ctx context.Context, page, pageSize *int) ([]entity.Customer, error)                              // Dùng để lấy danh sách customers (tất cả)
	FetchPendingStatusCustomers(ctx context.Context, page, pageSize *int) ([]entity.Customer, error)                   // OKEY Lấy danh sách tài khoản chưa xác minh
	FetchPhoneWithCache(ctx context.Context, phone *string) (bool, error)
	ChangePassword(ctx context.Context, phone, newPassword *string) error
}

type customerImpl struct {
	DB *postgresql.Database
}

func NewRepo(db *postgresql.Database) Repo {
	return &customerImpl{
		DB: db,
	}
}
