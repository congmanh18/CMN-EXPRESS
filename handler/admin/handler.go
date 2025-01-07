package admin

import (
	"express_be/usecase/customer"
	"express_be/usecase/delivery"
	"express_be/usecase/user"

	"github.com/labstack/echo/v4"
)

type Handler interface {
	HandleListPendingCustomer(c echo.Context) error
	HandleListPendingDeliveryPerson(c echo.Context) error
	HandleAllCustomers(c echo.Context) error
	HandleAllDeliveryPersons(c echo.Context) error
	HandleUpdateStatus(c echo.Context) error
}

type handlerImpl struct {
	adminUserUsecase      user.UserUsecase
	customerUsecase       customer.CustomerUsecase
	deliveryPersonUsecase delivery.DeliveryPersonUsecase
}

type HandlerInject struct {
	AdminUserUsecase      user.UserUsecase
	CustomerUsecase       customer.CustomerUsecase
	DeliveryPersonUsecase delivery.DeliveryPersonUsecase
}

func NewHandler(inj HandlerInject) Handler {
	return &handlerImpl{
		adminUserUsecase:      inj.AdminUserUsecase,
		customerUsecase:       inj.CustomerUsecase,
		deliveryPersonUsecase: inj.DeliveryPersonUsecase,
	}
}
