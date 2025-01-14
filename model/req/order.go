package req

type CreateOrderReq struct {
	SenderID        *string  `json:"sender_id" validate:"required"`
	ShopName        *string  `json:"shop_name" validate:"required"`
	SenderPhone     *string  `json:"sender_phone" validate:"required"`
	SenderAddress   *string  `json:"sender_address" validate:"required"`
	SenderLatitude  *float64 `json:"sender_latitude" validate:"required"`
	SenderLongitude *float64 `json:"sender_longitude" validate:"required"`

	ReceiverName      *string  `json:"receiver_name" validate:"required"`
	ReceiverPhone     *string  `json:"receiver_phone" validate:"required"`
	ReceiverAddress   *string  `json:"receiver_address" validate:"required"`
	ReceiverLatitude  *float64 `json:"receiver_latitude" validate:"required"`
	ReceiverLongitude *float64 `json:"receiver_longitude" validate:"required"`

	Product      *string  `json:"product" validate:"required"`
	Quantity     *int     `json:"quantity" validate:"required"`
	Weight       *float64 `json:"weight" validate:"required"`
	Dimensions   *string  `json:"dimensions" validate:"required"`
	DeclaredCode *float64 `json:"declared_code" validate:"required"`
}
