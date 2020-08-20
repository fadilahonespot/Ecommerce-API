package model

import "github.com/jinzhu/gorm"

type Basket struct {
	gorm.Model
	UserID  uint `gorm:"not null" json:"user_id"`
	MitraID uint `gorm:"not null" json:"mitra_id"`
}

type SubBasket struct {
	gorm.Model
	BasketID  uint `gorm:"not null" json:"basket_id"`
	ProductID uint `gorm:"not null" json:"product_id"`
	Quantity  int  `gorm:"not null" json:"quantity"`
	SubTotal  int  `gorm:"not null" json:"subtotal"`
}

func (e *Basket) TableName() string {
	return "basket"
}

func (e *SubBasket) TableName() string {
	return "sub_basket"
}
