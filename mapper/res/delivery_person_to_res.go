package res

import (
	"express_be/entity"
	"express_be/model/res"
)

func DeliveryPersonToRes(user *entity.DeliveryPersonDetails) res.DeliveryPersonRes {
	status := string(user.Status)
	return res.DeliveryPersonRes{
		ID:              user.ID,
		Phone:           user.Phone,
		SpecificAddress: user.SpecificAddress,
		Ward:            user.Ward,
		District:        user.District,
		City:            user.City,
		Status:          &status,

		IdentificationNumber: user.IdentificationNumber,
		LastName:             user.LastName,
		FirstName:            user.FirstName,
		MiddleName:           user.MiddleName,
		DateOfBirth:          user.DateOfBirth,
		Gender:               user.Gender,
		Nationality:          user.Nationality,
		PlaceOfOrigin:        user.PlaceOfOrigin,
		PlaceOfResidence:     user.PlaceOfResidence,
		SalaryRate:           user.SalaryRate,
	}
}
