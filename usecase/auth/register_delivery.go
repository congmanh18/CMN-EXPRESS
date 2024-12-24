package auth

import (
	"context"
	"errors"
	"express_be/repository/delivery"
	deliveryPersonEntity "express_be/repository/delivery/entity"

	"express_be/usecase"
)

type RegisterDeliveryPersonUseCase interface {
	CreateDeliveryPersonUsecase(ctx context.Context, deliveryPerson *deliveryPersonEntity.DeliveryPerson) *usecase.Error
}

type RegisterDeliveryPersonUseCaseImpl struct {
	DeliveryPersonRepo delivery.Repo
}

func (r *RegisterDeliveryPersonUseCaseImpl) CreateDeliveryPersonUsecase(ctx context.Context, deliveryPerson *deliveryPersonEntity.DeliveryPerson) *usecase.Error {
	// Kiểm tra id đã tồn tại chưa
	exists, _ := r.DeliveryPersonRepo.FindByPhone(ctx, deliveryPerson.ID)
	if exists != nil {
		return &usecase.Error{
			Code:    400,
			Message: "Phone already registered as delivery person",
			Err:     errors.New("delivery person already exists"),
		}
	}
	err := r.DeliveryPersonRepo.CreateDeliveryPerson(ctx, deliveryPerson)
	if err != nil {
		return &usecase.Error{
			Code:    500,
			Message: "Failed to create delivery person",
			Err:     err,
		}
	}

	return nil
}

func NewRegisterDeliveryPersonUseCase(deliveryPersonRepo delivery.Repo) RegisterDeliveryPersonUseCase {
	return &RegisterDeliveryPersonUseCaseImpl{
		DeliveryPersonRepo: deliveryPersonRepo,
	}
}
