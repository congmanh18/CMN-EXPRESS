package res

import (
	"express_be/model/res"
	userEntity "express_be/repository/user/entity"
)

func CustomerToRes(user *userEntity.CustomerDetails) res.CustomerRes {
	status := string(user.Status)
	return res.CustomerRes{
		ID:              user.ID,
		Phone:           user.Phone,
		SpecificAddress: user.SpecificAddress,
		Ward:            user.Ward,
		District:        user.District,
		City:            user.City,
		Status:          &status,

		IdentificationNumber: user.IdentificationNumber,
		FirstName:            user.FirstName,
		DateOfBirth:          user.DateOfBirth,
		Gender:               user.Gender,
		Nationality:          user.Nationality,
		PlaceOfOrigin:        user.PlaceOfOrigin,
		PlaceOfResidence:     user.PlaceOfResidence,
		AccountType:          user.AccountType,
		Latitude:             user.Latitude,
		Longitude:            user.Longtitude,
	}
}
