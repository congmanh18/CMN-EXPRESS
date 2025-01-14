package entity

type Receiver struct {
}

func (r *Receiver) TableName() string {
	return "receivers"
}
