package firebase

import (
	"context"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

// "./conf/express-2227f-firebase-adminsdk-vbu9s-83580ceea5.json"
func InitFirebase(pathsdk string) (*auth.Client, error) {
	ctx := context.Background()
	opt := option.WithCredentialsFile(pathsdk)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, err
	}

	client, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}

	return client, nil
}
