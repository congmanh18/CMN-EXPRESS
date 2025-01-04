package admin

import (
	"express_be/core/transport/http/response"
	"express_be/model/req"
	"express_be/model/res"

	"net/http"

	"github.com/labstack/echo/v4"
)

// HandleLoginAdmin implements Handler.
// @Summary Login Admin
// @Description Authenticate Admin using their phone number and password. Returns an Access Token and sets a Refresh Token in an HttpOnly cookie.
// @Tags Authentication
// @Accept json
// @Produce json
// @Param login body req.LoginRequest true "Login Request"
// @Router /auth/admin/login [post]
func (h *handlerImpl) HandleLogin(c echo.Context) error {
	var login req.LoginRequest
	if err := c.Bind(&login); err != nil {
		return response.Error(c, http.StatusBadRequest, "Invalid request body")
	}

	// 2. Validate dữ liệu
	if err := req.Validate.Struct(login); err != nil {
		return response.Error(c, http.StatusBadRequest, err.Error())
	}

	token, err := h.adminUsecase.LoginAdmin(c.Request().Context(), login.Phone, login.Password)
	if err != nil {
		return response.Error(c, http.StatusUnauthorized, err.Error())
	}

	resp := &res.LoginRes{
		AccessToken: *token.AccessToken,
		UserID:      "adminadmin",
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
