package auth

import (
	"express_be/core/transport/http/response"
	mapper "express_be/mapper/req"
	model "express_be/model/req"
	"net/http"

	"github.com/labstack/echo/v4"
)

// HandleCustomerRegistration handles customer registration
// @Summary Đăng ký khách hàng mới
// @Description Đăng ký khách hàng mới cung cấp số tối thiểu "account_type", "phone" và "password"
// @Tags Customers
// @Accept  json
// @Produce  json
// @Param   request  body  model.RegisterRequest  true  "Customer Registration Request" example({"phone": "0912345678", "password": "abc@1234", "account_type": "prepaid"})
// @Router /customers/register [post]
func (h *handlerImpl) HandleCustomerRegistration(c echo.Context) error {
	// 1. Parse dữ liệu đầu vào
	var req model.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return response.Error(c, http.StatusBadRequest, "Invalid request")
	}

	// 2. Validate dữ liệu
	if err := model.Validate.Struct(req); err != nil {
		return response.Error(c, http.StatusBadRequest, err.Error())
	}

	// 3. Define default and hash password
	customer := mapper.RegisterToCustomer(req)

	// 4. Gọi đến usecase để xử lý logic
	err := h.authUsecase.CreateCustomerUsecase(c.Request().Context(), customer)
	if err != nil {
		return response.Error(c, err.Code, err.Message)
	}

	// 5. Trả về kết quả
	return response.OK(c, http.StatusOK, "Register successfully", nil)
}

// HandleDeliveryPersonRegistration handles delivery person registration
// @Summary Đăng ký người giao hàng mới
// @Description Đăng ký người giao hàng mới cung cấp tối thiểu "phone" và "password"
// @Tags DeliveryPersons
// @Accept  json
// @Produce  json
// @Param   request  body  model.RegisterRequest  true  "Delivery Person Registration Request" example({"phone": "0912345678", "password": "abc@1234"})
// @Router /delivery-persons/register [post]
func (h *handlerImpl) HandleDeliveryPersonRegistration(c echo.Context) error {
	// 1. Parse dữ liệu đầu vào
	var req model.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return response.Error(c, http.StatusBadRequest, "Invalid request")
	}

	// 2. Validate dữ liệu
	if err := model.Validate.Struct(req); err != nil {
		return response.Error(c, http.StatusBadRequest, err.Error())
	}

	// 3. Define default and hash password
	deliveryPerson := mapper.RegisterToDeliveryPerson(req)

	// 4. Gọi đến usecase để xử lý logic
	err := h.authUsecase.CreateDeliveryPersonUsecase(c.Request().Context(), deliveryPerson)
	if err != nil {
		return response.Error(c, err.Code, err.Message)
	}

	// 5. Trả về kết quả
	return response.OK(c, http.StatusOK, "Register successfully", nil)
}
