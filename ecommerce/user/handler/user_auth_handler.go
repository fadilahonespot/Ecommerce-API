package handler

import (
	"ecommerce/model"
	"ecommerce/utils"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xlzd/gotp"
	"golang.org/x/crypto/bcrypt"
)

func (e *UserHandler) login(c *gin.Context) {
	var logUser = model.Login{}
	err := c.Bind(&logUser)
	if err != nil {
		fmt.Printf("[UserHandler.login] Error bind data %v \n", err)
		utils.HandleError(c, http.StatusBadRequest, "Opsss server somting wrong")
		return
	}
	err = utils.LogginValidation(logUser.Email, logUser.Password)
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, err.Error())
		return
	}
	user, err := e.userUsecase.ViewByEmail(logUser.Email)
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, err.Error())
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(logUser.Password))
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, "Password has wrong")
	return
	}
	token, err := utils.GenerateToken(int(user.ID), user.Password)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	jwt := model.Auth{
		Token: token,
	}
	utils.HandleSucces(c, jwt)
}

func (e *UserHandler) updatePassword(c *gin.Context) {
	user, err := e.userUsecase.ValidUser(c)
	if err != nil {
		utils.HandleError(c, http.StatusUnauthorized, err.Error())
		return
	}
	logAuth := model.UpdatePassword{}
	err = c.Bind(&logAuth)
	if err != nil {
		fmt.Printf("[userHandler.updatePassword] Error bind data %v \n", err)
		utils.HandleError(c, http.StatusInternalServerError, "Oopsss internal server error")
		return
	}
	if len(strings.Split(logAuth.NewPassword, "")) <= 8 {
		utils.HandleError(c, http.StatusBadRequest, "new password minimal 8 characters")
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(logAuth.OldPasswor))
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, "Old password incorrect")
		return
	}
	err = e.userUsecase.UpdatePassword(int(user.ID), logAuth.NewPassword)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.HandleSucces(c, "Update Password Success")
}

func (e *UserHandler) generateOTPPhone(c *gin.Context) {
	user, err := e.userUsecase.ValidUser(c)
	if err != nil {
		utils.HandleError(c, http.StatusUnauthorized, err.Error())
		return
	}
	var userAuth = model.UserAuth{}
	otp := utils.GenerateOTP()
	userAuth.AuthHP = otp

	var ch = make(chan bool)
	userAuth.AuthHPStatus = "active 1 minute"
	authShow := model.AuthHPShow{}
	otpUser, err := e.userUsecase.OtpByIDUser(int(user.ID))
	if err != nil {
		err = e.userUsecase.SendSmsOTP(otp, user.NoHP)
		if err != nil {
			utils.HandleError(c, http.StatusInternalServerError, err.Error())
			return
		}
		userAuth.UserID = uint(user.ID)
		newAuth, err := e.userUsecase.InsertOTP(&userAuth)
		if err != nil {
			utils.HandleError(c, http.StatusInternalServerError, err.Error())
			return
		}
		authShow.Otp = newAuth.AuthHP
		authShow.Active = newAuth.AuthHPStatus
		utils.HandleSucces(c, authShow)
		newAuth.AuthHPStatus = "expired"
		go e.timer(60, ch)
		go e.wacherUpdate(ch, newAuth)
		return
	}
	if otpUser.HPVerify == "verify" {
		utils.HandleError(c, http.StatusForbidden, "your mobile number has been verified")
		return
	}
	err = e.userUsecase.SendSmsOTP(otp, user.NoHP)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	upAuth, err := e.userUsecase.UpdateOTP(int(user.ID), &userAuth)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	authShow.Otp = upAuth.AuthHP
	authShow.Active = upAuth.AuthHPStatus
	utils.HandleSucces(c, authShow)
	upAuth.AuthHPStatus = "expired"
	go e.timer(60, ch)
	go e.wacherUpdate(ch, upAuth)
}

