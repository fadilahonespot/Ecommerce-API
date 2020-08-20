package model

import (
	"github.com/jinzhu/gorm"
)

type Transaction struct {
	gorm.Model
	UserID      uint `gorm:"not null" json:"user_id"`
	TotalPrince int  `gorm:"type:bigint; not null" json:"total_prince"`
}

type MitraTransaction struct {
	gorm.Model
	MitraTransactionNumber int    `gorm:"type:bigint; not null" json:"mitra_transaction_number"`
	MitraID                uint   `gorm:"not null" json:"mitra_id"`
	TransactionID          uint   `gorm:"not null" json:"transaction_id"`
	SubTotalMitra          int    `gorm:"type:bigint; not null" json:"sub_total"`
	Status                 string `gorm:"type:enum('waiting payment', 'waiting process', 'reject', 'process', 'send', 'received', 'confirmation')"`
}

type SubTransaction struct {
	gorm.Model
	MitraTransactionID uint   `gorm:"not null" json:"mitra_transaction_id"`
	ProductID          uint   `gorm:"not null" json:"product_id"`
	Note               string `json:"note"`
	Quantity           int    `gorm:"not null" json:"quantity"`
	SubTotal           int    `gorm:"not null" json:"sub_total"`
}

func (e *Transaction) TableName() string {
	return "transaction"
}

func (e *MitraTransaction) TableName() string {
	return "mitra_transaction"
}

func (e *SubTransaction) TableName() string {
	return "subtransaction"
}
