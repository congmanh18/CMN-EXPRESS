package user

import (
	"context"
	error "express_be/core/err"
	"express_be/repository/user/entity"
)

// AdminAcceptCustomer implements AdminUsecase.
func (c *userUsecaseImpl) UpdateStatus(ctx context.Context, customerID *string, approvalStatus *string) *error.Err {
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
	} else if approvalStatusValue == entity.Denied {
		accountStatus = entity.Blocked
	}

	if !isValid {
		return error.ErrInvalidFormat
	}

	err := c.userRepo.UpdateStatus(ctx, customerID, &approvalStatusValue, &accountStatus)
	if err != nil {
		return error.ErrInternalServer
	}
	return nil
}
