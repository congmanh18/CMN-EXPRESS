package customer

import (
	"context"
	"express_be/entity"

	"fmt"
)

func (r *customerImpl) Create(ctx context.Context, customer *entity.Customer) error {
	if err := r.DB.Executor.WithContext(ctx).Create(customer).Error; err != nil {
		return fmt.Errorf("failed to create customer: %w", err)
	}
	return nil
}
