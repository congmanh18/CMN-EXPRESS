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

// HandleSearch implements Handler.
// @Summary      Search paginated users
// @Description  Search a list of users (customers and delivery persons) with optional filters by status and role, including pagination.
// @Tags         Sprint1
// @Accept       json
// @Produce      json
// @Param        page        query     int     false  "Page number (default is 1)"       default(1)
// @Param        page_size   query     int     false  "Page size (default is 10)"       default(10)
// @Param        status      query     string  false  "Filter by customer status (e.g., pending, verified, blocked, active, inactive) Filter by delivery_person (e.g., on_duty, off_duty)"
// @Param        name      query     string  false  "Search by user name"
// @Param        phone      query     string  false  "Search by user phone"
// @Param        role        query     string  true   "Filter by user role (e.g., customer, delivery_person)"
// @Router       /search [get]
func (h *handlerImpl) HandleSearch(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(c.QueryParam("page_size"))
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	phone := c.QueryParam("phone")
	name := c.QueryParam("name")
	status := c.QueryParam("status")

	role := c.QueryParam("role")
	if role == "" {
		return response.Error(c, handlerError.ErrMissingField.Code, handlerError.ErrMissingField.Message)
	}

	customers, deliveryPersons, usecaseErr := h.userUsecase.Search(c.Request().Context(), &role, &phone, &name, &status, &page, &pageSize)
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
	resp := model.UserPaginationResponse{
		Customers:       customerResponses,
		DeliveryPersons: deliveryPersonResponses,
	}

	return response.OK(c, http.StatusOK, "success", resp)

}
