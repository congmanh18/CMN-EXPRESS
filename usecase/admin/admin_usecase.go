package admin

import (
	"express_be/repository/admin"
)

type AdminUsecase interface {
}

type adminUsecaseImpl struct {
	repo admin.Repo
}

func NewAdminUsecase(repo admin.Repo) AdminUsecase {
	return &adminUsecaseImpl{
		repo: repo,
	}
}
