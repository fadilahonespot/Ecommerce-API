package usecase

import (
	"ecommerce/middleware"
	"ecommerce/model"
	"ecommerce/ongkir"
	"ecommerce/transaction"
	"ecommerce/user"
	"ecommerce/utils"
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"github.com/xlzd/gotp"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecaseImpl struct {
	userRepo        user.UserRepo
	ongkirRepo      ongkir.OngkirRepo
	transactionRepo transaction.TransactionRepo
}

func CreateUserUsecase(userRepo user.UserRepo, ongkirRepo ongkir.OngkirRepo, transactionRepo transaction.TransactionRepo) user.UserUsecase {
	return &UserUsecaseImpl{userRepo, ongkirRepo, transactionRepo}
}

func (e *UserUsecaseImpl) ViewAll() (*[]model.User, error) {
	return e.userRepo.ViewAll()
}

func (e *UserUsecaseImpl) InsertDataUser(user *model.User) (*model.User, error) {
	tx := e.ongkirRepo.BeginTrans()
	city, err := e.ongkirRepo.GetCityById(user.IDKota)
	if err != nil {
		fmt.Printf("[UserUsecaseImpl.InsertDataUser] error get data city, %v \n", err)
		return nil, err
	}

	user.Kota = city.Type + " " + city.CityName
	user.Provinsi = city.Province
	user.KodePos = city.PostalCode
	user.Role = "user"

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		fmt.Printf("[UserUsecase.InsertDataUser] Error hashing password %v \n", err)
		return nil, fmt.Errorf("Oopss internal server error")
	}

	user.Password = string(hash)
	newUser, err := e.userRepo.InsertUser(user, tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	address := model.UserAddress{
		UserID:       newUser.ID,
		Label:        "Default",
		ReceipedName: newUser.Nama,
		MobileNumber: strconv.Itoa(newUser.NoHP),
		CityID:       newUser.IDKota,
		CityName:     city.CityName,
		Province:     city.Province,
		PostalCode:   city.PostalCode,
		Address:      newUser.Alamat,
	}
	_, err = e.transactionRepo.AddUserAddress(&address, tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	utils.WelcomeEmail("Selamat Datang di Ecommerce", "", newUser.Email, "", "asset/template/welcome_message.html")
	tx.Commit()
	return newUser, nil
}

func (e *UserUsecaseImpl) ViewByEmail(email string) (*model.User, error) {
	return e.userRepo.ViewByEmail(email)
}

func (e *UserUsecaseImpl) UpdateUser(id int, user *model.User) (*model.User, error) {
	tx := e.ongkirRepo.BeginTrans()
	city, err := e.ongkirRepo.GetCityById(user.IDKota)
	if err != nil {
		fmt.Printf("[UserUsecaseImpl.UpdateDataUser] error get data city, %v \n", err)
		return nil, err
	}

	user.Kota = city.Type + " " + city.CityName
	user.Provinsi = city.Province
	user.KodePos = city.PostalCode

	upUser, err := e.userRepo.UpdateUser(id, user, tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	address := model.UserAddress{
		UserID:       upUser.ID,
		Label:        "Default",
		ReceipedName: upUser.Nama,
		MobileNumber: strconv.Itoa(upUser.NoHP),
		CityID:       upUser.IDKota,
		CityName:     city.CityName,
		Province:     city.Province,
		PostalCode:   city.PostalCode,
		Address:      upUser.Alamat,
	}
	allAddress, err := e.transactionRepo.ViewUserAddressByUserId(int(upUser.ID))
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	_, err = e.transactionRepo.UpdateUserAddressById(int((*allAddress)[0].ID), &address, tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return upUser, nil
}

func (e *UserUsecaseImpl) UpdateOTP(idUser int, otp *model.UserAuth) (*model.UserAuth, error) {
	tx := e.ongkirRepo.BeginTrans()
	userAuth, err := e.userRepo.UpdateOTP(idUser, otp, tx)
	if err != nil {
		return nil, err
	}
	tx.Commit()
	return userAuth, nil
}

func (e *UserUsecaseImpl) ViewById(id int) (*model.User, error) {
	return e.userRepo.ViewById(id)
}

func (e *UserUsecaseImpl) UpdatePassword(id int, pass string) error {
	tx := e.ongkirRepo.BeginTrans()
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), 10)
	if err != nil {
		fmt.Printf("[UserUsecase.UpdatePassword] Error hashing password %v \n", err)
		return fmt.Errorf("Oopss internal server error")
	}
	var user = model.User{}
	user.Password = string(hash)
	_, err = e.userRepo.UpdateUser(id, &user, tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (e *UserUsecaseImpl) InsertOTP(otp *model.UserAuth) (*model.UserAuth, error) {
	return e.userRepo.InsertOTP(otp)
}

func (e *UserUsecaseImpl) OtpByIDUser(id int) (*model.UserAuth, error) {
	return e.userRepo.OtpByIDUser(id)
}

func (e *UserUsecaseImpl) SendSmsOTP(otp string, hp int) error {
	return e.userRepo.SendSmsOTP(otp, hp)
}

func (e *UserUsecaseImpl) InsertNewsletter(news *model.Newletter) (*model.Newletter, error) {
	return e.userRepo.InsertNewsletter(news)
}

func (e *UserUsecaseImpl) ViewAllNewsletter() (*[]model.Newletter, error) {
	return e.userRepo.ViewAllNewsletter()
}

func (e *UserUsecaseImpl) InsertMitra(mitra *model.Mitra) (*model.Mitra, error) {
	city, err := e.ongkirRepo.GetCityById(mitra.CityID)
	if err != nil {
		fmt.Printf("[userUsecase.InsertMitra] message error: %v \n", err)
		return nil, err
	}
	mitra.City = city.Type + " " + city.CityName
	mitra.Province = city.Province
	mitra.PostalCode = city.PostalCode
	newMitra, err := e.userRepo.InsertMitra(mitra)
	if err != nil {
		fmt.Printf("[userUsecase.InsertMitra] message error: %v \n", err)
		return nil, err
	}
	return newMitra, nil
}

func (e *UserUsecaseImpl) RequestMitraProcess(id int, status string) (*model.User, *model.Mitra, error) {
	tx := e.ongkirRepo.BeginTrans()
	mitra, err := e.userRepo.ViewMitraById(id)
	if err != nil {
		return nil, nil, err
	}
	if mitra.Status != "waiting for approval" {
		return nil, nil, fmt.Errorf("Partner status is not waiting for approval")
	}
	insertMitra := model.Mitra{}
	insertMitra.Status = status
	insertMitra.CityID = mitra.CityID
	upMitra, err := e.userRepo.UpdateMitraById(id, &insertMitra, tx)
	if err != nil {
		tx.Rollback()
		return nil, nil, err
	}

	if upMitra.Status == "approved" {
		couriers, err := e.ongkirRepo.ViewAllCourier()
		if err != nil {
			tx.Rollback()
			return nil, nil, err
		}
		for i := 0; i < len(*couriers); i++ {
			if i < 3 {
				var courierMitra = model.CourierMitra{
					CourierID: (*couriers)[i].ID,
					MitraID: upMitra.ID,
				}
				_, err = e.ongkirRepo.InsertCourierMitra(&courierMitra, tx)
				if err != nil {
					tx.Rollback()
					return nil, nil, err
				}
			}
		}
	}
	user, err := e.userRepo.ViewById(int(upMitra.UserID))
	if err != nil {
		tx.Rollback()
		return nil, nil, err
	}
	tx.Commit()
	return user, upMitra, nil
}

func (e *UserUsecaseImpl) ViewAllMitra() (*[]model.Mitra, error) {
	return e.userRepo.ViewAllMitra()
}

func (e *UserUsecaseImpl) ViewMitraByUserId(id int) (*model.Mitra, error) {
	return e.userRepo.ViewMitraByUserId(id)
}

func (e *UserUsecaseImpl) UpdateMitraById(id int, mitra *model.Mitra) (*model.Mitra, error) {
	tx := e.ongkirRepo.BeginTrans()
	city, err := e.ongkirRepo.GetCityById(mitra.CityID)
	if err != nil {
		fmt.Printf("[userUsecaseImpl.UpdateMitraById] error get city by id")
		return nil, err
	}
	mitra.City = city.Type + " " + city.CityName
	mitra.Province = city.Province
	mitra.PostalCode = city.PostalCode
	upMitra, err := e.userRepo.UpdateMitraById(id, mitra, tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return upMitra, nil
}

func (e *UserUsecaseImpl) ViewMitraById(id int) (*model.Mitra, error) {
	return e.userRepo.ViewMitraById(id)
}

func (e *UserUsecaseImpl) DataUploadMitra(c *gin.Context, idUser int) (*model.Mitra, []string, error) {
	var param = []string{"ktp", "npwp", "ktp_selfie"}
	var partImg []string
	for k := 0; k < len(param); k++ {
		files, err := c.FormFile(param[k])
		if err != nil {
			utils.ValidationRollbackImage(partImg)
			return nil, nil, fmt.Errorf("failed upload, file not found")
		}
		err = utils.FileImgValidation(files, partImg)
		if err != nil {
			return nil, nil, err
		}

		name := gotp.RandomSecret(12)
		part := viper.GetString("part.images") + name + ".jpg"
		err = c.SaveUploadedFile(files, part)
		if err != nil {
			fmt.Printf("[UserUsecase.DataUploadMitra] Error save images %v \n", err)
			utils.ValidationRollbackImage(partImg)
			return nil, nil, fmt.Errorf("Ooppss server someting wrong")
		}
		partImg = append(partImg, part)
	}
	var inputText = []string{"city_id", "store_name", "addres", "ktp_number", "npwp_number", "bank_name", "bank_accoun_name", "bank_accoun_number"}
	var dataTxt []string
	for j := 0; j < len(inputText); j++ {
		data := c.PostForm(inputText[j])
		if data == "" {
			utils.ValidationRollbackImage(partImg)
			return nil, nil, fmt.Errorf("column must be filled")
		}
		dataTxt = append(dataTxt, data)
	}
	cityId, err := strconv.Atoi(dataTxt[0])
	if err != nil {
		utils.ValidationRollbackImage(partImg)
		return nil, nil, fmt.Errorf("city id has be number")
	}
	if len(strings.Split(dataTxt[3], "")) != 16 {
		utils.ValidationRollbackImage(partImg)
		return nil, nil, fmt.Errorf("ktp nik must be 16 digits")
	}
	_, err = strconv.Atoi(dataTxt[3])
	if err != nil {
		utils.ValidationRollbackImage(partImg)
		return nil, nil, fmt.Errorf("ktp nik must be a number")
	}

	reqMitra := model.Mitra{
		UserID:           uint(idUser),
		StoreName:        dataTxt[1],
		CityID:           cityId,
		Address:          dataTxt[2],
		KTPNumber:        dataTxt[3],
		NPWPNumber:       dataTxt[4],
		KTPImage:         partImg[0],
		NPWPImage:        partImg[1],
		KTPSelfieImage:   partImg[2],
		BankName:         dataTxt[5],
		BankAccounName:   dataTxt[6],
		BankAccounNumber: dataTxt[7],
	}
	return &reqMitra, partImg, nil
}

func (e *UserUsecaseImpl) ValidUser(c *gin.Context) (*model.User, error) {
	userAuth, err := middleware.ExtractTokenAuth(c)
	if err != nil {
		return nil, err
	}
	user, err := e.userRepo.ViewById(userAuth.ID)
	if err != nil {
		return nil, err
	}
	if user.Password != userAuth.Password {
		return nil, fmt.Errorf("invalid token, please sign out and sign in again")
	}
	return user, nil
}

func (e *UserUsecaseImpl) ValidMitra(c *gin.Context) (*model.User, *model.Mitra, error) {
	user, err := e.ValidUser(c)
	if err != nil {
		return nil, nil, err
	}
	if user.Role == "user" {
		return nil, nil, fmt.Errorf("You are not allowed to add products, immediately upgrade to partners so you can sell")
	}
	mitra, err := e.userRepo.ViewMitraByUserId(int(user.ID))
	if err != nil {
		return nil, nil, fmt.Errorf("You are not yet registered as a partner")
	}
	if mitra.Status == "rejected" {
		return nil, nil, fmt.Errorf("Partner status is rejected, contact admin for more info")
	}
	if mitra.Status == "waiting for approval" {
		return nil, nil, fmt.Errorf("partner status is still under review")
	}
	return user, mitra, nil
}

func (e *UserUsecaseImpl) AdminOnly(c *gin.Context) (*model.User, error) {
	user, err := e.ValidUser(c)
	if err != nil {
		return nil, err
	}
	if user.Role != "admin" {
		return nil, fmt.Errorf("You are not an admin")
	}
	return user, nil
}

func (e *UserUsecaseImpl) ViewByMobileNumber(mobileNumber int) (*model.User, error) {
	return e.userRepo.ViewByMobileNumber(mobileNumber)
}

func (e *UserUsecaseImpl) OtpEmailVerifyUpdate(userAuth *model.UserAuth) error {
	tx := e.ongkirRepo.BeginTrans()
	userAuth.EmailVerify = "verify"
	logUser, err := e.userRepo.UpdateOTP(int(userAuth.UserID), userAuth, tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = e.CheckingUpdateUserVertifikasi(logUser, tx)
	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}

func (e *UserUsecaseImpl) OtpHandphoneVerifyUpdate(userAuth *model.UserAuth) error {
	tx := e.ongkirRepo.BeginTrans()
	userAuth.HPVerify = "verify"
	logAuth, err := e.userRepo.UpdateOTP(int(userAuth.UserID), userAuth, tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = e.CheckingUpdateUserVertifikasi(logAuth, tx)
	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}

func (e *UserUsecaseImpl) CheckingUpdateUserVertifikasi(userAuth *model.UserAuth, tx *gorm.DB) error {
	if userAuth.HPVerify == "verify" && userAuth.EmailVerify == "verify" {
		user, err := e.userRepo.ViewById(int(userAuth.UserID))
		if err != nil {
			tx.Rollback()
			return err
		}
		user.Status = "vertifiled"
		_, err = e.userRepo.UpdateUser(int(user.ID), user, tx)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return nil
}
