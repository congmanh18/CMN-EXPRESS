package res

import (
	"express_be/model/res"
	userEntity "express_be/repository/user/entity"
)

func CustomerToRes(user *userEntity.CustomerDetails) res.CustomerRes {
	status := string(user.Status)
	return res.CustomerRes{
		ID:             user.ID,
		Phone:          user.Phone,
		CurrentAddress: user.CurrentAddress,
		Status:         &status,

		IdentificationNumber: user.IdentificationNumber,
		FullName:             user.FullName,
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
