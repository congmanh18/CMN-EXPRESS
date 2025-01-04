package admin

import (
	"express_be/core/transport/http/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

// HandleVerifyCustomer implements Handler.
// @Summary Cập nhật trạng thái khách hàng
// @Description Cập nhật trạng thái của một khách hàng dựa trên ID
// @Tags Admin
// @Accept json
// @Produce json
// @Param id path string true "Customer ID"
// @Param status query string true "Trạng thái mới của khách hàng (active, inactive, pending, blocked, suspended, verified)"
// @Router /admin/customers/{id} [patch]
func (h *handlerImpl) HandleUpdateCustomerStatus(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return response.Error(c, http.StatusBadRequest, "id is required")
	}

	status := c.QueryParam("status")
	if status == "" {
		return response.Error(c, http.StatusBadRequest, "status is required")
	}

	usecaseErr := h.adminCustomerUsecase.AdminUpdateStatusCustomer(c.Request().Context(), &id, &status)
	if usecaseErr != nil {
		return response.Error(c, usecaseErr.Code, usecaseErr.Message)
	}

	return response.OK(c, http.StatusOK, "status update successful", nil)
}