func (e *UserHandler) generateOTPEmail(c *gin.Context) {
	secret := gotp.RandomSecret(42)
	otp := model.UserAuth{}
	var cr = make(chan bool)
	otp.AuthEmail = secret
	otp.AuthEmailStatus = "active 5 minute"
	user, err := e.userUsecase.ValidUser(c)
	if err != nil {
		utils.HandleError(c, http.StatusUnauthorized, err.Error())
		return
	}
	link := "http://" + string(c.Request.Host) + "/user/verify/email/" + strconv.Itoa(int(user.ID)) + "/" + otp.AuthEmail
	authShow := model.AuthEmailShow{}
	userAuth, err := e.userUsecase.OtpByIDUser(int(user.ID))
	if err != nil {
		otp.UserID = user.ID
		userAuth, err := e.userUsecase.InsertOTP(&otp)
		if err != nil {
			utils.HandleError(c, http.StatusInternalServerError, err.Error())
			return
		}
		utils.WelcomeEmail("Email Verifikasi Ecommerce", "asset/template/email_verifikasi.html", user.Email, link, "")
		authShow.Auth = userAuth.AuthEmail
		authShow.Active = userAuth.AuthEmailStatus
		utils.HandleSucces(c, authShow)
		userAuth.AuthEmailStatus = "expired"
		go e.timer(300, cr)
		go e.wacherUpdate(cr, userAuth)
		return
	}
	if userAuth.EmailVerify == "verify" {
		utils.HandleError(c, http.StatusForbidden, "Your email has been verified")
		return
	}
	upAuth, err := e.userUsecase.UpdateOTP(int(user.ID), &otp)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.WelcomeEmail("Email Verifikasi Ecommerce", "asset/template/email_verifikasi.html", user.Email, link, "")
	authShow.Auth = upAuth.AuthEmail
	authShow.Active = upAuth.AuthEmailStatus
	utils.HandleSucces(c, authShow)
	upAuth.AuthEmailStatus = "expired"
	go e.timer(300, cr)
	go e.wacherUpdate(cr, upAuth)
}

func (e *UserHandler) verifyOTPHp(c *gin.Context) {
	user, err := e.userUsecase.ValidUser(c)
	if err != nil {
		utils.HandleError(c, http.StatusUnauthorized, err.Error())
	}
	otp := c.PostForm("otp")
	userAuth, err := e.userUsecase.OtpByIDUser(int(user.ID))
	if err != nil {
		utils.HandleError(c, http.StatusNotFound, err.Error())
		return
	}
	if userAuth.AuthHP != otp {
		utils.HandleError(c, http.StatusForbidden, "authenfication failed, Incorrect otp code entered")
		return
	} 
	if userAuth.AuthHPStatus == "expired" {
		utils.HandleError(c, http.StatusForbidden, "OTP code has expired")
		return
	}
	err = e.userUsecase.OtpHandphoneVerifyUpdate(userAuth)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.HandleSucces(c, "Succes verification handphone number")
}

func (e *UserHandler) vertifyOTPEmail(c *gin.Context) {
	otp := c.Param("otp")
	idUser := c.Param("idUser")
	id, err := strconv.Atoi(idUser)
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, "authenfication failed, link is not valid")
		return
	}
	userAuth, err := e.userUsecase.OtpByIDUser(id)
	if err != nil {
		utils.HandleError(c, http.StatusNotFound, "authenfication failed, link is not valid")
		return
	}
	if userAuth.AuthEmail != otp {
		utils.HandleError(c, http.StatusBadRequest, "authenfication failed, link is not valid")
		return
	}
	if userAuth.AuthEmailStatus == "expired" {
		utils.HandleError(c, http.StatusForbidden, "verification link has expired")
		return
	}
	err = e.userUsecase.OtpEmailVerifyUpdate(userAuth)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	// utils.HandleSucces(c, logUser)
	c.Redirect(http.StatusPermanentRedirect, "/user/email/succes")
}

