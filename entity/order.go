package entity

import (
	"express_be/core/record"
)

type Status string

const (
	Created              Status = "created"               // Đơn hàng mới tạo
	PickingUp            Status = "picking_up"            // Đang lấy hàng
	PickedUp             Status = "picked_up"             // Đã lấy hàng
	PickupCanceled       Status = "pickup_canceled"       // Đơn hàng đã hủy lấy
	ReturnToHub          Status = "return_to_hub"         // Chuyển hàng về kho trung tâm
	AtHub                Status = "at_hub"                // Hàng đã đến kho
	OutForDelivery       Status = "out_for_delivery"      // Đang giao hàng
	Delivered            Status = "delivered"             // Giao hàng thành công
	Processing           Status = "processing"            // Đơn chờ xử lý
	ReattemptDelivery    Status = "reattempt_delivery"    // Đơn chờ phát lại
	Returning            Status = "returning"             // Đang chuyển hoàn
	Transferred          Status = "transferred"           // Đơn phát tiếp
	Refunded             Status = "refunded"              // Đơn đã trả
	Canceled             Status = "canceled"              // Đơn đã hủy giao
	CompensationChecking Status = "compensation_checking" // Đang xác minh bồi thường
	Compensated          Status = "compensated"           // Đã bồi thường
)

type Order struct {
	record.BaseEntity
	TrackingNumber         *string `json:"tracking_number"`
	WarehouseTransactionID *string `json:"warehouse_transaction_id"`

	PickerID    *string `gorm:"foreignKey:DeliveryPersonID" json:"picker_id"`
	DelivererID *string `gorm:"foreignKey:DeliveryPersonID" json:"deliverer_id"`

	SenderID              *string  `gorm:"foreignKey:CustomerID" json:"sender_id"`
	ShopName              *string  `json:"shop_name"`
	SenderPhone           *string  `json:"sender_phone"`
	SenderSpecificAddress *string  `json:"sender_specific_address"`
	SenderWard            *string  `json:"sender_ward"`
	SenderDistrict        *string  `json:"sender_district"`
	SenderCity            *string  `json:"sender_city"`
	SenderLatitude        *float64 `json:"sender_latitude"`
	SenderLongitude       *float64 `json:"sender_longitude"`
	SenderGeohash         *string  `json:"sender_geohash"`
	AutonomousCode        *string  `json:"autonomous_code"`

	ReceiverName            *string  `json:"receiver_name"`
	ReceiverPhone           *string  `json:"receiver_phone"`
	ReceiverSpecificAddress *string  `json:"receiver_specific_address"`
	ReceiverWard            *string  `json:"receiver_ward"`
	ReceiverDistrict        *string  `json:"receiver_district"`
	ReceiverCity            *string  `json:"receiver_city"`
	ReceiverLatitude        *float64 `json:"receiver_latitude"`
	ReceiverLongitude       *float64 `json:"receiver_longitude"`
	ReceiverGeohash         *string  `json:"receiver_geohash"`

	Product         *string  `json:"product"`
	Quantity        *int     `json:"quantity"`
	Weight          *float64 `json:"weight"`
	Dimensions      *string  `json:"dimensions"`
	DeclaredCod     *float64 `json:"declared_cod"`
	TotalCod        *float64 `json:"total_cod"`
	Status          Status   `json:"status"`
	SpecialFeatures *string  `json:"special_features"`
	PickUpNotes     *string  `json:"pick_up_notes"`
	DeliveryNotes   *string  `json:"delivery_notes"`

	ShippingFee *float64 `json:"shipping_fee"`
	TotalAmount *float64 `json:"total_amount"`
}

// OrderStatus []OrderStatus `gorm:"foreignKey:OrderID" json:"order_status"`
func (o *Order) TableName() string {
	return "orders"
}
