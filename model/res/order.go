package res

type OrderRes struct {
	ID             *string `json:"id"`
	TrackingNumber *string `json:"tracking_number"`

	PickerID       *string `json:"picker_id"`
	PickerPhone    *string `json:"picker_phone"`
	PickerFullName *string `json:"picker_full_name"`

	DelivererID       *string `json:"deliverer_id"`
	DelivererPhone    *string `json:"deliverer_phone"`
	DelivererFullName *string `json:"deliverer_full_name"`

	SenderID              *string `json:"sender_id"`
	ShopName              *string `json:"shop_name"`
	SenderPhone           *string `json:"sender_phone"`
	SenderSpecificAddress *string `json:"sender_specific_address"`
	SenderWard            *string `json:"sender_ward"`
	SenderDistrict        *string `json:"sender_district"`
	SenderCity            *string `json:"sender_city"`
	AutonomousCode        *string `json:"autonomous_code"`

	ReceiverName            *string `json:"receiver_name"`
	ReceiverPhone           *string `json:"receiver_phone"`
	ReceiverSpecificAddress *string `json:"receiver_specific_address"`
	ReceiverWard            *string `json:"receiver_ward"`
	ReceiverDistrict        *string `json:"receiver_district"`
	ReceiverCity            *string `json:"receiver_city"`

	Product         *string  `json:"product"`
	Quantity        *int     `json:"quantity"`
	Weight          *float64 `json:"weight"`
	Dimensions      *string  `json:"dimensions"`
	DeclaredCod     *float64 `json:"declared_cod"`
	TotalCod        *float64 `json:"total_cod"`
	Status          *string  `json:"status"`
	SpecialFeatures *string  `json:"special_features"`
	PickUpNotes     *string  `json:"pick_up_notes"`
	DeliveryNotes   *string  `json:"delivery_notes"`

	ShippingFee *float64 `json:"shipping_fee"`
	TotalAmount *float64 `json:"total_amount"`
}
