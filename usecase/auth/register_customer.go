package auth

import (
	"context"
	customerEntity "express_be/repository/customer/entity"
	userEntity "express_be/repository/user/entity"

	"express_be/usecase"
)

// func (c *authUsecaseImpl) CreateCustomer(ctx context.Context, user *user.User, customer *customerEntity.Customer) *usecase.Error {
// 	// Gọi hàm từ tầng repository để bắt đầu giao dịch
// 	err := c.userRepo.CreateAndLinkCustomer(ctx, user, customer)
// 	if err != nil {
// 		return &usecase.Error{
// 			Code:    500,
// 			Message: "Failed to create customer. Please try again later. " + err.Error(),
// 			Err:     err,
// 		}
// 	}
// 	return nil
// }

func (c *authUsecaseImpl) CreateCustomer(ctx context.Context, user *userEntity.User, customer *customerEntity.Customer) *usecase.Error {
	// Thêm user trước
	err := c.userRepo.Create(ctx, user)
	if err != nil {
		return &usecase.Error{
			Code:    500,
			Message: "Failed to create user. Please try again later. " + err.Error(),
			Err:     err,
		}
	}

	// Sau đó thêm customer
	err = c.customerRepo.Create(ctx, customer)
	if err != nil {
		return &usecase.Error{
			Code:    500,
			Message: "Failed to create customer. Please try again later. " + err.Error(),
			Err:     err,
		}
	}

	return nil
}
