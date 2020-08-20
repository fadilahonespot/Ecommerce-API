package model

import "github.com/jinzhu/gorm"

type Newletter struct {
	gorm.Model
	DateSend    string `gorm:"type:date; not null" json:"date_send"`
	Subject     string `gorm:"not null" json:"subject"`
	Article     string `gorm:"type:text; not null" json:"article"`
	MessageSend int    `gorm:"not null" json:"succes_send_message"`
	UserID      uint   `gorm:"not null" json:"admin_id"`
}

type NewsletterInput struct {
	Subjec  string `json:"subject"`
	Article string `json:"article"`
}

func (e *Newletter) TableName() string {
	return "newsletter"
}
