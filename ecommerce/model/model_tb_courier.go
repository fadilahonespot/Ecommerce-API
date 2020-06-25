package model

import "github.com/jinzhu/gorm"

type Courier struct {
	gorm.Model
	CourierName string `gorm:"not null" json:"courier_name"`
}

type CourierMitra struct {
	gorm.Model
	CourierID uint `gorm:"not null" json:"courier_id"`
	MitraID   uint `gorm:"not null" json:"mitra_id"`
}

// type CourierTrx struct {
// 	gorm.Model
// 	MitraTransactionID uint   `gorm:"not null" json:"mitra_transaction_id"`
// 	UserAddressID      uint   `gorm:"not null" json:"user_address_id"`
// 	CourierName        string `gorm:"not null" json:"courier"`
// 	CourierService     string `gorm:"not null" json:"courier_service"`
// 	CourierCost        int    `json:"courier_cost"`
// 	ResiNumber         string `json:"resi_number"`
// }

func (e *Courier) TableName() string {
	return "courier"
}

func (e *CourierMitra) TableName() string {
	return "courier_mitra"
}

// func (e *CourierTrx) TableName() string {
// 	return "courier_trx"
// }
