package dashboard

import (
	"express_be/usecase/user"

	"github.com/labstack/echo/v4"
)

type Handler interface {
	HandleListUsers(c echo.Context) error
	HandleUpdateUserStatus(c echo.Context) error
}

type handlerImpl struct {
	userUsecase user.UserUsecase
}

type HandlerInject struct {
	UserUsecase user.UserUsecase
}

func NewHandler(inj HandlerInject) Handler {
	return &handlerImpl{
		userUsecase: inj.UserUsecase,
	}
}
