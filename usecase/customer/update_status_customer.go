package customer

import (
	"context"
	"express_be/repository/customer/entity"
	"express_be/usecase"
	"fmt"
)

// AdminAcceptCustomer implements AdminUsecase.
func (c *adminUsecaseImpl) AdminUpdateStatusCustomer(ctx context.Context, customerID *string, status *string) *usecase.Error {
	// Danh sách các trạng thái hợp lệ
	validStatuses := []string{
		string(entity.Pending),
		string(entity.Verified),
		string(entity.Blocked),
		string(entity.Active),
		string(entity.Inactive),
		string(entity.Suspended),
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

	err := c.customerRepo.UpdateStatus(ctx, customerID, &statusValue)
	if err != nil {
		return &usecase.Error{
			Code:    500,
			Message: err.Error(),
			Err:     err,
		}
	}
	return nil
}
