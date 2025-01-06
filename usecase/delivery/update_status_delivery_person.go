package delivery

import (
	"context"
	"express_be/repository/delivery/entity"
	"express_be/usecase"
	"fmt"
)

func (c *adminUseCaseImpl) AdminUpdateStatusDeliveryPerson(ctx context.Context, deliveryPersonID *string, approvalStatus *string) *usecase.Error {
	validStatuses := []string{
		string(entity.Accepted),
		string(entity.Denied),
	}

	approvalStatusValue := entity.ApprovalStatus(*approvalStatus)
	// Kiểm tra trạng thái hợp lệ
	isValid := false
	for _, validStatus := range validStatuses {
		if *approvalStatus == validStatus {
			isValid = true
			break
		}
	}

	var accountStatus entity.Status
	if approvalStatusValue == entity.Accepted {
		accountStatus = entity.Verified
	} else {
		accountStatus = entity.Blocked
	}

	if !isValid {
		return &usecase.Error{
			Code:    400,
			Message: "Invalid approval status value",
			Err:     fmt.Errorf("approval status '%s' is not a valid value", *approvalStatus),
		}
	}

	err := c.deliveryPersonRepo.UpdateStatus(ctx, deliveryPersonID, &approvalStatusValue, &accountStatus)
	if err != nil {
		return &usecase.Error{
			Code:    500,
			Message: err.Error(),
			Err:     err,
		}
	}
	return nil
}
