package customer

import (
	"context"
	customerEntity "express_be/repository/customer/entity"

	"fmt"
)

func (r *customerImpl) Create(ctx context.Context, customer *customerEntity.Customer) error {
	if err := r.DB.Executor.WithContext(ctx).Create(customer).Error; err != nil {
		return fmt.Errorf("failed to create customer: %w", err)
	}
	return nil
}
