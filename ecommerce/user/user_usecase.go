package user

import (
	"ecommerce/model"

	"github.com/gin-gonic/gin"
)

type UserUsecase interface {
	ViewAll()(*[]model.User, error)
	InsertDataUser(pembeli *model.User) (*model.User, error)
	ViewByEmail(email string) (*model.User, error)
	UpdateUser(id int, user *model.User)(*model.User, error)
	ViewById(id int)(*model.User, error)
	RequestMitraProcess(id int, status string) (*model.User, *model.Mitra, error)
	UpdatePassword(id int, pass string) error
	InsertOTP(otp *model.UserAuth) (*model.UserAuth, error)
	UpdateOTP(idUser int, otp *model.UserAuth) (*model.UserAuth, error)
	OtpByIDUser(id int)(*model.UserAuth, error)
	SendSmsOTP(otp string, hp int) error
	InsertNewsletter(news *model.Newletter) (*model.Newletter, error)
	ViewAllNewsletter()(*[]model.Newletter, error)
	InsertMitra(mitra *model.Mitra)(*model.Mitra, error)
	ViewAllMitra()(*[]model.Mitra, error)
	ViewMitraByUserId(id int)(*model.Mitra, error)
	UpdateMitraById(id int, mitra *model.Mitra)(*model.Mitra, error)
	ViewMitraById(id int)(*model.Mitra, error)
	DataUploadMitra(c *gin.Context, idUser int)(*model.Mitra, []string, error)
	ValidMitra(c *gin.Context) (*model.User, *model.Mitra, error)
	AdminOnly(c *gin.Context) (*model.User, error)
	ValidUser(c *gin.Context) (*model.User, error)
	ViewByMobileNumber(mobileNumber int)(*model.User, error)
	OtpEmailVerifyUpdate(userAuth *model.UserAuth) error
	OtpHandphoneVerifyUpdate(userAuth *model.UserAuth) error
}