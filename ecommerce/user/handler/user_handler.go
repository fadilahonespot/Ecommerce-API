package handler

import (
	"ecommerce/middleware"
	"ecommerce/model"
	"ecommerce/user"
	"ecommerce/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/xlzd/gotp"
)

type UserHandler struct {
	userUsecase user.UserUsecase
}

func CreateUserHandler(r *gin.Engine, userUsecase user.UserUsecase) {
	userHandler := UserHandler{userUsecase}

	r.POST("/register", userHandler.register)
	r.POST("/login", userHandler.login)
	r.GET("/user/verify/email/:idUser/:otp", userHandler.vertifyOTPEmail)
	r.GET("/user/password/reset/:idUser/:auth", userHandler.verifyAuthForgetPass)
	r.POST("/user/password/forget", userHandler.generateAuthForgetPass)

	// Redirect page`
	r.GET("/user/email/succes", userHandler.succesMessageEmailVerification)
	r.GET("/user/password/forget/succes", userHandler.succesResetForgetPass)

	r2 := r.Group("/user").Use(middleware.TokenVerifikasiMiddleware())
	r2.GET("/profile", userHandler.userProfile)
	r2.PUT("/update", userHandler.userUpdate)
	r2.PUT("/password", userHandler.updatePassword)
	r2.GET("/otp/phone", userHandler.generateOTPPhone)
	r2.GET("/otp/email", userHandler.generateOTPEmail)
	r2.GET("/verify/hp", userHandler.verifyOTPHp)
	r2.POST("/upgrate", userHandler.upgrateMitra)

	r3 := r.Group("/admin").Use(middleware.TokenVerifikasiMiddleware())
	r3.PUT("/user/:id", userHandler.userUpdateAdmin)
	r3.GET("/user", userHandler.viewAllUser)
	r3.POST("/newsletter", userHandler.blashEmailVertifiled)
	r3.GET("/newsletter", userHandler.historyNewsletter)
	r3.GET("/user/:id", userHandler.getUserById)
	r3.GET("/mitra/waiting", userHandler.checkMitraWaiting)
	r3.PUT("/mitra/approved/:id", userHandler.approveMitra)
	r3.PUT("/mitra/reject/:id", userHandler.rejectMitra)

	r4 := r.Group("/mitra").Use(middleware.TokenVerifikasiMiddleware())
	r4.GET("/profile", userHandler.mitraProfile)
	r4.PUT("/profile", userHandler.updateMitraProfile)
	r4.PUT("/profile/photo", userHandler.updateProfileImg)
}


func (e *UserHandler) register(c *gin.Context) {
	var user = model.User{}
	err := c.Bind(&user)
	if err != nil {
		fmt.Printf("[UserHandler.insertUser] error bind data %v \n", err)
		utils.HandleError(c, http.StatusBadRequest, "Oppss Sever has be wrong")
		return
	}
	err = utils.RegisterValidasi(&user)
	if err != nil {
		utils.HandleError(c, http.StatusForbidden, err.Error())
		return
	}
	_, err = e.userUsecase.ViewByEmail(user.Email)
	if err == nil {
	utils.HandleError(c, http.StatusConflict, "email is already registered")
		return
	}
	noHp, err := strconv.Atoi("62" + strconv.Itoa(user.NoHP))
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, "no hp has be number")
		return
	}
	user.NoHP = noHp
	_, err = e.userUsecase.ViewByMobileNumber(noHp)
	if err == nil {
		utils.HandleError(c, http.StatusConflict, "mobile number already registered")
		return
	}
	newUser, err := e.userUsecase.InsertDataUser(&user)
	if err != nil {
	utils.HandleError(c, http.StatusBadRequest, err.Error())
		return
	}
	utils.HandleSucces(c, newUser)
}

func (e *UserHandler) userProfile(c *gin.Context) {
	user, err := e.userUsecase.ValidUser(c)
	if err != nil {
		utils.HandleError(c, http.StatusUnauthorized, err.Error())
		return
	}
	utils.HandleSucces(c, user)
}

