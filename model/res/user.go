package res

type CustomerRes struct {
	ID              *string `json:"id"`
	Phone           *string `json:"phone"`
	SpecificAddress *string `json:"specific_address"`
	Ward            *string `json:"ward"`
	District        *string `json:"district"`
	City            *string `json:"city"`
	AccountType     *string `json:"account_type"`

	Status    *string  `json:"status"`
	Latitude  *float64 `json:"latitude"`
	Longitude *float64 `json:"longitude"`

	IdentificationNumber *string `json:"identification_number"`
	LastName             *string `json:"last_name"`
	FirstName            *string `json:"first_name"`
	MiddleName           *string `json:"middle_name"`
	DateOfBirth          *string `json:"date_of_birth"`
	Gender               *string `json:"gender"`
	Nationality          *string `json:"nationality"`
	PlaceOfOrigin        *string `json:"place_of_origin"`
	PlaceOfResidence     *string `json:"place_of_residence"`
}

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
	LastName             *string  `json:"last_name"`
	FirstName            *string  `json:"first_name"`
	MiddleName           *string  `json:"middle_name"`
	DateOfBirth          *string  `json:"date_of_birth"`
	Gender               *string  `json:"gender"`
	Nationality          *string  `json:"nationality"`
	PlaceOfOrigin        *string  `json:"place_of_origin"`
	PlaceOfResidence     *string  `json:"place_of_residence"`
}

type PaginationResponse struct {
	Page            int                 `json:"page"`
	PageSize        int                 `json:"page_size"`
	TotalPages      int                 `json:"total_pages,omitempty"` // Nếu cần tổng số trang
	TotalCount      int64               `json:"total_count,omitempty"` // Nếu cần tổng số bản ghi
	Customers       []CustomerRes       `json:"customers,omitempty"`
	DeliveryPersons []DeliveryPersonRes `json:"delivery_persons,omitempty"`
	Orders          []OrderRes          `json:"orders,omitempty"`
}
