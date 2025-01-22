package user

import (
	handlerError "express_be/core/err"
	"express_be/core/transport/http/response"
	mapper "express_be/mapper/res"
	model "express_be/model/res"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// HandleListUsers implements Handler.
// @Summary      Fetch paginated users
// @Description  Get a list of users (customers and delivery persons) with optional filters by status and role, including pagination.
// @Tags         User-information
// @Accept       json
// @Produce      json
// @Param Authorization header string true "Bearer token" default(Bearer <access-token>)
// @Param        page        query     int     false  "Page number (default is 1)"       default(1)
// @Param        page_size   query     int     false  "Page size (default is 10)"       default(10)
// @Param        status      query     string  false  "Filter by customer status (e.g., pending, verified, blocked, active, inactive) Filter by delivery_person (e.g., on_duty, off_duty)"
// @Param        role        query     string  true   "Filter by user role (e.g., customer, delivery_person)"
// @Security BearerAuth
// @Router       /users [get]
func (h *handlerImpl) HandleListUsers(c echo.Context) error {
	roleCheck, ok := c.Get("role").(string)
	if !ok {
		return response.Error(c, handlerError.ErrTokenMissing.Code, handlerError.ErrTokenMissing.Message)
	}

	if roleCheck != "admin" && roleCheck != "accounting" {
		return response.Error(c, handlerError.ErrAccessDenied.Code, handlerError.ErrAccessDenied.Message)
	}

	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(c.QueryParam("page_size"))
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	status := c.QueryParam("status")
	role := c.QueryParam("role")
	if role == "" {
		return response.Error(c, handlerError.ErrMissingField.Code, handlerError.ErrMissingField.Message)
	}

	customers, deliveryPersons, usecaseErr := h.userUsecase.GetUsers(c.Request().Context(), &status, &role, &page, &pageSize)
	if usecaseErr != nil {
		return response.Error(
			c,
			usecaseErr.Code,
			usecaseErr.Message)
	}

	var customerResponses []model.CustomerRes
	for _, customer := range customers {
		customerRes := mapper.CustomerToRes(&customer)
		customerResponses = append(customerResponses, customerRes)
	}

	var deliveryPersonResponses []model.DeliveryPersonRes
	for _, deliveryPerson := range deliveryPersons {
		deliveryPersonResponse := mapper.DeliveryPersonToRes(&deliveryPerson)
		deliveryPersonResponses = append(deliveryPersonResponses, deliveryPersonResponse)
	}

	resp := model.PaginationResponse{
		Page:            page,
		PageSize:        pageSize,
		Customers:       customerResponses,
		DeliveryPersons: deliveryPersonResponses,
	}

	return response.OK(c, http.StatusOK, "success", resp)

}