func (e *UserHandler) generateAuthForgetPass(c *gin.Context) {
	var email = model.InputEmailForget{}
	err := c.Bind(&email)
	if err != nil {
		fmt.Printf("[UserHandler.generateForgetPass] Error bind data %v \n", err)
		utils.HandleError(c, http.StatusInternalServerError, "Oppss internal server error")
		return
	}
	user, err := e.userUsecase.ViewByEmail(email.Email)
	if err != nil {
		utils.HandleError(c, http.StatusNotFound, err.Error())
		return
	}
	auth := gotp.RandomSecret(40)
	var userAuth = model.UserAuth{}
	userAuth.AuthPasswordStatus = "active 5 minute"
	var showAuth = model.ShowAuthForgetPass{}
	userAuth.AuthPassword = auth
	var cd = make(chan bool)
	link := "http://" + string(c.Request.Host) + "/user/password/reset/" + strconv.Itoa(int(user.ID)) + "/" + auth
	_, err = e.userUsecase.OtpByIDUser(int(user.ID))
	if err != nil {
		userAuth.UserID = user.ID
		newAuth, err := e.userUsecase.InsertOTP(&userAuth)
		if err != nil {
			utils.HandleError(c, http.StatusInternalServerError, err.Error())
			return
		}
		showAuth.Auth = newAuth.AuthPassword
		showAuth.Active = newAuth.AuthPasswordStatus
		utils.ResetPassEmail(user.Email, "Reset Password Akun Ecommerce", "asset/template/reset_password.html", link)
		utils.HandleSucces(c, showAuth)
		newAuth.AuthPasswordStatus = "expired"
		go e.timer(300, cd)
		go e.wacherUpdate(cd, newAuth)
		return
	}
	upAuth, err := e.userUsecase.UpdateOTP(int(user.ID), &userAuth)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	showAuth.Auth = upAuth.AuthPassword
	showAuth.Active = upAuth.AuthPasswordStatus
	utils.ResetPassEmail(user.Email, "Reset Password Akun Ecommerce", "asset/template/reset_password.html", link)
	utils.HandleSucces(c, showAuth)
	upAuth.AuthPasswordStatus = "expired"
	go e.timer(300, cd)
	go e.wacherUpdate(cd, upAuth)
}

func (e *UserHandler) verifyAuthForgetPass(c *gin.Context) {
	idStr := c.Param("idUser")
	auth := c.Param("auth")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.HandleError(c, http.StatusForbidden, "link is not valid")
		return
	}
	userAuth, err := e.userUsecase.OtpByIDUser(id)
	if err != nil {
		utils.HandleError(c, http.StatusForbidden, "link is not valid")
		return
	}
	if auth != userAuth.AuthPassword {
		utils.HandleError(c, http.StatusForbidden, "link is not valid")
		return
	}

	if userAuth.AuthPasswordStatus == "expired" {
		utils.HandleError(c, http.StatusForbidden, "link is expired")
		return
	}

	pass := gotp.RandomSecret(6)
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), 10)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, "Oopss server has wrong")
		return
	}
	user, err := e.userUsecase.ViewById(int(userAuth.UserID))
	if err != nil {
		utils.HandleError(c, http.StatusNotFound, err.Error())
		return
	}
	user.Password = string(hash)
	upUser, err := e.userUsecase.UpdateUser(int(userAuth.UserID), user)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.ResetPassEmailSucces(upUser.Email, "Reset Password Succes", "asset/template/email_succes_forget_pass.html", pass)
	c.Redirect(http.StatusPermanentRedirect, "/user/password/forget/succes")
}

func (e *UserHandler) timer(timeout int, ch chan<- bool) {
	time.AfterFunc(time.Duration(timeout)*time.Second, func(){
		ch <- true
	})
}

func (e *UserHandler) wacherUpdate(ch <-chan bool, auth *model.UserAuth) {
	<-ch
	auth.EmailVerify = ""
	auth.HPVerify = ""
	_, err := e.userUsecase.UpdateOTP(int(auth.UserID), auth)
		if err != nil {
			fmt.Printf("[Userhandler.wacherUpdate] failed update otp in expired %v \n", err)
			return
		}
}

func (e *UserHandler) succesResetForgetPass(c *gin.Context) {
	temp, err := utils.ReadTemplateHtml("asset/template/page_succes_change_pass.html")
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, "failed redirect page")
		return
	}
	c.Data(http.StatusOK, "text/html; charset=utf-8", temp)
}

func (e *UserHandler) succesMessageEmailVerification(c *gin.Context) {
	temp, err := utils.ReadTemplateHtml("asset/template/succes_verification_email.html")
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, "failed redirect page")
		return
	}
	c.Data(http.StatusOK, "text/html; charset=utf-8", temp)
}

