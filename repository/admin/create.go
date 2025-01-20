package admin

import (
	"context"
	"express_be/entity"

	"fmt"
)

func (r *adminImpl) Create(ctx context.Context, admin *entity.Admin) error {
	if err := r.DB.Executor.WithContext(ctx).Create(admin).Error; err != nil {
		return fmt.Errorf("failed to create admin: %w", err)
	}
	return nil
}
