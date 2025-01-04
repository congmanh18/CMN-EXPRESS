package auth

import (
	"express_be/core/transport/http/response"
	"express_be/model/req"
	"express_be/model/res"
	"net/http"

	"github.com/labstack/echo/v4"
)

// HandleLogin handles login for different user roles
// @Summary Login
// @Description Login for different roles (admin, customer, delivery_person, accounting)
// @Tags Authentication
// @Accept json
// @Produce json
// @Param login body req.LoginRequest true "Login Request"
// @Router /login [post]
func (h *handlerImpl) HandleLogin(c echo.Context) error {
	var login req.LoginRequest
	if err := c.Bind(&login); err != nil {
		return response.Error(c, http.StatusBadRequest, "Invalid request body")
	}

	if err := req.Validate.Struct(login); err != nil {
		return response.Error(c, http.StatusBadRequest, err.Error())
	}

	switch *login.Role {
	case "admin":
		token, err := h.authUsecase.LoginAdmin(c.Request().Context(), login.Phone, login.Password)
		if err != nil {
			return response.Error(c, err.Code, err.Message)
		}

		resp := &res.LoginRes{
			AccessToken: *token.AccessToken,
		}
		return response.OK(c, http.StatusOK, "Login successfully", resp)
	case "customer":
		token, user, err := h.authUsecase.LoginCustomer(c.Request().Context(), login.Phone, login.Password)
		if err != nil {
			return response.Error(c, err.Code, err.Message)
		}

		resp := &res.LoginRes{
			AccessToken: *token.AccessToken,
			UserID:      *user.ID,
		}
		return response.OK(c, http.StatusOK, "Login successfully", resp)
	case "delivery_person":
		token, user, err := h.authUsecase.LoginDeliveryPerson(c.Request().Context(), login.Phone, login.Password)
		if err != nil {
			return response.Error(c, err.Code, err.Message)
		}

		resp := &res.LoginRes{
			AccessToken: *token.AccessToken,
			UserID:      *user.ID,
		}
		return response.OK(c, http.StatusOK, "Login successfully", resp)
	case "accounting":
		token, user, err := h.authUsecase.LoginAccounting(c.Request().Context(), login.Phone, login.Password)
		if err != nil {
			return response.Error(c, err.Code, err.Message)
		}

		resp := &res.LoginRes{
			AccessToken: *token.AccessToken,
			UserID:      *user.ID,
		}
		return response.OK(c, http.StatusOK, "Login successfully", resp)
	default:
		return response.Error(c, http.StatusUnauthorized, "Invalid role")
	}
}

// c.SetCookie(&http.Cookie{
// 	Name:     "refresh_token_id",
// 	Value:    *token.RefreshToken,
// 	HttpOnly: false, // Không cho phép JavaScript truy cập
// 	Secure:   false, // Chỉ gửi qua HTTPS
// 	Path:     "/",   // Áp dụng trên toàn bộ ứng dụng
// })
