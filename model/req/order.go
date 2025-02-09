package req

type CreateOrderReq struct {
	ShopName              *string  `json:"shop_name" validate:"required"`
	SenderPhone           *string  `json:"sender_phone" validate:"required"`
	SenderSpecificAddress *string  `json:"sender_specific_address"`
	SenderWard            *string  `json:"sender_ward"`
	SenderDistrict        *string  `json:"sender_district"`
	SenderCity            *string  `json:"sender_city"`
	SenderLatitude        *float64 `json:"sender_latitude" validate:"required"`
	SenderLongitude       *float64 `json:"sender_longitude" validate:"required"`

	ReceiverName            *string  `json:"receiver_name" validate:"required"`
	ReceiverPhone           *string  `json:"receiver_phone" validate:"required"`
	ReceiverSpecificAddress *string  `json:"receiver_specific_address"`
	ReceiverWard            *string  `json:"receiver_ward"`
	ReceiverDistrict        *string  `json:"receiver_district"`
	ReceiverCity            *string  `json:"receiver_city"`
	ReceiverLatitude        *float64 `json:"receiver_latitude" validate:"required"`
	ReceiverLongitude       *float64 `json:"receiver_longitude" validate:"required"`

	Product      *string  `json:"product" validate:"required"`
	Quantity     *int     `json:"quantity" validate:"required"`
	Weight       *float64 `json:"weight" validate:"required"`
	Dimensions   *string  `json:"dimensions" validate:"required"`
	DeclaredCode *float64 `json:"declared_code" validate:"required"`

	PickUpNotes   *string `json:"pick_up_notes" validate:"required"`
	DeliveryNotes *string `json:"delivery_notes" validate:"required"`
}

type LocationReq struct {
	Latitude  float64 `json:"latitude" validate:"required"`
	Longitude float64 `json:"longitude" validate:"required"`
}
