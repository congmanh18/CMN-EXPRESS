package admin

import (
	"express_be/core/transport/http/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

// HandleUpdateDeliveryPersonStatus implements Handler.
// @Summary Cập nhật trạng thái của người giao hàng
// @Description Cập nhật trạng thái của một người giao hàng dựa trên ID
// @Tags Admin
// @Accept json
// @Produce json
// @Param id path string true "DeliveryPerson ID"
// @Param approval-status query string true "Trạng thái mới của người giao hàng (accept, deny)"
// @Router /admin/delivery-persons/{id} [patch]
func (h *handlerImpl) HandleUpdateDeliveryPersonStatus(c echo.Context) error {
	// Lấy ID từ path
	id := c.Param("id")
	if id == "" {
		return response.Error(c, http.StatusBadRequest, "id is required")
	}

	status := c.QueryParam("approval-status")
	if status == "" {
		return response.Error(c, http.StatusBadRequest, "status is required")
	}

	usecaseErr := h.adminDeliveryPersonUsecase.AdminUpdateStatusDeliveryPerson(c.Request().Context(), &id, &status)
	if usecaseErr != nil {
		return response.Error(c, usecaseErr.Code, usecaseErr.Message)
	}

	return response.OK(c, http.StatusOK, "status update successful", nil)
}
