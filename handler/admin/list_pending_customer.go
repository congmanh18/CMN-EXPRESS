package admin

import (
	"express_be/core/transport/http/response"
	mapper "express_be/mapper/res"
	model "express_be/model/res"

	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// HandleListPendingCustomer lists all pending customers
// @Summary Liệt kê khách hàng đang "PENDING"
// @Description Truy xuất danh sách khách hàng được phân trang với trạng thái "PENDING"
// @Tags Admin
// @Accept  json
// @Produce  json
// @Param   page       query   int    false  "Page number, defaults to 1"
// @Param   page_size  query   int    false  "Page size, defaults to 10"
// @Router /admin/customers/pending [get]
func (h handlerImpl) HandleListPendingCustomer(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page < 1 {
		page = 1 // Mặc định là trang đầu tiên nếu không hợp lệ
	}

	pageSize, err := strconv.Atoi(c.QueryParam("page_size"))
	if err != nil || pageSize < 1 {
		pageSize = 10 // Mặc định 10 bản ghi mỗi trang nếu không hợp lệ
	}

	customers, usecaseErr := h.adminCustomerUsecase.AdminGetPendingCustomers(c.Request().Context(), &page, &pageSize)
	if usecaseErr != nil {
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
