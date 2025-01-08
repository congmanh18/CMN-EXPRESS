package dashboard

import (
	"express_be/usecase/customer"
	"express_be/usecase/delivery"
	"express_be/usecase/user"

	"github.com/labstack/echo/v4"
)

type Handler interface {
	HandleListUsers(c echo.Context) error
	HandleUpdateUserStatus(c echo.Context) error
}

type handlerImpl struct {
	userUsecase           user.UserUsecase
	customerUsecase       customer.CustomerUsecase
	deliveryPersonUsecase delivery.DeliveryPersonUsecase
}

type HandlerInject struct {
	UserUsecase           user.UserUsecase
	CustomerUsecase       customer.CustomerUsecase
	DeliveryPersonUsecase delivery.DeliveryPersonUsecase
}

func NewHandler(inj HandlerInject) Handler {
	return &handlerImpl{
		userUsecase:           inj.UserUsecase,
		customerUsecase:       inj.CustomerUsecase,
		deliveryPersonUsecase: inj.DeliveryPersonUsecase,
	}
}
