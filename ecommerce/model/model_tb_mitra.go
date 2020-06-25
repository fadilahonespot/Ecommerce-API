package model

import (
	"github.com/jinzhu/gorm"
)

type Mitra struct {
	gorm.Model
	UserID           uint   `gorm:"not null" json:"user_id"`
	StoreName        string `gorm:"not null" json:"store_name"`
	CityID           int    `gorm:"not null" json:"city_id"`
	City             string `gorm:"not null" json:"city"`
	Province         string `gorm:"not null" json:"province"`
	PostalCode       string `gorm:"not null" json:"postal_code"`
	Address          string `gorm:"not null" json:"address"`
	KTPNumber        string `gorm:"not null" json:"ktp_number"`
	NPWPNumber       string `gorm:"not null" json:"npwp_number"`
	BankName         string `gorm:"not null" json:"bank_name"`
	BankAccounName   string `gorm:"not null" json:"bank_accoun_name"`
	BankAccounNumber string `gorm:"not null" json:"bank_accoun_number"`
	Status           string `gorm:"type:enum('approved', 'waiting for approval', 'rejected'); default:'waiting for approval'" json:"status"`
	KTPImage         string `gorm:"not null" json:"ktp_image"`
	NPWPImage        string `gorm:"not null" json:"npwp_image"`
	KTPSelfieImage   string `gorm:"not null" json:"ktp_selfie_img"`
	MitraImg         string `json:"mitra_profile_photo"`
	MitraCover       string `json:"mitra_profile_cover"`
}

func (e *Mitra) TableName() string {
	return "mitra"
}
