package delivery

import (
	"express_be/core/transport/http/response"
	mapper "express_be/mapper/req"
	model "express_be/model/req"

	"net/http"

	"github.com/labstack/echo/v4"
)

// HandleUpdateDeliveryPerson updates delivery person information
// @Summary Cập nhật thông tin người giao hàng
// @Description Cập nhật chi tiết người giao hàng cụ thể theo ID
// @Tags DeliveryPersons
// @Accept  json
// @Produce  json
// @Param   id      path      string                       true  "Delivery Person ID"
// @Param   request body      model.UpdateDeliveryPersonReq true  "Delivery Person Update Request"
// @Router /delivery-persons/{id} [put]
func (h *handlerImpl) HandleUpdateDeliveryPerson(c echo.Context) error {
	// 1. Kiểm tra id từ url
	id := c.Param("id")
	if id == "" {
		return response.Error(c, http.StatusBadRequest, "id is required")
	}

	// 2. Parse dữ liệu đầu vào
	var req model.UpdateDeliveryPersonReq
	if err := c.Bind(&req); err != nil {
		return response.Error(c, http.StatusBadRequest, "Invalid request body")
	}

	// 3. Validate dữ liệu
	if err := model.Validate.Struct(req); err != nil {
		return response.Error(c, http.StatusBadRequest, err.Error())
	}

	// 4. mapping req thành entity
	deliveryPerson, user := mapper.UpdateToDeliveryPerson(req)

	// 5. Thực hiện cập nhật
	if err := h.deliveryPersonUsecase.UpdateInfoDeliveryPerson(c.Request().Context(), user, deliveryPerson, &id); err != nil {
		return response.Error(c, http.StatusInternalServerError, err.Error())
	}

	// 6. Trả về kết quả
	return response.OK(c, http.StatusOK, "Updated successfully", nil)
}
