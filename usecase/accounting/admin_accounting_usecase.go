package accounting

import (
	"express_be/repository/accounting"
)

type AdminUsecase interface {
}

type adminUsecaseImpl struct {
	repo accounting.Repo
}

func NewAdminUsecase(repo accounting.Repo) AdminUsecase {
	return &adminUsecaseImpl{repo: repo}
}
