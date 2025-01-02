package auth

import (
	"context"
	"errors"
	customerEntity "express_be/repository/customer/entity"
	"fmt"

	"express_be/usecase"

	"github.com/lib/pq"
)

func (c *authUsecaseImpl) CreateCustomerUsecase(ctx context.Context, customer *customerEntity.Customer) *usecase.Error {
	err := c.customerRepo.CreateCustomer(ctx, customer)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) && pqErr.Code == "23505" { // SQLSTATE 23505: duplicate key
			return &usecase.Error{
				Code:    400,
				Message: fmt.Sprintf("Phone already registered: %s", *customer.Phone),
				Err:     errors.New("customer already exists"),
			}
		}

		return &usecase.Error{
			Code:    500,
			Message: "Failed to create customer - " + err.Error(),
			Err:     err,
		}
	}

	return nil
}
