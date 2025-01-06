package auth

import (
	"express_be/usecase/auth"

	"github.com/labstack/echo/v4"
)

type Handler interface {
	HandleRegister(c echo.Context) error
	HandleLogin(c echo.Context) error
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
