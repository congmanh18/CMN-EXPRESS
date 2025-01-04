package firebase

import (
	"context"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

type FirebaseAuthService struct {
	client *auth.Client
}

func NewFirebaseAuthService(credPath string) (*FirebaseAuthService, error) {
	opt := option.WithCredentialsFile(credPath)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, err
	}

	client, err := app.Auth(context.Background())
	if err != nil {
		return nil, err
	}

	return &FirebaseAuthService{client: client}, nil
}

func (s *FirebaseAuthService) VerifyFirebaseToken(ctx context.Context, token string) (*auth.Token, error) {
	return s.client.VerifyIDToken(ctx, token)
}
