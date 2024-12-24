package auth

import (
	"express_be/usecase/auth"

	"github.com/labstack/echo/v4"
)

type Handler interface {
	HandleCustomerRegistration(c echo.Context) error
	HandleDeliveryPersonRegistration(c echo.Context) error
}

type handlerImpl struct {
	registerCustomerUseCase       auth.RegisterCustomerUseCase
	registerDeliveryPersonUseCase auth.RegisterDeliveryPersonUseCase
}

type HandlerInject struct {
	RegisterCustomerUseCase       auth.RegisterCustomerUseCase
	RegisterDeliveryPersonUseCase auth.RegisterDeliveryPersonUseCase
}

func NewHandler(inj HandlerInject) Handler {
	return &handlerImpl{
		registerCustomerUseCase:       inj.RegisterCustomerUseCase,
		registerDeliveryPersonUseCase: inj.RegisterDeliveryPersonUseCase,
	}
}
