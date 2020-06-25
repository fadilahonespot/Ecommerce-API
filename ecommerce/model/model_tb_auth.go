package model

import (
	"github.com/jinzhu/gorm"
)

type UserAuth struct {
	gorm.Model
	UserID             uint   `gorm:"not null" json:"user_id"`
	AuthHP             string `json:"auth_hp"`
	AuthHPStatus       string `gorm:"type:enum('expired', 'active 1 minute'); default:'expired'" json:"auth_hp_status"`
	AuthEmail          string `json:"auth_email"`
	AuthEmailStatus    string `gorm:"type:enum('expired', 'active 5 minute'); default:'expired'" json:"auth_email_status"`
	AuthPassword       string `json:"auth_password"`
	AuthPasswordStatus string `gorm:"type:enum('expired', 'active 5 minute'); default:'expired'" json:"auth_password_status"`
	HPVerify           string `gorm:"type:enum('verify', 'not verify'); default:'not verify'" json:"hp_verify"`
	EmailVerify        string `gorm:"type:enum('verify', 'not verify'); default:'not verify'" json:"email_verify"`
}

func (e *UserAuth) TableName() string {
	return "user_auth"
}
