package admin

import (
	"context"

	"fmt"
)

// ChangePassword implements Repo.
func (a *adminImpl) ChangePassword(ctx context.Context, phone *string, newPassword *string) error {
	result := a.DB.Executor.WithContext(ctx).
		Model(&Admin{}).
		Where("phone =?", *phone).
		Update("password", newPassword)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no admin found with: %s", *phone)
	}

	return nil
}
