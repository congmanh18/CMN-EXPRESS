package auth

import (
	handlerError "express_be/core/err"

	"express_be/core/transport/http/response"
	mapper "express_be/mapper/req"
	model "express_be/model/req"
	"net/http"

	"github.com/labstack/echo/v4"
)

// HandleRegister handles register for different user roles
// @Summary Register
// @Description Register for different roles (admin, accounting, customer, delivery_person) account_type customer (prepaid, postpaid)
// @Description Example customer payload:
//
//	@Description {
//	@Description     "phone": "0990123456",
//	@Description     "password": "wrightdaniel",
//	@Description     "last_name": "Wright",
//	@Description     "middle_name": "Thomas",
//	@Description     "first_name": "Daniel",
//	@Description     "specific_address": "789 Tran Hung Dao",
//	@Description     "ward": "Duong Dong",
//	@Description     "district": "Duong Dong",
//	@Description     "city": "Phu Quoc",
//	@Description     "identification_number": "321987654",
//	@Description     "gender": "male",
//	@Description     "date_of_birth": "1994-09-09",
//	@Description     "nationality": "Vietnamese",
//	@Description     "place_of_origin": "Phu Quoc",
//	@Description     "place_of_residence": "Phu Quoc",
//	@Description     "latitude": 10.2899,
//	@Description     "longtitude": 103.9840,
//	@Description     "account_type": "postpaid",
//	@Description     "role": "customer"
//	@Description   }
//
// @Tags Sprint1
// @Accept json
// @Produce json
// @Param register body req.RegisterRequest true "Register Request Example"
// @Router /register [post]
func (h handlerImpl) HandleRegister(c echo.Context) error {
	// 1. Parse dữ liệu đầu vào
	var req model.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return response.Error(c, handlerError.ErrMissingField.Code, handlerError.ErrMissingField.Message)
	}

	// 2. Validate dữ liệu
	if err := model.Validate.Struct(req); err != nil {
		return response.Error(c, handlerError.ErrInvalidFormat.Code, handlerError.ErrInvalidFormat.Message)
	}
	switch req.Role {
	case "admin":
		admin, user := mapper.RegisterToAdmin(req)
		err := h.authUsecase.CreateAdmin(c.Request().Context(), user, admin)
		if err != nil {
			return response.Error(c, err.Code, err.Message)
		}
	case "customer":
		customer, user := mapper.RegisterToCustomer(req)
		err := h.authUsecase.CreateCustomer(c.Request().Context(), user, customer)
		if err != nil {
			return response.Error(c, err.Code, err.Message)
		}
	case "delivery_person":
		deliveryPerson, user := mapper.RegisterToDeliveryPerson(req)
		err := h.authUsecase.CreateDeliveryPerson(c.Request().Context(), user, deliveryPerson)
		if err != nil {
			return response.Error(c, err.Code, err.Message)
		}
	case "accounting":
		accounting, user := mapper.RegisterToAccounting(req)
		err := h.authUsecase.CreateAccounting(c.Request().Context(), user, accounting)
		if err != nil {
			return response.Error(c, err.Code, err.Message)
		}
	default:
		return response.Error(c, handlerError.ErrInvalidRole.Code, handlerError.ErrInvalidRole.Message)
	}

	return response.OK(c, http.StatusOK, "success", nil)
}
