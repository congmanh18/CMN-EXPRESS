package admin

import (
	"express_be/core/transport/http/response"
	mapper "express_be/mapper/res"
	model "express_be/model/res"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// HandleListPendingDeliveryPerson lists all pending delivery persons
// @Summary Liệt kê những người giao hàng đang "PENDING"
// @Description Truy xuất danh sách được phân trang của những người giao hàng với trạng thái "PENDING"
// @Tags Admin
// @Accept  json
// @Produce  json
// @Param   page       query   int    false  "Page number, defaults to 1"
// @Param   page_size  query   int    false  "Page size, defaults to 10"
// @Router /admin/delivery-persons/pending [get]
func (h handlerImpl) HandleListPendingDeliveryPerson(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page < 1 {
		page = 1 // Mặc định là trang đầu tiên nếu không hợp lệ
	}

	pageSize, err := strconv.Atoi(c.QueryParam("page_size"))
	if err != nil || pageSize < 1 {
		pageSize = 10 // Mặc định 10 bản ghi mỗi trang nếu không hợp lệ
	}

	deliveryPersons, usecaseErr := h.deliveryPersonUsecase.GetPendingDeliveryPersons(c.Request().Context(), &page, &pageSize)
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
