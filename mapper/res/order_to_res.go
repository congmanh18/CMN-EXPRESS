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

		SenderID:       orderDetail.SenderID,
		ShopName:       orderDetail.ShopName,
		SenderPhone:    orderDetail.SenderPhone,
		SenderAddress:  orderDetail.SenderAddress,
		AutonomousCode: orderDetail.AutonomousCode,

		ReceiverPhone:   orderDetail.ReceiverPhone,
		ReceiverName:    orderDetail.ReceiverName,
		ReceiverAddress: orderDetail.ReceiverAddress,
		Product:         orderDetail.Product,
		Quantity:        orderDetail.Quantity,
		Weight:          orderDetail.Weight,
		Dimensions:      orderDetail.Dimensions,
		DeclaredCod:     orderDetail.DeclaredCod,
		TotalCod:        orderDetail.TotalCod,
		Status:          (*string)(&orderDetail.Status),
		SpecialFeatures: orderDetail.SpecialFeatures,
		PickUpNotes:     orderDetail.PickUpNotes,
		DeliveryNotes:   orderDetail.DeliveryNotes,
	}
}

func OrderToRes(order *entity.Order) res.OrderRes {
	return res.OrderRes{
		ID:             order.ID,
		TrackingNumber: order.TrackingNumber,

		PickerID:    order.PickerID,
		DelivererID: order.DelivererID,

		SenderID:       order.SenderID,
		ShopName:       order.ShopName,
		SenderPhone:    order.SenderPhone,
		SenderAddress:  order.SenderAddress,
		AutonomousCode: order.AutonomousCode,

		ReceiverPhone:   order.ReceiverPhone,
		ReceiverName:    order.ReceiverName,
		ReceiverAddress: order.ReceiverAddress,
		Product:         order.Product,
		Quantity:        order.Quantity,
		Weight:          order.Weight,
		Dimensions:      order.Dimensions,
		DeclaredCod:     order.DeclaredCod,
		TotalCod:        order.TotalCod,
		Status:          (*string)(&order.Status),
		SpecialFeatures: order.SpecialFeatures,
		PickUpNotes:     order.PickUpNotes,
		DeliveryNotes:   order.DeliveryNotes,
	}
}
