package auth

import (
	"context"
	"errors"
	"express_be/repository/customer"
	customerEntity "express_be/repository/customer/entity"
	"fmt"

	"express_be/usecase"
)

type RegisterCustomerUseCase interface {
	CreateCustomerUsecase(ctx context.Context, customer *customerEntity.Customer) *usecase.Error
}

type RegisterCustomerUseCaseImpl struct {
	CustomerRepo customer.Repo
}

func (c *RegisterCustomerUseCaseImpl) CreateCustomerUsecase(ctx context.Context, customer *customerEntity.Customer) *usecase.Error {
	// Kiểm tra phone đã tồn tại chưa
	exists, err := c.CustomerRepo.FindByPhone(ctx, customer.Phone)
	if err != nil {
		return &usecase.Error{
			Code:    500,
			Message: "Failed to check phone existence",
			Err:     err,
		}
	}
	if exists != nil {
		return &usecase.Error{
			Code:    400,
			Message: fmt.Sprintf("Phone already registered as customer: %s", *exists),
			Err:     errors.New("customer already exists"),
		}
	}

	// Tạo khách hàng mới
	err = c.CustomerRepo.CreateCustomer(ctx, customer)
	if err != nil {
		return &usecase.Error{
			Code:    500,
			Message: "Failed to create customer",
			Err:     err,
		}
	}

	return nil
}

func NewRegisterCustomerUseCase(customerRepo customer.Repo) RegisterCustomerUseCase {
	return &RegisterCustomerUseCaseImpl{
		CustomerRepo: customerRepo,
	}
}
