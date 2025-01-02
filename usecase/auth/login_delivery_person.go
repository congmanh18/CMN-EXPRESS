package auth

import (
	"context"
	"express_be/core/security"
	"express_be/usecase"
)

func (a *authUsecaseImpl) LoginDeliveryPerson(ctx context.Context, phone, password *string) (*string, *usecase.Error) {
	deliveryPerson, err := a.customerRepo.FindByPhone(ctx, phone)
	if err != nil {
		return nil, &usecase.Error{
			Code:    401,
			Message: "Invalid phone number or password",
			Err:     err,
		}
	}

	if !security.VerifyPassword(*password, *deliveryPerson.PasswordHash) {
		return nil, &usecase.Error{
			Code:    401,
			Message: "Invalid phone number or password",
			Err:     err,
		}
	}

	return deliveryPerson.ID, nil
}
