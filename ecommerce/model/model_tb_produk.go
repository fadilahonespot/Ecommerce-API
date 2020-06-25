package model

import (
	"github.com/jinzhu/gorm"
)

type Product struct { //in `proto` package
	gorm.Model
	Name          string `gorm:"not null" json:"name"`
	Prince        int    `gorm:"not null" json:"prince"`
	Stock         string `gorm:"not null" json:"stock"`
	Weight        int    `gorm:"type:bigint; not null" json:"weight"`
	Condition     string `gorm:"type:enum('new', 'secound'); not null" json:"condition"`
	Brand         string `gorm:"not null" json:"brand"`
	Description   string `gorm:"type:text; not null" json:"description"`
	MinPurchase   int    `gorm:"not null" json:"minimal_purchase"`
	Sold          int    `gorm:"default:0" json:"sold"`
	MitraID       uint   `gorm:"not null" json:"mitra_id"`
	SubcatagoryID uint   `gorm:"not null" json:"subcatagory_id"`
}

type ProductReview struct {
	gorm.Model
	UserID    uint   `gorm:"not null" json:"user_id"`
	ProductID uint   `gorm:"not null" json:"product_id"`
	Rating    int    `gorm:"type:int(5); not null" json:"rating"`
	Review    string `gorm:"not null" json:"review"`
}

type ProductImg struct {
	gorm.Model
	ProductID uint   `gorm:"not null" json:"product_id"`
	PartImg   string `gorm:"not null" json:"part_img"`
}

type ProductDiscussion struct {
	gorm.Model
	ProductID  uint   `gorm:"not null" json:"product_id"`
	UserID     uint   `gorm:"not null" json:"user_id"`
	Discussion string `gorm:"not null" json:"discussion"`
}

type ProductSubDiscussion struct {
	gorm.Model
	ProductDiscussionID uint   `gorm:"not null" json:"product_discussion_id"`
	UserID              uint   `gorm:"not null" json:"user_id"`
	ReplyDiscussion     string `json:"reply discussion"`
}

func (e *Product) TableName() string {
	return "product"
}

func (e *ProductReview) TableName() string {
	return "product_review"
}

func (e *ProductImg) TableName() string {
	return "product_img"
}

func (e *ProductDiscussion) TableName() string {
	return "product_discussion"
}
