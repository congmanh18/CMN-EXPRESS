package res

import (
	"express_be/core/record"
	"express_be/entity"
)

func CustomerInfoToDetail(user *entity.User, customer *entity.Customer) *entity.CustomerDetails {
	return &entity.CustomerDetails{
		User: entity.User{
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
		Longitude:   customer.Longitude,
	}

}

func DeliveryPersonInfoToDetail(user *entity.User, deliveryPerson *entity.DeliveryPerson) *entity.DeliveryPersonDetails {
	return &entity.DeliveryPersonDetails{
		User: entity.User{
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
