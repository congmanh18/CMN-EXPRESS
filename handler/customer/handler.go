package customer

import (
	"express_be/usecase/customer"

	"github.com/labstack/echo/v4"
)

type Handler interface {
	HandleGetInfoCustomer(c echo.Context) error
	HandleUpdateCustomer(c echo.Context) error
	HandleDeleteCustomer(c echo.Context) error
}

type handlerImpl struct {
	customerUsecase customer.CustomerUsecase
}

type HandlerInject struct {
	CustomerUsecase customer.CustomerUsecase
}

func NewHandler(inj HandlerInject) Handler {
	return &handlerImpl{customerUsecase: inj.CustomerUsecase}
}
