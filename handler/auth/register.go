package auth

import (
	"express_be/core/transport/http/response"
	mapper "express_be/mapper/req"
	model "express_be/model/req"
	"net/http"

	"github.com/labstack/echo/v4"
)

// HandleRegister handles register for different user roles
// @Summary Register
// @Description Register for different roles (admin, accounting, customer, delivery_person) account_type customer (prepaid, postpaid)
// @Description Example payload:
// @Description
//
//	@Description {
//	@Description 	"account_type": "prepaid",
//	@Description 	"current_address": "Shop Address of Customer",
//	@Description 	"date_of_birth": "23/10/2002",
//	@Description	"full_name": "Nguyen Cong Manh",
//	@Description	"gender": "Nam",
//	@Description	"identification_number": "052202014579",
//	@Description	"latitude": 37.7749,
//	@Description	"longtitude": 122.4194,
//	@Description	"nationality": "VN",
//	@Description	"password": "strongpassword123",
//	@Description	"phone": "0977683511",
//	@Description	"place_of_origin": "Hoài Sơn, Thị xã Hoài Nhơn, Bình Định",
//	@Description	"place_of_residence": "Thôn Phú Nông, Hoài Sơn, Hoài Nhơn, Bình Định",
//	@Description	"role": "customer"
//	@Description }
//
// @Tags Authentication
// @Accept json
// @Produce json
//
//	@Param register body req.RegisterRequest true "Register Request Example"
//
// @Router /register [post]
func (h handlerImpl) HandleRegister(c echo.Context) error {
	// 1. Parse dữ liệu đầu vào
	var req model.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return response.Error(c, http.StatusBadRequest, "Invalid request")
	}

	// 2. Validate dữ liệu
	if err := model.Validate.Struct(req); err != nil {
		return response.Error(c, http.StatusBadRequest, err.Error())
	}
	switch req.Role {
	case "admin":
		admin := mapper.ReqToAdmin(req)
		err := h.authUsecase.CreateAdmin(c.Request().Context(), admin)
		if err != nil {
			return response.Error(c, err.Code, err.Message)
		}
	case "customer":
		customer := mapper.ReqToCustomer(req)
		err := h.authUsecase.CreateCustomer(c.Request().Context(), customer)
		if err != nil {
			return response.Error(c, err.Code, err.Message)
		}
	case "delivery_person":
		deliveryPerson := mapper.ReqToDeliveryPerson(req)
		err := h.authUsecase.CreateDeliveryPerson(c.Request().Context(), deliveryPerson)
		if err != nil {
			return response.Error(c, err.Code, err.Message)
		}
	case "accounting":
		accounting := mapper.ReqToAccounting(req)
		err := h.authUsecase.CreateAccounting(c.Request().Context(), accounting)
		if err != nil {
			return response.Error(c, err.Code, err.Message)
		}
	case "driver":
	default:
		return response.Error(c, http.StatusUnauthorized, "Invalid role")

	}

	return response.OK(c, http.StatusOK, "Register successfully", nil)
}
