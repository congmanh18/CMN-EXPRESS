package admin

import (
	"context"
	"express_be/core/security"
	mapper "express_be/mapper/req"
	"express_be/repository/admin"

	"express_be/repository/token"
	"express_be/usecase"
	"time"
)

type AdminUsecase interface {
	// AddUser adds a new user to the system.
	CreateAdmin(ctx context.Context, user *admin.Admin) *usecase.Error
	LoginAdmin(ctx context.Context, phone, password *string) (*security.Token, *usecase.Error)
}

type adminUsecaseImpl struct {
	repo      admin.Repo
	tokenRepo token.Repo
}

func NewAdminUsecase(repo admin.Repo, tokenRepo token.Repo) AdminUsecase {
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

func (a *adminUsecaseImpl) LoginAdmin(ctx context.Context, phone, password *string) (*security.Token, *usecase.Error) {
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

	accessTokenDuration := time.Hour * 8
	refreshTokenDuration := time.Hour * 24 * 14
	// jwt := "secret_key"
	scToken, err := security.GenToken(*admin.ID, accessTokenDuration, refreshTokenDuration)
	if err != nil {
		return nil, &usecase.Error{
			Code:    500,
			Message: "internal server error",
			Err:     err,
		}
	}

	token := mapper.SecureTokenToTokenEntity(scToken, admin.ID, refreshTokenDuration)
	err = a.tokenRepo.SaveToken(ctx, token)
	if err != nil {
		return nil, &usecase.Error{
			Code:    500,
			Message: "failed to save token" + err.Error(),
			Err:     err,
		}
	}
	return scToken, nil
}
