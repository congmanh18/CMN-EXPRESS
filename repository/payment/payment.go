package payment

type Payment struct{}

func (p *Payment) TableName() string {
	return "payments"
}
