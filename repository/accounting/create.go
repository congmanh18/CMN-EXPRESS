package accounting

import (
	"context"
	"express_be/entity"

	"fmt"
)

func (r *accountingImpl) Create(ctx context.Context, accounting *entity.Accounting) error {
	if err := r.DB.Executor.WithContext(ctx).Create(accounting).Error; err != nil {
		return fmt.Errorf("failed to create accounting: %w", err)
	}
	return nil
}
