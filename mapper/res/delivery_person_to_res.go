package res

import (
	"express_be/model/res"
	userEntity "express_be/repository/user/entity"
)

func DeliveryPersonToRes(user *userEntity.DeliveryPersonDetails) res.DeliveryPersonRes {
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
