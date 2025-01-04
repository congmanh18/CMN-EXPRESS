package admin

import (
	"express_be/usecase/accounting"
	"express_be/usecase/admin"
	"express_be/usecase/customer"
	"express_be/usecase/delivery"

	"github.com/labstack/echo/v4"
)

type Handler interface {
	HandleRegister(c echo.Context) error
	HandleListPendingCustomer(c echo.Context) error
	HandleListPendingDeliveryPerson(c echo.Context) error
	HandleAllCustomers(c echo.Context) error
	HandleAllDeliveryPersons(c echo.Context) error
	HandleUpdateCustomerStatus(c echo.Context) error
	HandleUpdateDeliveryPersonStatus(c echo.Context) error
	HandleRegisterAccounting(c echo.Context) error
}

type handlerImpl struct {
	adminCustomerUsecase       customer.AdminUsecase
	adminDeliveryPersonUsecase delivery.AdminUsecase
	adminAccountingUsecase     accounting.AdminUsecase
	adminUsecase               admin.AdminUsecase
}

type HandlerInject struct {
	AdminCustomerUsecase       customer.AdminUsecase
	AdminDeliveryPersonUsecase delivery.AdminUsecase
	AdminAccountingUsecase     accounting.AdminUsecase
	AdminUsecase               admin.AdminUsecase
}

func NewHandler(inj HandlerInject) Handler {
	return &handlerImpl{
		adminCustomerUsecase:       inj.AdminCustomerUsecase,
		adminDeliveryPersonUsecase: inj.AdminDeliveryPersonUsecase,
		adminAccountingUsecase:     inj.AdminAccountingUsecase,
		adminUsecase:               inj.AdminUsecase,
	}
}
