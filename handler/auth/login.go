package auth

import (
	handlerError "express_be/core/err"
	"express_be/core/transport/http/response"
	"express_be/model/req"
	"express_be/model/res"
	"net/http"

	"github.com/labstack/echo/v4"
)

// HandleLogin handles login for different user roles
// @Summary Login
// @Description Login for different roles (admin, customer, delivery_person, accounting)
// @Description Example customer payload:
//
//	@Description ``` {
//	@Description	"phone": "0977683511",
//	@Description	"password": "nguyenmanhcong"
//	@Description } ```
//
// @Tags Authentication
// @Accept json
// @Produce json
// @Param login body req.LoginRequest true "Login Request"
// @Router /login [post]
func (h *handlerImpl) HandleLogin(c echo.Context) error {
	var login req.LoginRequest
	if err := c.Bind(&login); err != nil {
		return response.Error(c, handlerError.ErrMissingField.Code, handlerError.ErrMissingField.Message)
	}

	if err := req.Validate.Struct(login); err != nil {
		return response.Error(c, handlerError.ErrInvalidFormat.Code, handlerError.ErrInvalidFormat.Message)
	}

	token, err := h.authUsecase.Login(c.Request().Context(), login.Phone, login.Password)
	if err != nil {
		return response.Error(c, err.Code, err.Message)
	}

	resp := &res.LoginRes{
		AccessToken:  *token.AccessToken,
		RefreshToken: *token.RefreshToken,
	}
	return response.OK(c, http.StatusOK, "success", resp)
}

// c.SetCookie(&http.Cookie{
// 	Name:     "refresh_token_id",
// 	Value:    *token.RefreshToken,
// 	HttpOnly: false, // Không cho phép JavaScript truy cập
// 	Secure:   false, // Chỉ gửi qua HTTPS
// 	Path:     "/",   // Áp dụng trên toàn bộ ứng dụng
// })
