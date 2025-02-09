package user

import (
	handlerError "express_be/core/err"
	"express_be/core/transport/http/response"
	mapper "express_be/mapper/req"
	"net/http"

	model "express_be/model/req"

	"github.com/labstack/echo/v4"
)

// HandleUpdateInfo handles update info for different user roles
// @Summary Update
// @Description Update for different roles (customer, delivery_person) account_type customer (prepaid, postpaid)
// @Description Example user payload:
// @Tags User-information
// @Accept json
// @Produce json
// @Param id path string true "UserID"
// @Param update_info body model.UpdateUserReq true "Update Request Example"
// @Security BearerAuth
// @Router /users/{id} [put]
func (h handlerImpl) HandleUpdateInfo(c echo.Context) error {
	roleCheck, ok := c.Get("role").(string)
	if !ok {
		return response.Error(c, handlerError.ErrTokenMissing.Code, handlerError.ErrTokenMissing.Message)
	}

	if roleCheck == "" {
		return response.Error(c, handlerError.ErrAccessDenied.Code, handlerError.ErrAccessDenied.Message)
	}

	id := c.Param("id")
	if id == "" {
		return response.Error(c, handlerError.ErrMissingField.Code, handlerError.ErrMissingField.Message)
	}

	var req model.UpdateUserReq
	if err := c.Bind(&req); err != nil {
		return response.Error(c, handlerError.ErrMissingField.Code, handlerError.ErrMissingField.Message)
	}

	if err := model.Validate.Struct(req); err != nil {
		return response.Error(c, handlerError.ErrInvalidFormat.Code, handlerError.ErrInvalidFormat.Message)
	}
	switch req.Role {
	case "customer":
		customer, user := mapper.UpdateToCustomer(req)
		err := h.userUsecase.UpdateCustomerInfo(c.Request().Context(), &id, user, customer)
		if err != nil {
			return response.Error(c, err.Code, err.Message)
		}
	case "delivery_person":
		delivery, user := mapper.UpdateToDeliveryPerson(req)
		err := h.userUsecase.UpdateDeliveryPersonInfo(c.Request().Context(), &id, user, delivery)
		if err != nil {
			return response.Error(c, err.Code, err.Message)
		}
	default:
		return response.Error(c, handlerError.ErrInvalidRole.Code, handlerError.ErrInvalidRole.Message)
	}
	return response.OK(c, http.StatusOK, "success", nil)

}
