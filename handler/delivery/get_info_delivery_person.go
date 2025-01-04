package delivery

import (
	"express_be/core/transport/http/response"
	mapper "express_be/mapper/res"
	"net/http"

	"github.com/labstack/echo/v4"
)

// HandleGetInfoDeliveryPerson retrieves delivery person information
// @Summary Lấy thông tin người giao hàng
// @Description Truy xuất thông tin chi tiết về người giao hàng cụ thể theo ID
// @Tags DeliveryPersons
// @Accept  json
// @Produce  json
// @Param   id    path      string  true  "Delivery Person ID"
// @Router /delivery-persons/{id} [get]
func (h *handlerImpl) HandleGetInfoDeliveryPerson(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return response.Error(c, http.StatusBadRequest, "id is required")
	}

	deliveryPerson, usecaseErr := h.deliveryPersonUsecase.GetInfoDeliveryPerson(c.Request().Context(), &id)
	if usecaseErr != nil {
		return response.Error(
			c,
			usecaseErr.Code,
			usecaseErr.Message,
		)
	}

	resp := mapper.DeliveryPersonToRes(deliveryPerson)
	return response.OK(c, http.StatusOK, "Delivery person information", resp)
}
