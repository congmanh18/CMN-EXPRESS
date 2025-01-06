package accounting

import "context"

// FindByPhone implements Repo.
func (a *accountingImpl) FindByPhone(ctx context.Context, phone *string) (*Accounting, error) {
	var result *Accounting
	query := a.DB.Executor.WithContext(ctx).
		Model(&Accounting{}).
		Where("phone =?", phone).
		First(&result)

	if query.Error != nil {
		return nil, query.Error
	}

	return result, nil
}
