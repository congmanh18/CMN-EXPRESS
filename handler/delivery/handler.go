package delivery

import (
	"express_be/usecase/delivery"

	"github.com/labstack/echo/v4"
)

type Handler interface {
	HandleGetInfoDeliveryPerson(c echo.Context) error
	HandleUpdateDeliveryPerson(c echo.Context) error
	HandleDeleteDeliveryPerson(c echo.Context) error
	HandleChangePassword(c echo.Context) error
}

type handlerImpl struct {
	deliveryPersonUsecase delivery.DeliveryPersonUsecase
}

type HandlerInject struct {
	DeliveryPersonUsecase delivery.DeliveryPersonUsecase
}

func NewHandler(inj HandlerInject) Handler {
	return &handlerImpl{
		deliveryPersonUsecase: inj.DeliveryPersonUsecase,
	}
}
