package admin

import "context"

// FindByPhone implements Repo.
func (a *adminImpl) FindByPhone(ctx context.Context, phone *string) (*Admin, error) {
	var result Admin
	query := a.DB.Executor.WithContext(ctx).
		Model(&Admin{}).
		Where("phone =?", phone).
		First(&result)

	if query.Error != nil {
		return nil, query.Error
	}

	return &result, nil
}
