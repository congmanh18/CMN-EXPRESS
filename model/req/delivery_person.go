package req

type UpdateDeliveryPersonReq struct {
	CurrentAddress       string `json:"current_address,omitempty"`
	IdentificationNumber string `json:"identification_number,omitempty"`
	FullName             string `json:"full_name,omitempty"`
	DateOfBirth          string `json:"date_of_birth,omitempty"`
	Gender               string `json:"gender,omitempty"`
	Nationality          string `json:"nationality,omitempty"`
	PlaceOfOrigin        string `json:"place_of_origin"`
	PlaceOfResidence     string `json:"place_of_residence"`
}
