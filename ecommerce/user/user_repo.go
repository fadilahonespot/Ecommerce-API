package user

import (
	"ecommerce/model"

	"github.com/jinzhu/gorm"
)

type UserRepo interface {
	ViewAll()(*[]model.User, error)
	InsertUser(pembeli *model.User, tx *gorm.DB)(*model.User, error)
	ViewByEmail(email string) (*model.User, error)
	UpdateUser(id int, user *model.User, tx *gorm.DB)(*model.User, error)
	ViewById(id int)(*model.User, error)
	InsertOTP(otp *model.UserAuth) (*model.UserAuth, error)
	UpdateOTP(idUser int, otp *model.UserAuth, tx *gorm.DB) (*model.UserAuth, error)
	OtpByIDUser(id int)(*model.UserAuth, error)
	SendSmsOTP(otp string, hp int) error
	InsertNewsletter(news *model.Newletter) (*model.Newletter, error)
	ViewAllNewsletter()(*[]model.Newletter, error)
	InsertMitra(mitra *model.Mitra)(*model.Mitra, error)
	ViewAllMitra()(*[]model.Mitra, error)
	ViewMitraByUserId(id int)(*model.Mitra, error)
	UpdateMitraById(id int, mitra *model.Mitra, tx *gorm.DB)(*model.Mitra, error)
	ViewMitraById(id int)(*model.Mitra, error)
	ViewByMobileNumber(mobileNumber int)(*model.User, error)
}