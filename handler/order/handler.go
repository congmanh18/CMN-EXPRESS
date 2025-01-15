package order

import (
	"express_be/usecase/order"

	"github.com/labstack/echo/v4"
)

type Handler interface {
	HandleCreate(c echo.Context) error
	HandleUpdateOrderStatus(c echo.Context) error
}

type handlerImpl struct {
	orderUsecase order.OrderUsecase
}

type HandlerInject struct {
	OrderUsecase order.OrderUsecase
}

func NewHandler(inj HandlerInject) Handler {
	return &handlerImpl{
		orderUsecase: inj.OrderUsecase,
	}
}
