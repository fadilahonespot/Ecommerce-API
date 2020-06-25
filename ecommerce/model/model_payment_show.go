package model

type PaymentShow struct {
	PaymentMethodID uint             `json:"payment_method_id"`
	PaymentType     string           `json:"payment_type"`
	Payments        []SubPaymentShow `json:"payments"`
}

type SubPaymentShow struct {
	PaymentID       uint   `json:"payment_id"`
	RekeningAccount string `gorm:"not null" json:"rek_account"`
	RekeningName    string `gorm:"not null" json:"rek_name"`
	RekeningNumber  string `gorm:"not null" json:"rek_number"`
}
