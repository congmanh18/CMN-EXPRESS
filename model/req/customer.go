package req

type UpdateCustomerReq struct {
	SpecificAddress      string  `json:"specific_address"`
	Ward                 string  `json:"ward"`
	District             string  `json:"district"`
	City                 string  `json:"city"`
	IdentificationNumber string  `json:"identification_number,omitempty"`
	FullName             string  `json:"full_name,omitempty"`
	DateOfBirth          string  `json:"date_of_birth,omitempty"`
	Gender               string  `json:"gender,omitempty"`
	Nationality          string  `json:"nationality,omitempty"`
	PlaceOfOrigin        string  `json:"place_of_origin"`
	PlaceOfResidence     string  `json:"place_of_residence"`
	CustomerAccountType  string  `json:"account_type,omitempty"`
	Latitude             float64 `json:"latitude,omitempty"`
	Longtitude           float64 `json:"longtitude,omitempty"`
}
