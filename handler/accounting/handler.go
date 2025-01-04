package accounting

import (
	"express_be/usecase/accounting"
)

type Handler interface {
}

type handlerImpl struct {
	accountingUsecase accounting.AccountingUsecase
}

type HandlerInject struct {
	AccountingUsecase accounting.AccountingUsecase
}

func NewHandler(inj HandlerInject) Handler {
	return &handlerImpl{inj.AccountingUsecase}
}
