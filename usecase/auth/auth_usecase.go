package auth

import (
	"context"
	"express_be/repository/customer"
	customerEntity "express_be/repository/customer/entity"
	"express_be/repository/delivery"
	deliveryPersonEntity "express_be/repository/delivery/entity"
	"express_be/usecase"
)

type AuthUsecase interface {
	CreateCustomerUsecase(ctx context.Context, customer *customerEntity.Customer) *usecase.Error
	CreateDeliveryPersonUsecase(ctx context.Context, deliveryPerson *deliveryPersonEntity.DeliveryPerson) *usecase.Error
	VerifyCustomerUsecase(ctx context.Context, uid *string) *usecase.Error
	VerifyDeliveryPersonUsecase(ctx context.Context, uid *string) *usecase.Error
	LoginCustomer(ctx context.Context, phone, password *string) (*string, *usecase.Error)
	LoginDeliveryPerson(ctx context.Context, phone, password *string) (*string, *usecase.Error)
}

type authUsecaseImpl struct {
	customerRepo       customer.Repo
	deliveryPersonRepo delivery.Repo
}

func NewAuthUsecase(customerRepo customer.Repo, deliveryPersonRepo delivery.Repo) AuthUsecase {
	return &authUsecaseImpl{
		customerRepo:       customerRepo,
		deliveryPersonRepo: deliveryPersonRepo,
	}
}
