package delivery

import (
	"context"
	"express_be/repository/delivery/entity"
	"express_be/usecase"
	"fmt"
)

func (c *adminUseCaseImpl) AdminUpdateStatusDeliveryPerson(ctx context.Context, deliveryPersonID *string, status *string) *usecase.Error {
	validStatuses := []string{
		string(entity.Pending),
		string(entity.Verified),
		string(entity.Blocked),
		string(entity.Active),
		string(entity.Inactive),
		string(entity.Suspended),
		string(entity.OnDuty),
		string(entity.OffDuty),
	}

	statusValue := entity.Status(*status)
	// Kiểm tra trạng thái hợp lệ
	isValid := false
	for _, validStatus := range validStatuses {
		if *status == validStatus {
			isValid = true
			break
		}
	}

	if !isValid {
		return &usecase.Error{
			Code:    400,
			Message: "Invalid status value",
			Err:     fmt.Errorf("status '%s' is not a valid value", *status),
		}
	}

	err := c.deliveryPersonRepo.UpdateStatus(ctx, deliveryPersonID, &statusValue)
	if err != nil {
		return &usecase.Error{
			Code:    500,
			Message: err.Error(),
			Err:     err,
		}
	}
	return nil
}
