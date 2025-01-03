package admin

import (
	"express_be/core/transport/http/response"
	mapper "express_be/mapper/res"
	model "express_be/model/res"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// HandleListPendingDeliveryPerson lists all delivery persons
// @Summary Liệt kê tất cả những người giao hàng
// @Description Truy xuất danh sách được phân trang của tất cả những người giao hàng
// @Tags Admin
// @Accept  json
// @Produce  json
// @Param   page       query   int    false  "Page number, defaults to 1"
// @Param   page_size  query   int    false  "Page size, defaults to 10"
// @Router /admin/delivery-persons/all [get]
func (h *handlerImpl) HandleAllDeliveryPersons(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(c.QueryParam("page_size"))
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	deliveryPersons, usecaseErr := h.adminDeliveryPersonUsecase.AdminGetAllDeliveryPersons(c.Request().Context(), &page, &pageSize)
	if usecaseErr != nil {
		return response.Error(
			c,
			usecaseErr.Code,
			usecaseErr.Message,
		)
	}

	var deliveryPersonResponses []model.DeliveryPersonRes
	for _, deliveryPerson := range deliveryPersons {
		deliveryPersonResponse := mapper.DeliveryPersonToRes(&deliveryPerson)
		deliveryPersonResponses = append(deliveryPersonResponses, deliveryPersonResponse)
	}

	resp := model.DeliveryPaginationResponse{
		Page:     page,
		PageSize: pageSize,
		List:     deliveryPersonResponses,
	}

	return response.OK(c, http.StatusOK, "success", resp)
}
