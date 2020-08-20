package model

import "github.com/jinzhu/gorm"

type Watchlist struct {
	gorm.Model
	ProductID uint
	UserID uint
}

func (e *Watchlist) TableName() string {
	return "wachlist"
}