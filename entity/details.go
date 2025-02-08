package entity

type OrderDetail struct {
	Order
	PickerPhone       *string `json:"picker_phone"`
	PickerFullName    *string `json:"picker_full_name"`
	DelivererPhone    *string `json:"deliverer_phone"`
	DelivererFullName *string `json:"deliverer_full_name"`
}
