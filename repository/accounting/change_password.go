package accounting

import (
	"context"
	"fmt"
)

// ChangePassword implements Repo.
func (a *accountingImpl) ChangePassword(ctx context.Context, phone *string, newPassword *string) error {
	result := a.DB.Executor.WithContext(ctx).
		Model(&Accounting{}).
		Where("phone =?", *phone).
		Update("password", newPassword)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no accounting found with: %s", *phone)
	}

	return nil
}
