package repo

import (
	"ecommerce/model"
	"ecommerce/user"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

type UserRepoImpl struct {
	DB *gorm.DB
}

func CreateUserRepo(DB *gorm.DB) user.UserRepo {
	return &UserRepoImpl{DB}
}

func (e *UserRepoImpl) ViewAll()(*[]model.User, error) {
	var user []model.User
	err := e.DB.Find(&user).Error
	if err != nil {
		fmt.Printf("[UserRepoImpl.ViewAll] Error execute query %v \n", err)
		return nil, fmt.Errorf("failed view all data buyer")
	}
	return &user, nil
}

func (e *UserRepoImpl) InsertUser(user *model.User, tx *gorm.DB) (*model.User, error) {
	err := tx.Save(&user).Error
	if err != nil {
		fmt.Printf("[UserRepoImpl.InsertUser] Error execute query %v \n", err)
		return nil, fmt.Errorf("failed add data user")
	}
	return user, nil
}

func (e *UserRepoImpl) ViewByEmail(email string) (*model.User, error) {
	var user = model.User{}
	err := e.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		fmt.Printf("[UserRepoImpl.ViewByEmail] Error execute query %v \n", err)
		return nil, fmt.Errorf("email is not exsis")
	}
	return &user, nil
}

func (e *UserRepoImpl) ViewByMobileNumber(mobileNumber int)(*model.User, error) {
	var user = model.User{}
	err := e.DB.Where("no_hp = ?", mobileNumber).First(&user).Error
	if err != nil {
		fmt.Printf("[UserRepoImpl.ViewByMobileNumber] error execute query %v \n", err)
		return nil, fmt.Errorf("mobile number is not exsis")
	}
	return &user, nil
}

func (e *UserRepoImpl) UpdateUser(id int, user *model.User, tx *gorm.DB)(*model.User, error) {
	var newUser = model.User{}
	err := tx.Table("user").Where("ID = ?", id).First(&newUser).Update(&user).Error
	if err != nil {
		fmt.Printf("[UserRepoImpl.UpdateData] Error execute query %v \n", err)
		return nil, fmt.Errorf("failed update data")
	}
	return &newUser, nil
}

func (e *UserRepoImpl) ViewById(id int)(*model.User, error) {
	var user = model.User{}
	err := e.DB.Table("user").Where("ID = ?", id).First(&user).Error
	if err != nil {
		fmt.Printf("[UserRepoImpl.ViewById] Error execute query %v \n", err)
		return nil, fmt.Errorf("id is not exist")
	}
	return &user, nil
} 

func (e *UserRepoImpl) InsertOTP(otp *model.UserAuth) (*model.UserAuth, error) {
	err := e.DB.Save(&otp).Error
	if err != nil {
		fmt.Printf("[UserHandler.InsertOTP] Error execute query %v \n", err)
		return nil, fmt.Errorf("failed input data")
	}
	return otp, nil
}

func (e *UserRepoImpl) UpdateOTP(idUser int, otp *model.UserAuth, tx *gorm.DB) (*model.UserAuth, error) {
	var upOTP = model.UserAuth{}
	err := tx.Table("user_auth").Where("user_id = ?", idUser).First(&upOTP).Update(&otp).Error
	if err != nil {
		fmt.Printf("[UserRepoImpl.UpdateOTP] error execute query %v \n", err)
		return nil, fmt.Errorf("failed to update data")
	}
	return &upOTP, nil
}

func (e *UserRepoImpl) OtpByIDUser(id int)(*model.UserAuth, error) {
	var otp = model.UserAuth{}
	err := e.DB.Table("user_auth").Where("user_id = ?", id).First(&otp).Error
	if err != nil {
		fmt.Printf("[UserRepoImpl.OtpByIDUser] error execute query %v \n", err)
		return nil, fmt.Errorf("id is not exist")
	}
	return &otp, nil
}

func (e *UserRepoImpl) InsertNewsletter(news *model.Newletter) (*model.Newletter, error) {
	err := e.DB.Save(&news).Error
	if err != nil {
		fmt.Printf("[userRepo.InserNewsletter] error execute query %v \n", err)
		return nil, fmt.Errorf("failed add data newsletter")
	}
	return news, nil
}

func (e *UserRepoImpl) ViewAllNewsletter()(*[]model.Newletter, error) {
	var newData []model.Newletter
	err := e.DB.Find(&newData).Error
	if err != nil {
		fmt.Printf("[userRepo.ViewAllNewsletter] Error execute query %v \n", err)
		return nil, fmt.Errorf("failed view data newsletter")
	}
	return &newData, nil
}

func (e *UserRepoImpl) SendSmsOTP(otp string, hp int) error {
	_, err := http.NewRequest("GET", viper.GetString("medansms.url") + "sms_api.php?action=kirim_sms&email=" + 
	viper.GetString("medansms.email") + "&passkey=" + viper.GetString("medansms.token") + "&no_tujuan=" + 
	strconv.Itoa(hp) + "&pesan=" + viper.GetString("message_otp") + otp, nil)
	if err != nil {
		fmt.Printf("[UserRepoImpl.SendSmsOTP] Error execute url, %v \n", err)
		return fmt.Errorf("Failed send otp message")
	}
	return nil
}

func (e *UserRepoImpl) InsertMitra(mitra *model.Mitra)(*model.Mitra, error) {
	err := e.DB.Save(&mitra).Error
	if err != nil {
		fmt.Printf("[userRepoImpl.InsertMitra] Error execute query %v \n", err)
		return nil, fmt.Errorf("failed insert data mitra")
	}
	return mitra, nil
}

func (e *UserRepoImpl) ViewAllMitra()(*[]model.Mitra, error) {
	var mitra []model.Mitra
	err := e.DB.Find(&mitra).Error
	if err != nil {
		fmt.Printf("[userRepoImpl.ViewAllMitra] error execute query %v \n", err)
		return nil, fmt.Errorf("failed view all data mitra")
	}
	return &mitra, err
}

func (e *UserRepoImpl) UpdateMitraById(id int, mitra *model.Mitra, tx *gorm.DB)(*model.Mitra, error) {
	var upMitra = model.Mitra{}
	err := tx.Table("mitra").Where("ID = ?", id).First(&upMitra).Update(&mitra).Error
	if err != nil {
		fmt.Printf("[userRepoImpl.UpdateMitra] error execute query %v \n", err)
		return nil, fmt.Errorf("failed update data mitra")
	}
	return &upMitra, nil
}

func (e *UserRepoImpl) ViewMitraByUserId(id int)(*model.Mitra, error) {
	var mitra = model.Mitra{}
	err := e.DB.Table("mitra").Where("user_id = ?", id).First(&mitra).Error
	if err != nil {
		fmt.Printf("[userRepoImpl.ViewMitraByUserId] error execute query %v \n", err)
		return nil, fmt.Errorf("id user is not exist")
	}
	return &mitra, nil
}

func (e *UserRepoImpl) ViewMitraById(id int)(*model.Mitra, error) {
	var mitra = model.Mitra{}
	err := e.DB.Table("mitra").Where("ID = ?", id).First(&mitra).Error
	if err != nil {
		fmt.Printf("[userRepoImpl.ViewMitraById] error execute query %v \n", err)
		return nil, fmt.Errorf("id mitra is not exist")
	}
	return &mitra, nil
}
