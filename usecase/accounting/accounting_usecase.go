package accounting

import (
	"express_be/repository/accounting"
)

type AccountingUsecase interface {
}

type accountingUsecaseImpl struct {
	repo accounting.Repo
}

func NewAccountingUsecase(repo accounting.Repo) AccountingUsecase {
	return &accountingUsecaseImpl{repo}
}
