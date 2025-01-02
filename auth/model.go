package auth

import (
	"github.com/golang-jwt/jwt/v4"
)

type JWTClaims struct {
	Phone string `json:"phone"`
	jwt.RegisteredClaims
}
