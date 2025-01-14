package res

import (
	"express_be/core/record"
	customerEntity "express_be/repository/customer/entity"
	deliveryPersonEntity "express_be/repository/delivery/entity"
	userEntity "express_be/repository/user/entity"
)

func CustomerInfoToDetail(user *userEntity.User, customer *customerEntity.Customer) *userEntity.CustomerDetails {
	return &userEntity.CustomerDetails{
		User: userEntity.User{
			BaseEntity: record.BaseEntity{
				ID: user.ID,
			},
			Phone:                user.Phone,
			SpecificAddress:      user.SpecificAddress,
			Ward:                 user.Ward,
			District:             user.District,
			City:                 user.City,
			IdentificationNumber: user.IdentificationNumber,
			LastName:             user.LastName,
			MiddleName:           user.MiddleName,
			FirstName:            user.FirstName,
			DateOfBirth:          user.DateOfBirth,
			Gender:               user.Gender,
			Nationality:          user.Nationality,
			PlaceOfOrigin:        user.PlaceOfOrigin,
			PlaceOfResidence:     user.PlaceOfResidence,
			Role:                 user.Role,
			Status:               user.Status,
		},
		AccountType: (*string)(&customer.AccountType),
		Latitude:    customer.Latitude,
		Longtitude:  customer.Longtitude,
	}

}

func DeliveryPersonInfoToDetail(user *userEntity.User, deliveryPerson *deliveryPersonEntity.DeliveryPerson) *userEntity.DeliveryPersonDetails {
	return &userEntity.DeliveryPersonDetails{
		User: userEntity.User{
			BaseEntity: record.BaseEntity{
				ID: user.ID,
			},
			Phone:                user.Phone,
			SpecificAddress:      user.SpecificAddress,
			Ward:                 user.Ward,
			District:             user.District,
			City:                 user.City,
			IdentificationNumber: user.IdentificationNumber,
			LastName:             user.LastName,
			MiddleName:           user.MiddleName,
			FirstName:            user.FirstName,
			DateOfBirth:          user.DateOfBirth,
			Gender:               user.Gender,
			Nationality:          user.Nationality,
			PlaceOfOrigin:        user.PlaceOfOrigin,
			PlaceOfResidence:     user.PlaceOfResidence,
			Role:                 user.Role,
			Status:               user.Status,
		},
		SalaryRate: deliveryPerson.SalaryRate,
	}
}
