package entity

import (
	"express_be/core/record"
)

type Order struct {
	record.BaseEntity

	DeliveryPersonID       *string `json:"delivery_person_id"`
	WarehouseTransactionID *string `json:"warehouse_transaction_id"`

	SenderID        *uint    `gorm:"foreignKey:SenderID" json:"sender_id"`
	SenderName      *string  `json:"sender_name"`
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
	Status      *string  `json:"status"`

	AdditionalServices *string  `json:"additional_services"`
	ShippingFee        *float64 `json:"shipping_fee"`
	TotalAmount        *float64 `json:"total_amount"`

	// Quan hệ với bảng con OrderStatus
	OrderStatus       []OrderStatus       `gorm:"foreignKey:OrderID" json:"order_status"`
	AdditionalService []AdditionalService `gorm:"foreignKey:OrderID" json:"additional_service"`
	CODTransaction    []CODTransaction    `gorm:"foreignKey:OrderID" json:"cod_transaction"`
}

func (o *Order) TableName() string {
	return "orders"
}
