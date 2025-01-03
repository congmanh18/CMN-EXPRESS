package res

import (
	"express_be/model/res"
	"express_be/repository/delivery/entity"
)

func DeliveryPersonToRes(deliveryPerson *entity.DeliveryPerson) res.DeliveryPersonRes {
	status := string(deliveryPerson.Status)
	return res.DeliveryPersonRes{
		ID:             deliveryPerson.ID,
		Phone:          deliveryPerson.Phone,
		CurrentAddress: deliveryPerson.CurrentAddress,
		SalaryRate:     deliveryPerson.SalaryRate,

		IdentificationNumber: deliveryPerson.IdentificationNumber,
		FullName:             deliveryPerson.FullName,
		DateOfBirth:          deliveryPerson.DateOfBirth,
		Gender:               deliveryPerson.Gender,
		Nationality:          deliveryPerson.Nationality,
		PlaceOfOrigin:        deliveryPerson.PlaceOfOrigin,
		PlaceOfResidence:     deliveryPerson.PlaceOfResidence,
		Status:               &status,
	}
}
