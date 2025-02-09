package res

import (
	"express_be/entity"
	"express_be/model/res"
)

func OrderDetailToRes(orderDetail *entity.OrderDetail) res.OrderRes {
	return res.OrderRes{
		ID:             orderDetail.ID,
		TrackingNumber: orderDetail.TrackingNumber,

		PickerID:       orderDetail.PickerID,
		PickerPhone:    orderDetail.PickerPhone,
		PickerFullName: orderDetail.PickerFullName,

		DelivererID:       orderDetail.DelivererID,
		DelivererPhone:    orderDetail.DelivererPhone,
		DelivererFullName: orderDetail.DelivererFullName,

		ShopName:              orderDetail.ShopName,
		SenderPhone:           orderDetail.SenderPhone,
		SenderSpecificAddress: orderDetail.SenderSpecificAddress,
		SenderWard:            orderDetail.SenderWard,
		SenderDistrict:        orderDetail.SenderDistrict,
		SenderCity:            orderDetail.SenderCity,
		AutonomousCode:        orderDetail.AutonomousCode,

		ReceiverPhone:           orderDetail.ReceiverPhone,
		ReceiverName:            orderDetail.ReceiverName,
		ReceiverSpecificAddress: orderDetail.ReceiverSpecificAddress,
		ReceiverWard:            orderDetail.ReceiverWard,
		ReceiverDistrict:        orderDetail.ReceiverDistrict,
		ReceiverCity:            orderDetail.ReceiverCity,
		Product:                 orderDetail.Product,
		Quantity:                orderDetail.Quantity,
		Weight:                  orderDetail.Weight,
		Dimensions:              orderDetail.Dimensions,
		DeclaredCod:             orderDetail.DeclaredCod,
		TotalCod:                orderDetail.TotalCod,
		Status:                  (*string)(&orderDetail.Status),
		SpecialFeatures:         orderDetail.SpecialFeatures,
		PickUpNotes:             orderDetail.PickUpNotes,
		DeliveryNotes:           orderDetail.DeliveryNotes,
	}
}
