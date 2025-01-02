package admin

import (
	"context"
	"express_be/core/security"
	"express_be/repository/admin"
	"express_be/usecase"
)

type AdminUsecase interface {
	// AddUser adds a new user to the system.
	CreateAdmin(ctx context.Context, user *admin.Admin) *usecase.Error
	LoginAdmin(ctx context.Context, phone, password *string) (*string, *usecase.Error)
}

type adminUsecaseImpl struct {
	repo admin.Repo
}

func NewAdminUsecase(repo admin.Repo) AdminUsecase {
	return &adminUsecaseImpl{
		repo: repo,
	}
}

func (a *adminUsecaseImpl) CreateAdmin(ctx context.Context, user *admin.Admin) *usecase.Error {
	err := a.repo.CreateAdmin(ctx, user)
	if err != nil {
		return &usecase.Error{
			Code:    500,
			Message: err.Error(),
			Err:     err,
		}
	}
	return nil
}

func (a *adminUsecaseImpl) LoginAdmin(ctx context.Context, phone, password *string) (*string, *usecase.Error) {
	admin, err := a.repo.FindByPhone(ctx, phone)
	if err != nil {
		return nil, &usecase.Error{
			Code:    401,
			Message: "Invalid phone or password",
			Err:     err,
		}
	}
	if !security.VerifyPassword(*password, *admin.Password) {
		return nil, &usecase.Error{
			Code:    401,
			Message: "Invalid phone or password",
			Err:     err,
		}
	}

	token := "password"
	// token, err := a.authService.GenerateToken(*admin.ID, "admin")
	// if err != nil {
	// 	return nil, &usecase.Error{
	// 		Code:    500,
	// 		Message: "Failed to generate token",
	// 		Err:     err,
	// 	}
	// }
	return &token, nil
}
