package admin

import (
	"express_be/core/transport/http/response"
	mapper "express_be/mapper/res"
	model "express_be/model/res"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// HandleListPendingCustomer lists all customers
// @Summary Liệt kê tất cả khách hàng
// @Description Truy xuất danh sách tất cả khách hàng được phân trang
// @Tags Admin
// @Accept  json
// @Produce  json
// @Param   page       query   int    false  "Page number, defaults to 1"
// @Param   page_size  query   int    false  "Page size, defaults to 10"
// @Router /admin/customers/all [get]
func (h *handlerImpl) HandleAllCustomers(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(c.QueryParam("page_size"))
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	customers, usecaseErr := h.customerUsecase.GetAllCustomers(c.Request().Context(), &page, &pageSize)
	if err != nil {
		return response.Error(
			c,
			usecaseErr.Code,
			usecaseErr.Message,
		)
	}

	var customerResponses []model.CustomerRes
	for _, customer := range customers {
		customerRes := mapper.CustomerToRes(&customer)
		customerResponses = append(customerResponses, customerRes)
	}

	resp := model.CustomerPaginationResponse{
		Page:     page,
		PageSize: pageSize,
		List:     customerResponses,
	}

	return response.OK(c, http.StatusOK, "success", resp)
}
