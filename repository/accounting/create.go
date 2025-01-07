package accounting

import (
	"context"
	accountingEntity "express_be/repository/accounting/entity"

	"fmt"
)

func (r *accountingImpl) Create(ctx context.Context, accounting *accountingEntity.Accounting) error {
	if err := r.DB.Executor.WithContext(ctx).Create(accounting).Error; err != nil {
		return fmt.Errorf("failed to create accounting: %w", err)
	}
	return nil
}
