package auth

import (
	"express_be/auth"
	"express_be/core/transport/http/response"
	mapper "express_be/mapper/req"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *handlerImpl) HandleCustomerRegistration(c echo.Context) error {
	// 1. Parse dữ liệu đầu vào
	var req auth.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return response.Error(c, http.StatusBadRequest, "Invalid request")
	}

	// 2. Kiểm tra các trường bắt buộc
	if req.Phone == "" || req.Password == "" {
		return response.Error(c, http.StatusBadRequest, "Missing required fields")
	}

	// 3. Define default and hash password
	customer := mapper.RegisterToCustomer(req)

	// 4. Gọi đến usecase để xử lý logic
	err := h.registerCustomerUseCase.CreateCustomerUsecase(c.Request().Context(), customer)
	if err != nil {
		return response.Error(c, err.Code, err.Message)
	}

	// 5. Trả về kết quả
	return response.OK(c, http.StatusOK, "Register success", nil)
}

func (h *handlerImpl) HandleDeliveryPersonRegistration(c echo.Context) error {
	// 1. Parse dữ liệu đầu vào
	var req auth.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return response.Error(c, http.StatusBadRequest, "Invalid request")
	}

	// 2. Kiểm tra các trường bắt buộc
	if req.Phone == "" || req.Password == "" {
		return response.Error(c, http.StatusBadRequest, "Missing required fields")
	}

	// 3. Define default and hash password
	deliveryPerson := mapper.RegisterToDeliveryPerson(req)

	// 4. Gọi đến usecase để xử lý logic
	err := h.registerDeliveryPersonUseCase.CreateDeliveryPersonUsecase(c.Request().Context(), deliveryPerson)
	if err != nil {
		return response.Error(c, err.Code, err.Message)
	}

	// 5. Trả về kết quả
	return response.OK(c, http.StatusOK, "Register success", nil)
}
