package accounting

import "context"

// Create implements Repo.
func (a *accountingImpl) CreateAccounting(ctx context.Context, accounting *Accounting) error {
	return a.DB.Executor.WithContext(ctx).Debug().Create(&accounting).Error
}
