package model

import (
	"github.com/jinzhu/gorm"
)

type PaymentMethod struct {
	gorm.Model
	PaymentType string `gorm:"not null" json:"payment_type"`
}

type Payment struct {
	gorm.Model
	PaymentMethodID uint   `gorm:"not null" json:"payment_method_id"`
	RekeningAccount string `gorm:"not null" json:"rek_account"`
	RekeningName    string `gorm:"not null" json:"rek_name"`
	RekeningNumber  string `gorm:"not null" json:"rek_number"`
}

type PaymentDetail struct {
	gorm.Model
	PaymentID              uint   `gorm:"not null" json:"payment_id"`
	TransactionID          uint   `gorm:"not null" json:"transaction_id"`
	TotalPayment           int    `gorm:"type:bigint; not null" json:"total_payment"`
	RekeningSendAccound    string `json:"rek_send_acccount"`
	RekeningNumSendAccound string `json:"rek_num_send_account"`
	TotalPaymentReceiped   int    `gorm:"type:bigint" json:"total_payment_receiped"`
	Status                 string `gorm:"type:enum('waiting', 'failed', 'receiped')" json:"status"`
}

func (e *Payment) TableName() string {
	return "payment"
}

func (e *PaymentMethod) TableName() string {
	return "payment_method"
}

func (e *PaymentDetail) TableName() string {
	return "payment_detail"
}
