package res

type DeliveryPersonRes struct {
	ID                   *string  `json:"id"`
	Phone                *string  `json:"phone"`
	SpecificAddress      *string  `json:"specific_address"`
	Ward                 *string  `json:"ward"`
	District             *string  `json:"district"`
	City                 *string  `json:"city"`
	SalaryRate           *float64 `json:"salary_rate"`
	Status               *string  `json:"status"`
	IdentificationNumber *string  `json:"identification_number"`
	FullName             *string  `json:"full_name"`
	DateOfBirth          *string  `json:"date_of_birth"`
	Gender               *string  `json:"gender"`
	Nationality          *string  `json:"nationality"`
	PlaceOfOrigin        *string  `json:"place_of_origin"`
	PlaceOfResidence     *string  `json:"place_of_residence"`
}

type DeliveryPaginationResponse struct {
	Page       int                 `json:"page"`
	PageSize   int                 `json:"page_size"`
	TotalPages int                 `json:"total_pages,omitempty"`
	TotalCount int                 `json:"total_count,omitempty"`
	List       []DeliveryPersonRes `json:"list"`
}
