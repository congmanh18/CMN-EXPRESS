package delivery

import (
	"express_be/core/transport/http/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

// HandleDeleteDeliveryPerson deletes a delivery person (soft delete)
// @Summary Xóa tài khoản người giao hàng
// @Description Thực hiện xóa người giao hàng theo ID
// @Tags DeliveryPersons
// @Accept  json
// @Produce  json
// @Param   id    path      string  true  "Delivery Person ID"
// @Router /delivery-persons/{id} [delete]
func (h *handlerImpl) HandleDeleteDeliveryPerson(c echo.Context) error {
	// 1. Kiểm tra id từ url
	id := c.Param("id")
	if id == "" {
		return response.Error(c, http.StatusBadRequest, "id is required")
	}

	// 2. Xóa người giao hàng từ cơ sở dữ liệu
	if err := h.deliveryPersonUsecase.SoftDeleteDeliveryPerson(c.Request().Context(), &id); err != nil {
		return response.Error(c, http.StatusInternalServerError, "Failed to delete: "+err.Error())
	}

	// 3. Trả về kết quả thành công
	return response.OK(c, http.StatusOK, "Deleted successfully", nil)
}
