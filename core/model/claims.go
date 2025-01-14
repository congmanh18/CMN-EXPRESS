package token

import (
	"github.com/golang-jwt/jwt/v5"
)

type JwtCustomClaims struct {
	ID   string `json:"id"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}
