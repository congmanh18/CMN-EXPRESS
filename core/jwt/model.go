package jwt

import "github.com/golang-jwt/jwt/v4"

type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type JwtCustomClaims struct {
	ID                   string `json:"id"`   // User ID
	Role                 string `json:"role"` // User Role
	jwt.RegisteredClaims        // Thông tin chuẩn như `exp`, `iat`
}