func (e *UserHandler) userUpdate(c *gin.Context) {
	user, err := e.userUsecase.ValidUser(c)
	if err != nil {
		utils.HandleError(c, http.StatusForbidden, err.Error())
		return
	}
	var logUser = model.User{}
	err = c.Bind(&logUser)
	if err != nil {
		fmt.Printf("[UserHandler.userUpdate] Error bind data body %v \n", err)
		utils.HandleError(c, http.StatusInternalServerError, "Ooopss internal server error")
		return
	}
	if logUser.ID != 0 || logUser.Password != "" || logUser.Role != "" || logUser.Status != "" {
		utils.HandleError(c, http.StatusForbidden, "update not allowed")
		return
	}
	noHp, err := strconv.Atoi("62" + strconv.Itoa(logUser.NoHP))
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, "mobile number has be number")
		return
	}
	if logUser.Email != user.Email || logUser.NoHP != user.NoHP {
		allUser, err := e.userUsecase.ViewAll()
		if err != nil {
			utils.HandleError(c, http.StatusInternalServerError, err.Error())
			return
		}
		for i := 0; i < len(*allUser); i++ {
			if (*allUser)[i].Email == logUser.Email {
				utils.HandleError(c, http.StatusConflict, "email is already in another account")
				return
			}
			if (*allUser)[i].NoHP == noHp {
				utils.HandleError(c, http.StatusConflict, "mobile number is already in another account")
				return
			}
		}
	}
	var userAuth = model.UserAuth{}
	if user.Email != logUser.Email && logUser.Email != "" {
		userAuth.AuthEmailStatus = "not verify"
		_, err := e.userUsecase.UpdateOTP(int(user.ID), &userAuth)
		if err != nil {
			utils.HandleError(c, http.StatusInternalServerError, err.Error())
			return
		}
		user.Status = "not vertifiled"
		_, err = e.userUsecase.UpdateUser(int(user.ID), user)
		if err != nil {
			utils.HandleError(c, http.StatusInternalServerError, err.Error())
			return
		}
	}
	if user.NoHP != logUser.NoHP && user.NoHP != 0 {
		userAuth.AuthHPStatus = "not verify"
		_, err = e.userUsecase.UpdateOTP(int(user.ID), &userAuth)
		if err != nil {
			utils.HandleError(c, http.StatusInternalServerError, err.Error())
			return
		}
		user.Status = "not vertifiled"
		_, err = e.userUsecase.UpdateUser(int(user.ID), user)
		if err != nil {
			utils.HandleError(c, http.StatusInternalServerError, err.Error())
			return
		}
	}

	upUser, err := e.userUsecase.UpdateUser(int(user.ID), &logUser)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.HandleSucces(c, upUser)
}

