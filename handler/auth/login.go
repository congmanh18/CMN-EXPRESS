package auth

import (
	"express_be/core/transport/http/response"
	"express_be/model/req"
	"express_be/model/res"
	"net/http"

	"github.com/labstack/echo/v4"
)

// HandleLoginCustomer godoc
// @Summary Login Customer
// @Description Authenticate a customer using their phone number and password. Returns an Access Token and sets a Refresh Token in an HttpOnly cookie.
// @Tags Authentication
// @Accept json
// @Produce json
// @Param login body req.LoginRequest true "Login Request"
// @Router /auth/customers/login [post]
func (h *handlerImpl) HandleLoginCustomer(c echo.Context) error {
	var login req.LoginRequest
	if err := c.Bind(&login); err != nil {
		return response.Error(c, http.StatusBadRequest, "Invalid request body")
	}

	// 2. Validate dữ liệu
	if err := req.Validate.Struct(login); err != nil {
		return response.Error(c, http.StatusBadRequest, err.Error())
	}

	// 3. Xử lý đăng nhập
	token, customer, err := h.authUsecase.LoginCustomer(c.Request().Context(), login.Phone, login.Password)
	if err != nil {
		return response.Error(c, err.Code, err.Message)
	}

	resp := &res.LoginRes{
		AccessToken: *token.AccessToken,
		UserID:      *customer.ID,
	}

	c.SetCookie(&http.Cookie{
		Name:     "refresh_token_id",
		Value:    *token.RefreshToken,
		HttpOnly: false, // Không cho phép JavaScript truy cập
		Secure:   false, // Chỉ gửi qua HTTPS
		Path:     "/",   // Áp dụng trên toàn bộ ứng dụng
	})

	return response.OK(c, http.StatusOK, "Login successfully", resp)
}

// Handle Delivery Persons
// @Summary Login Delivery Persons
// @Description Authenticate Delivery Persons using their phone number and password. Returns an Access Token and sets a Refresh Token in an HttpOnly cookie.
// @Tags Authentication
// @Accept json
// @Produce json
// @Param login body req.LoginRequest true "Login Request"
// @Router /auth/delivery-persons/login [post]
func (h *handlerImpl) HandleLoginDeliveryPerson(c echo.Context) error {
	var login req.LoginRequest
	if err := c.Bind(&login); err != nil {
		return response.Error(c, http.StatusBadRequest, "Invalid request body")
	}

	// 2. Validate dữ liệu
	if err := req.Validate.Struct(login); err != nil {
		return response.Error(c, http.StatusBadRequest, err.Error())
	}

	token, deliveryPerson, err := h.authUsecase.LoginDeliveryPerson(c.Request().Context(), login.Phone, login.Password)
	if err != nil {
		return response.Error(c, err.Code, err.Message)
	}

	resp := &res.LoginRes{
		AccessToken: *token.AccessToken,
		UserID:      *deliveryPerson.ID,
	}

	c.SetCookie(&http.Cookie{
		Name:     "refresh_token_id",
		Value:    *token.RefreshToken,
		HttpOnly: false, // Không cho phép JavaScript truy cập
		Secure:   false, // Chỉ gửi qua HTTPS
		Path:     "/",   // Áp dụng trên toàn bộ ứng dụng
	})

	return response.OK(c, http.StatusOK, "Login successfully", resp)
}

// HandleLoginAccounting implements Handler.
// @Summary Login Accounting
// @Description Authenticate Accounting using their phone number and password. Returns an Access Token and sets a Refresh Token in an HttpOnly cookie.
// @Tags Authentication
// @Accept json
// @Produce json
// @Param login body req.LoginRequest true "Login Request"
// @Router /auth/accountings/login [post]
func (h *handlerImpl) HandleLoginAccounting(c echo.Context) error {
	var login req.LoginRequest
	if err := c.Bind(&login); err != nil {
		return response.Error(c, http.StatusBadRequest, "Invalid request body")
	}

	// 2. Validate dữ liệu
	if err := req.Validate.Struct(login); err != nil {
		return response.Error(c, http.StatusBadRequest, err.Error())
	}

	token, accouting, err := h.authUsecase.LoginAccounting(c.Request().Context(), login.Phone, login.Password)
	if err != nil {
		return response.Error(c, err.Code, err.Message)
	}

	resp := &res.LoginRes{
		AccessToken: *token.AccessToken,
		UserID:      *accouting.ID,
	}

	c.SetCookie(&http.Cookie{
		Name:     "refresh_token_id",
		Value:    *token.RefreshToken,
		HttpOnly: false, // Không cho phép JavaScript truy cập
		Secure:   false, // Chỉ gửi qua HTTPS
		Path:     "/",   // Áp dụng trên toàn bộ ứng dụng
	})

	return response.OK(c, http.StatusOK, "Login successfully", resp)
}
