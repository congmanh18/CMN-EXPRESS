package req

import (
	model "express_be/model/req"
	"express_be/repository/order/entity"

	"github.com/mmcloughlin/geohash"
)

func CreateOrderReqToOrder(req model.CreateOrderReq) *entity.Order {
	senderGeohash := geohash.Encode(*req.SenderLatitude, *req.SenderLongitude)
	receiverGeohash := geohash.Encode(*req.ReceiverLatitude, *req.ReceiverLongitude)

	return &entity.Order{
		SenderID:        req.SenderID,
		ShopName:        req.ShopName,
		SenderPhone:     req.SenderPhone,
		SenderAddress:   req.SenderAddress,
		SenderLatitude:  req.SenderLatitude,
		SenderLongitude: req.SenderLongitude,
		SenderGeohash:   &senderGeohash,

		ReceiverName:      req.ReceiverName,
		ReceiverPhone:     req.ReceiverPhone,
		ReceiverAddress:   req.ReceiverAddress,
		ReceiverLatitude:  req.ReceiverLatitude,
		ReceiverLongitude: req.ReceiverLongitude,
		ReceiverGeohash:   &receiverGeohash,

		Product:     req.Product,
		Quantity:    req.Quantity,
		Weight:      req.Weight,
		Dimensions:  req.Dimensions,
		DeclaredCod: req.DeclaredCode,
	}
}
