package customer

import (
	"express_be/core/transport/http/response"
	mapper "express_be/mapper/req"
	model "express_be/model/req"

	"net/http"

	"github.com/labstack/echo/v4"
)

// HandleUpdateCustomer updates customer information
// @Summary Cập nhật thông tin khách hàng
// @Description Cập nhật chi tiết của một khách hàng cụ thể theo ID
// @Description Example customer payload:
//
//	@Description ``` {
//	@Description 	"account_type": "prepaid",
//	@Description 	"current_address": "Shop Address of Customer",
//	@Description 	"date_of_birth": "23/10/2002",
//	@Description	"full_name": "Nguyen Cong Manh",
//	@Description	"gender": "Nam",
//	@Description	"identification_number": "052202014579",
//	@Description	"nationality": "VN",
//	@Description	"place_of_origin": "Hoài Sơn, Thị xã Hoài Nhơn, Bình Định",
//	@Description	"place_of_residence": "Thôn Phú Nông, Hoài Sơn, Hoài Nhơn, Bình Định",
//	@Description	"latitude": 37.7749,
//	@Description	"longtitude": 122.4194
//	@Description } ```
//
// @Tags Customers
// @Accept  json
// @Produce  json
// @Param   id      path      string  true  "Customer ID"
// @Param   request body      model.UpdateCustomerReq true  "Customer Update Request"
// @Router /customers/{id} [put]
func (h *handlerImpl) HandleUpdateCustomer(c echo.Context) error {
	// 1. Kiểm tra id từ url
	id := c.Param("id")
	if id == "" {
		return response.Error(c, http.StatusBadRequest, "id is required")
	}

	// 2. Parse dữ liệu đầu vào
	var req model.UpdateCustomerReq
	if err := c.Bind(&req); err != nil {
		return response.Error(c, http.StatusBadRequest, "Invalid request")
	}

	// 3. Validate dữ liệu
	if err := model.Validate.Struct(req); err != nil {
		return response.Error(c, http.StatusBadRequest, err.Error())
	}

	// 4. mapping req thành entity
	customer, user := mapper.UpdateToCustomer(req)

	// 5. Thực hiện cập nhật
	if err := h.customerUsecase.UpdateInfoCustomer(c.Request().Context(), user, customer, &id); err != nil {
		return response.Error(c, http.StatusInternalServerError, "Failed to update: "+err.Error())
	}

	// 6. Trả về kết quả
	return response.OK(c, http.StatusOK, "Updated successfully", nil)
}
