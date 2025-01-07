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
// @Param approval-status query string true "Trạng thái mới của khách hàng (accept, deny)"
// @Router /admin/users/{id} [patch]
func (h *handlerImpl) HandleUpdateStatus(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return response.Error(c, http.StatusBadRequest, "id is required")
	}

	status := c.QueryParam("approval-status")
	if status == "" {
		return response.Error(c, http.StatusBadRequest, "status is required")
	}

	usecaseErr := h.adminUserUsecase.UpdateStatus(c.Request().Context(), &id, &status)
	if usecaseErr != nil {
		return response.Error(c, usecaseErr.Code, usecaseErr.Message)
	}

	return response.OK(c, http.StatusOK, "status update successful", nil)
}
