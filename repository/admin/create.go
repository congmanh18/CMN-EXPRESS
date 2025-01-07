package admin

import (
	"context"
	adminEntity "express_be/repository/admin/entity"

	"fmt"
)

func (r *adminImpl) Create(ctx context.Context, admin *adminEntity.Admin) error {
	if err := r.DB.Executor.WithContext(ctx).Create(admin).Error; err != nil {
		return fmt.Errorf("failed to create admin: %w", err)
	}
	return nil
}
