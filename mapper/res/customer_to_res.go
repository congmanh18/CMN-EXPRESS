package res

import (
	"express_be/model/res"
	"express_be/repository/customer/entity"
)

func CustomerToRes(customer *entity.Customer) res.CustomerRes {
	status := string(customer.Status)
	return res.CustomerRes{
		ID:             customer.ID,
		Phone:          customer.Phone,
		CurrentAddress: customer.CurrentAddress,
		AccountType:    (*string)(&customer.AccountType),

		Status:    &status,
		Latitude:  customer.Latitude,
		Longitude: customer.Longtitude,

		IdentificationNumber: customer.IdentificationNumber,
		FullName:             customer.FullName,
		DateOfBirth:          customer.DateOfBirth,
		Gender:               customer.Gender,
		Nationality:          customer.Nationality,
		PlaceOfOrigin:        customer.PlaceOfOrigin,
		PlaceOfResidence:     customer.PlaceOfResidence,
	}
}
