package admin

import (
	"express_be/auth"
	core "express_be/core/firebase"
	"express_be/core/transport/http/response"

	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type AdminHandler struct {
	firebaseAuth *core.FirebaseAuthService
	jwtSecret    string
}

func NewAdminHandler(firebaseAuth *core.FirebaseAuthService, jwtSecret string) *AdminHandler {
	return &AdminHandler{
		firebaseAuth: firebaseAuth,
		jwtSecret:    jwtSecret,
	}
}

func (h *AdminHandler) Login(c echo.Context) error {
	// Lấy Firebase Token từ body request
	var req struct {
		FirebaseToken string `json:"firebaseToken"`
	}
	if err := c.Bind(&req); err != nil {
		return response.Error(c, http.StatusBadRequest, "Error binding request")
	}

	// Xác minh Firebase Token
	tokenData, err := h.firebaseAuth.VerifyFirebaseToken(c.Request().Context(), req.FirebaseToken)
	if err != nil {
		return response.Error(c, http.StatusUnauthorized, "Invalid or expired Firebase token")
	}

	// Lấy số điện thoại từ token
	phone, ok := tokenData.Claims["phone_number"].(string)
	if !ok {
		return response.Error(c, http.StatusUnauthorized, "Phone number not found in Firebase token")
	}

	// Tạo JWT
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &auth.JWTClaims{
		Phone: phone,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := jwtToken.SignedString([]byte(h.jwtSecret))
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, "Failed to generate JWT")
	}

	// Trả về JWT
	return response.OK(c, http.StatusOK, "Login successful", map[string]string{"token": tokenString})
}
