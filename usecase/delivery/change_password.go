package delivery

import (
	"context"
	"express_be/core/security"
	"express_be/usecase"
)

// ChangePassword implements DeliveryPersonUsecase.
func (d *deliveryPersonUsecaseImpl) ChangePassword(ctx context.Context, phone *string, password *string) *usecase.Error {
	hashedPassword, err := security.HashPassword(*password)
	if err != nil {
		return &usecase.Error{
			Code:    400,
			Message: "Error hashing password: " + err.Error(),
			Err:     err,
		}
	}

	err = d.deliveryPersonRepo.ChangePassword(ctx, phone, &hashedPassword)
	if err != nil {
		return &usecase.Error{
			Code:    500,
			Message: "Error changing password: " + err.Error(),
			Err:     err,
		}
	}
	return nil
}
