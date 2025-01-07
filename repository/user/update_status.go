package user

import (
	"context"
	"express_be/repository/user/entity"
	"fmt"
)

func (c *userImpl) UpdateStatus(ctx context.Context, id *string, approval_status *entity.ApprovalStatus, status *entity.Status) error {
	result := c.DB.Executor.WithContext(ctx).
		Model(&entity.User{}).
		Where("id = ?", *id).
		Updates(map[string]interface{}{
			"approval_status": approval_status,
			"status":          status,
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no customer found with id: %s", *id)
	}

	return nil // Thành công
}
