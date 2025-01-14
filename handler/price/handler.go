package price

import (
	"express_be/usecase/price"

	"github.com/labstack/echo/v4"
)

type Handler interface {
	HandleCreate(c echo.Context) error
	HandleRead(c echo.Context) error
	HandleUpdate(c echo.Context) error
	HandleDelete(c echo.Context) error
}

type handlerImpl struct {
	priceUsecase price.PriceUsecase
}

type HandlerInject struct {
	PriceUsecase price.PriceUsecase
}

func NewHandler(inj HandlerInject) Handler {
	return &handlerImpl{
		priceUsecase: inj.PriceUsecase,
	}
}
