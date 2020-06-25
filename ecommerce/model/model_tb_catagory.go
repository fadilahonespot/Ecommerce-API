package model

import "github.com/jinzhu/gorm"

type ResponsCat struct {
	Success bool       `json:"Success" example:"true"`
	Message string     `json:"Message" example:"Success"`
	Data    []Catagory `json:"data,omitempty"`
}

type Catagory struct { //in `proto` package
	gorm.Model
	UserID       uint   `gorm:"not null" json:"user_id" example:"2"`
	CatagoryName string `gorm:"not null" json:"catagory_name" example:"Electronik"`
}

type SubCatagory struct {
	gorm.Model
	SubCatagoryName string `gorm:"not null" json:"subcatagory_name"`
	CatagoryID      uint   `gorm:"not null" json:"catagory_id"`
	UserID          uint   `gorm:"not null" json:"user_id"`
}

func (e *Catagory) TableName() string {
	return "catagory"
}

func (e *SubCatagory) TableName() string {
	return "subcatagory"
}
