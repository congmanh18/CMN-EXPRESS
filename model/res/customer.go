package res

type CustomerRes struct {
	ID             *string `json:"id"`
	Phone          *string `json:"phone"`
	CurrentAddress *string `json:"current_address"`
	AccountType    *string `json:"account_type"`

	Status    *string  `json:"status"`
	Latitude  *float64 `json:"latitude"`
	Longitude *float64 `json:"longitude"`

	IdentificationNumber *string `json:"identification_number"`
	FullName             *string `json:"full_name"`
	DateOfBirth          *string `json:"date_of_birth"`
	Gender               *string `json:"gender"`
	Nationality          *string `json:"nationality"`
	PlaceOfOrigin        *string `json:"place_of_origin"`
	PlaceOfResidence     *string `json:"place_of_residence"`
}

type CustomerPaginationResponse struct {
	Page       int           `json:"page"`
	PageSize   int           `json:"page_size"`
	TotalPages int           `json:"total_pages,omitempty"` // Nếu cần tổng số trang
	TotalCount int64         `json:"total_count,omitempty"` // Nếu cần tổng số bản ghi
	List       []CustomerRes `json:"list"`
}
