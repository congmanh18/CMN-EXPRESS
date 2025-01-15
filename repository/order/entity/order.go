package entity

import (
	"express_be/core/record"
)

type Status string

const (
	Created          Status = "created"
	PickingUp        Status = "picking_up"
	PickedUp         Status = "picked_up"
	PickupCanceled   Status = "pickup_canceled"
	ReturnToHub      Status = "return_to_hub"
	AtHub            Status = "at_hub"
	OutForDelivery   Status = "out_for_delivery"
	Delivered        Status = "delivered"
	Returned         Status = "returned"
	DeliveryCanceled Status = "delivery_canceled"
)

type Order struct {
	record.BaseEntity

	PickerID               *string `gorm:"foreignKey:DeliveryPersonID" json:"picker_id"`
	DelivererID            *string `gorm:"foreignKey:DeliveryPersonID" json:"deliverer_id"`
	WarehouseTransactionID *string `json:"warehouse_transaction_id"`

	SenderID        *string  `gorm:"foreignKey:CustomerID" json:"sender_id"`
	ShopName        *string  `json:"shop_name"`
	SenderPhone     *string  `json:"sender_phone"`
	SenderAddress   *string  `json:"sender_address"`
	SenderLatitude  *float64 `json:"sender_latitude"`
	SenderLongitude *float64 `json:"sender_longitude"`
	SenderGeohash   *string  `json:"sender_geohash"`

	ReceiverName      *string  `json:"receiver_name"`
	ReceiverPhone     *string  `json:"receiver_phone"`
	ReceiverAddress   *string  `json:"receiver_address"`
	ReceiverLatitude  *float64 `json:"receiver_latitude"`
	ReceiverLongitude *float64 `json:"receiver_longitude"`
	ReceiverGeohash   *string  `json:"receiver_geohash"`

	Product     *string  `json:"product"`
	Quantity    *int     `json:"quantity"`
	Weight      *float64 `json:"weight"`
	Dimensions  *string  `json:"dimensions"`
	DeclaredCod *float64 `json:"declared_cod"`
	TotalCod    *float64 `json:"total_cod"`
	Status      Status   `json:"status"`

	ShippingFee *float64 `json:"shipping_fee"`
	TotalAmount *float64 `json:"total_amount"`

	// Quan hệ với bảng con OrderStatus
	OrderStatus []OrderStatus `gorm:"foreignKey:OrderID" json:"order_status"`
}

func (o *Order) TableName() string {
	return "orders"
}
