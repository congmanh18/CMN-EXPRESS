package auth

import (
	"express_be/core/transport/http/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *handlerImpl) HandleVerifyCustomer(c echo.Context) error {
	// 1. Lấy authid từ context
	authid, ok := c.Get("authid").(string)
	if !ok || authid == "" {
		return response.Error(c, http.StatusUnauthorized, "unauthorized access")
	}

	// 2. Lấy uid từ URL parameter
	uid := c.Param("id")
	if uid == "" {
		return response.Error(c, http.StatusBadRequest, "invalid user id")
	}

	// 3. Update verification status
	if err := h.authUsecase.VerifyCustomerUsecase(c.Request().Context(), &uid); err != nil {
		return response.Error(c, http.StatusInternalServerError, "failed to verify customer")
	}

	// 4. Trả về kết quả
	return response.SimpleOK(c, http.StatusOK, nil)
}

func (h *handlerImpl) HandleVerifyDeliveryPerson(c echo.Context) error {
	// 1. Lấy authid từ context
	authid, ok := c.Get("authid").(string)
	if !ok || authid == "" {
		return response.Error(c, http.StatusUnauthorized, "unauthorized access")
	}

	// 2. Lấy uid từ URL parameter
	uid := c.Param("id")
	if uid == "" {
		return response.Error(c, http.StatusBadRequest, "invalid delivery person id")
	}

	// 3. Update verification status
	if err := h.authUsecase.VerifyDeliveryPersonUsecase(c.Request().Context(), &uid); err != nil {
		return response.Error(c, http.StatusInternalServerError, "failed to verify delivery person")
	}

	// 4. Trả về kết quả
	return response.SimpleOK(c, http.StatusOK, nil)
}
