package res

import (
	"express_be/entity"
	"express_be/model/res"
)

func CustomerToRes(user *entity.CustomerDetails) res.CustomerRes {
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
		LastName:             user.LastName,
		MiddleName:           user.MiddleName,
		FirstName:            user.FirstName,
		DateOfBirth:          user.DateOfBirth,
		Gender:               user.Gender,
		Nationality:          user.Nationality,
		PlaceOfOrigin:        user.PlaceOfOrigin,
		PlaceOfResidence:     user.PlaceOfResidence,
		AccountType:          user.AccountType,
		Latitude:             user.Latitude,
		Longitude:            user.Longitude,
	}
}
