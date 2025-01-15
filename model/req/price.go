package req

type PriceReq struct {
	Region    *string  `json:"region"`
	Area      *string  `json:"area"`
	BasePrice *float64 `json:"base_price"`
}
