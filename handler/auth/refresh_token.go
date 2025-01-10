package auth

import (
	handlerError "express_be/core/err"
	"express_be/core/security"
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
	// 1. Lấy refresh_token_id từ cookie
	cookie, err := c.Cookie("refresh_token_id")
	if err != nil {
		return response.Error(c, handlerError.ErrTokenMissing.Code, handlerError.ErrTokenMissing.Message)
	}

	refreshTokenID := cookie.Value
	// 2. Kiểm tra RefreshToken trong database/cache
	userID, err := h.authUsecase.ValidateRefreshToken(c.Request().Context(), &refreshTokenID)
	if userID == nil {
		return response.Error(c, handlerError.ErrValidateTokenFailure.Code, handlerError.ErrValidateTokenFailure.Message)
	}

	// 3. Tạo AccessToken mới
	accessTokenDuration := time.Hour * 8
	newAccessToken, err := security.GenAccessToken(*userID, accessTokenDuration)
	if err != nil {
		return response.Error(c, handlerError.ErrGentokenFailure.Code, handlerError.ErrGentokenFailure.Message)
	}

	// 4. Trả AccessToken mới
	return response.OK(c, http.StatusOK, "success", map[string]interface{}{
		"access_token": newAccessToken,
	})
}
