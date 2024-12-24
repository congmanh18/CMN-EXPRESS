package auth

import (
	"context"
	"fmt"

	"firebase.google.com/go/auth"
)

type AuthService interface {
	VerifyToken(ctx context.Context, token string) (*auth.Token, error)
	CheckAccountType(token *auth.Token, requiredType string) error
}

type firebaseAuthService struct {
	client *auth.Client
}

func NewFirebaseAuthService(client *auth.Client) AuthService {
	return &firebaseAuthService{client: client}
}

func (s *firebaseAuthService) VerifyToken(ctx context.Context, token string) (*auth.Token, error) {
	tokenData, err := s.client.VerifyIDToken(ctx, token)
	if err != nil {
		return nil, err
	}
	return tokenData, nil
}

func (s *firebaseAuthService) CheckAccountType(token *auth.Token, requiredType string) error {
	// Kiểm tra claims
	claims := token.Claims
	if claims == nil {
		return fmt.Errorf("unauthorized: missing claims")
	}

	// Kiểm tra account type
	accountType, ok := claims["account_type"].(string)
	if !ok || accountType != requiredType {
		return fmt.Errorf("unauthorized: required account type %s, but got %s", requiredType, accountType)
	}

	return nil
}
