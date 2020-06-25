package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Nama     string `gorm:"not null" json:"name"`
	Email    string `gorm:"not null" json:"email"`
	NoHP     int    `gorm:"type:bigint; not null" json:"mobile_number"`
	IDKota   int    `gorm:"not null" json:"city_id"`
	Kota     string `json:"city"`
	Provinsi string `json:"province"`
	KodePos  string `json:"postal_code"`
	Alamat   string `gorm:"type:text; not null" json:"addres"`
	Password string `gorm:"not null" json:"password"`
	Role     string `gorm:"type:enum('user', 'mitra', 'admin'); default:'user'" json:"role"`
	Status   string `gorm:"type:enum('vertifiled', 'not vertifiled'); default:'not vertifiled'" json:"status"`
}

func (e *User) TableName() string {
	return "user"
}