func (e *UserHandler) upgrateMitra(c *gin.Context) {
	user, err := e.userUsecase.ValidUser(c)
	if err != nil {
		utils.HandleError(c, http.StatusNotFound, err.Error())
		return
	}
	if user.Status == "not vertifiled" {
		utils.HandleError(c, http.StatusForbidden, "verify your email and cellphone number to submit an upgrade to partners")
		return
	}
	mitra, err := e.userUsecase.ViewMitraByUserId(int(user.ID))
	if err == nil {
		if mitra.Status == "waiting for approval" {
			utils.HandleError(c, http.StatusForbidden, "You've registered before, your data is being reviewed")
			return
		}
		if mitra.Status == "approved" {
			utils.HandleError(c, http.StatusForbidden, "You are already registered as our partner")
			return
		}
		dataMitra, partImg, err := e.userUsecase.DataUploadMitra(c, int(user.ID))
		if err != nil {
			utils.HandleError(c, http.StatusForbidden, err.Error())
			return
		}
		dataMitra.Status = "waiting for approval"
		upMitra, err := e.userUsecase.UpdateMitraById(int(mitra.ID), dataMitra)
		if err != nil {
			utils.ValidationRollbackImage(partImg)
			utils.HandleError(c, http.StatusInternalServerError, "Oopss server someting wrong")
			return
		}
		utils.WelcomeEmail("Ecommerce - Permintaan Upgrate Mitra Sedang Ditinjau Kembali","", user.Email, "", "asset/template/email_permintaan_upgrate_mitra.html")
		utils.HandleSucces(c, upMitra)
		return
	}

	reqMitra, partImg, err := e.userUsecase.DataUploadMitra(c, int(user.ID))
	if err != nil {
		utils.HandleError(c, http.StatusForbidden, err.Error())
		return
	}
	newMitra, err := e.userUsecase.InsertMitra(reqMitra)
	if err != nil {
		utils.ValidationRollbackImage(partImg)
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	
	utils.WelcomeEmail("Ecommerce - Permintaan Upgrate Mitra Sedang Ditinjau","", user.Email, "", "asset/template/email_permintaan_upgrate_mitra.html")
	utils.HandleSucces(c, newMitra)
}

func (e *UserHandler) mitraProfile(c *gin.Context) {
	_, mitra, err := e.userUsecase.ValidMitra(c)
	if err != nil {
		utils.HandleError(c, http.StatusUnauthorized, err.Error())
		return
	}
	utils.HandleSucces(c, mitra)
}

func (e *UserHandler) updateMitraProfile(c *gin.Context) {
	_, mitra, err := e.userUsecase.ValidMitra(c)
	if err != nil {
		utils.HandleError(c, http.StatusUnauthorized, err.Error())
		return
	}
	var inMitra = model.Mitra{}
	err = c.Bind(&inMitra)
	if err != nil {
		fmt.Printf("[updateMitraProfile] Error bind data mitra %v \n", err)
		utils.HandleError(c, http.StatusInternalServerError, "Oopss internal server error")
		return
	}
	if inMitra.ID != 0 || inMitra.KTPImage != "" || inMitra.KTPNumber != "" || inMitra.KTPSelfieImage != "" || inMitra.NPWPImage != "" || inMitra.NPWPNumber != "" || inMitra.Status != "" || inMitra.UserID != 0 {
		utils.HandleError(c, http.StatusBadRequest, "one of the update data is not permitted, please contact the admin")
		return
	}
	if inMitra.CityID == 0 {
		inMitra.CityID = mitra.CityID
	}
	upMitra, err := e.userUsecase.UpdateMitraById(int(mitra.ID), &inMitra)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.HandleSucces(c, upMitra)
}

func (e *UserHandler) updateProfileImg(c *gin.Context) {
	_, mitra, err := e.userUsecase.ValidMitra(c)
	if err != nil {
		utils.HandleError(c, http.StatusUnauthorized, err.Error())
		return
	}
	str := []string{"profile_photo", "profile_cover"}
	var partImg []string
	for i := 0; i < len(str); i++ {
		file, err := c.FormFile(str[i])
		if err != nil {
			utils.ValidationRollbackImage(partImg)
			fmt.Printf("[userHandler.updateProfileImg] error from file %v \n", err)
			utils.HandleError(c, http.StatusInternalServerError, "Oopss internal server error")
			return
		}
		err = utils.FileImgValidation(file, partImg)
		if err != nil {
			utils.HandleError(c, http.StatusForbidden, err.Error())
			return
		}
		name := gotp.RandomSecret(12)
		part := viper.GetString("part.images") + name + ".jpg"
		err = c.SaveUploadedFile(file, part)
		if err != nil {
			utils.ValidationRollbackImage(partImg)
			fmt.Printf("[userHandler.updateProfileImg] error save file upload %v \n", err)
			utils.HandleError(c, http.StatusInternalServerError, "Oopss internal server errror")
			return
		}
		partImg = append(partImg, part)
	}
	var inMitra = model.Mitra{
		MitraImg: partImg[0],
		MitraCover: partImg[1],
		CityID: mitra.CityID,
	}
	upMitra, err := e.userUsecase.UpdateMitraById(int(mitra.ID), &inMitra)
	if err != nil {
		utils.ValidationRollbackImage(partImg)
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	if mitra.MitraImg != "" || mitra.MitraCover != "" {
		var oldImg = []string{mitra.MitraImg, mitra.MitraCover}
		utils.ValidationRollbackImage(oldImg)
	}
	utils.HandleSucces(c, upMitra)
}





