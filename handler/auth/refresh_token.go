package auth

import (
	handlerError "express_be/core/err"
	"express_be/core/transport/http/response"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// HandleRefreshToken godoc
// @Summary Refresh Access Token
// @Description Refresh Access Token using a valid Refresh Token from headers.
// @Tags Authentication
// @Accept json
// @Produce json
// @Param refresh_token_id header string true "The refresh token identifier stored in headers"
// @Router /refresh-token [post]
func (h *handlerImpl) HandleRefreshToken(c echo.Context) error {
	// 1. Lấy Refresh Token từ cookie
	cookie, err := c.Cookie("refresh_token_id")
	if err != nil || cookie.Value == "" {
		return response.Error(c, handlerError.ErrTokenMissing.Code, handlerError.ErrTokenMissing.Message)
	}

	refreshTokenID := cookie.Value

	// 2. Xác thực Refresh Token trong cơ sở dữ liệu/cache
	userID, role, usecaseErr := h.authUsecase.ValidateRefreshToken(c.Request().Context(), &refreshTokenID)
	if usecaseErr != nil {
		return response.Error(c, usecaseErr.Code, usecaseErr.Message)
	}

	// 3. Tạo Access Token mới
	accessTokenDuration := time.Hour * 8
	newAccessToken, err := h.authUsecase.GenAccessToken(*userID, *role, accessTokenDuration)
	if err != nil {
		return response.Error(c, handlerError.ErrGentokenFailure.Code, handlerError.ErrGentokenFailure.Message)
	}

	// 4. Trả Access Token mới
	return response.OK(c, http.StatusOK, "success", newAccessToken)
}
