package auth

import (
	"express_be/usecase/auth"

	"github.com/labstack/echo/v4"
)

type Handler interface {
	HandleCustomerRegistration(c echo.Context) error
	HandleLoginCustomer(c echo.Context) error
	HandleDeliveryPersonRegistration(c echo.Context) error
	HandleLoginDeliveryPerson(c echo.Context) error
	HandleRefreshToken(c echo.Context) error
}

type handlerImpl struct {
	authUsecase auth.AuthUsecase
}

type HandlerInject struct {
	AuthUsecase auth.AuthUsecase
}

func NewHandler(inj HandlerInject) Handler {
	return &handlerImpl{
		authUsecase: inj.AuthUsecase,
	}
}
