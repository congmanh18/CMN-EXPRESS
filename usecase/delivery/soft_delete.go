package delivery

import (
	"context"
	"express_be/usecase"
)

// SoftDeleteDeliveryPerson implements DeliveryPersonUsecase.
func (d *deliveryPersonUsecaseImpl) SoftDeleteDeliveryPerson(ctx context.Context, id *string) *usecase.Error {
	err := d.deliveryPersonRepo.DeleteDeliveryPerson(ctx, id)
	if err != nil {
		return &usecase.Error{
			Code:    500,
			Message: "Failed to delete delivery person" + err.Error(),
			Err:     err,
		}
	}

	return nil
}
