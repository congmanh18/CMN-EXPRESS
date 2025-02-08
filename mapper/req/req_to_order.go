package req

import (
	"express_be/core/pointer"
	"express_be/core/record"
	"express_be/core/utils/gen"
	"express_be/entity"
	model "express_be/model/req"

	"github.com/google/uuid"
	"github.com/mmcloughlin/geohash"
)

func CreateOrderReqToOrder(req model.CreateOrderReq, senderID *string) *entity.Order {
	senderGeohash := geohash.Encode(*req.SenderLatitude, *req.SenderLongitude)
	receiverGeohash := geohash.Encode(*req.ReceiverLatitude, *req.ReceiverLongitude)

	id, _ := gen.GenerateShipmentCode()
	return &entity.Order{
		BaseEntity: record.BaseEntity{
			ID: pointer.String(uuid.New().String()),
		},
		TrackingNumber:        pointer.String(id),
		SenderID:              senderID,
		ShopName:              req.ShopName,
		SenderPhone:           req.SenderPhone,
		SenderSpecificAddress: req.SenderSpecificAddress,
		SenderWard:            req.SenderWard,
		SenderDistrict:        req.SenderDistrict,
		SenderCity:            req.SenderCity,
		SenderLatitude:        req.SenderLatitude,
		SenderLongitude:       req.SenderLongitude,
		SenderGeohash:         &senderGeohash,

		ReceiverName:            req.ReceiverName,
		ReceiverPhone:           req.ReceiverPhone,
		ReceiverSpecificAddress: req.ReceiverSpecificAddress,
		ReceiverWard:            req.ReceiverWard,
		ReceiverDistrict:        req.ReceiverDistrict,
		ReceiverCity:            req.ReceiverCity,
		ReceiverLatitude:        req.ReceiverLatitude,
		ReceiverLongitude:       req.ReceiverLongitude,
		ReceiverGeohash:         &receiverGeohash,

		Product:     req.Product,
		Quantity:    req.Quantity,
		Weight:      req.Weight,
		Dimensions:  req.Dimensions,
		DeclaredCod: req.DeclaredCode,

		PickUpNotes:   req.PickUpNotes,
		DeliveryNotes: req.DeliveryNotes,
		Status:        entity.Created,
	}
}
