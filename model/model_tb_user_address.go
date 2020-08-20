package model

import "github.com/jinzhu/gorm"

type UserAddress struct {
	gorm.Model
	UserID       uint   `gorm:"not null" json:"user_id"`
	Label        string `gorm:"not null" json:"label"`
	ReceipedName string `gorm:"not null" json:"receiped_name"`
	MobileNumber string `gorm:"type:bigint; not null" json:"mobile_number"`
	CityID       int    `gorm:"not null" json:"city_id"`
	CityName     string `gorm:"not null" json:"city"`
	Province     string `gorm:"not null" json:"province"`
	PostalCode   string `gorm:"not null" json:"postal_code"`
	Address      string `gorm:"type:text; not null" json:"address"`
}

func (e *UserAddress) TableName() string {
	return "user_address"
}
