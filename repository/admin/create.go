package admin

import "context"

// CreateAdmin implements Repo.
func (a *adminImpl) CreateAdmin(ctx context.Context, admin *Admin) error {
	return a.DB.Executor.WithContext(ctx).Debug().Create(&admin).Error
}
