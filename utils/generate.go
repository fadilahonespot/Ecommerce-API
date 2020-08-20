package utils

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"github.com/xlzd/gotp"
)

func GenerateToken(id int, pass string) (string, error) {
	secret := viper.GetString("security.secret")
	claims := jwt.MapClaims{}
	
	claims["user_id"] = id
	claims["pass"] = pass 

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		fmt.Printf("[Utils.GenerateToken] Error to signature generate token, %v \n", err)
		return "", fmt.Errorf("Failed generate token")
	}
	return tokenString, nil
}

func GenerateOTP() string {
	secret := gotp.RandomSecret(16)
	generate := gotp.NewDefaultTOTP(secret)
	otp := generate.At(1)
	return otp
}

func GenerateRandomNumber() int {
	min := 100
	max := 999
	rand.Seed(time.Now().Unix())
    return rand.Intn(max-min) + min
}

func GenerateNoTransaction(number uint) int {
	var transactionNum = 400005878374
	result := uint(transactionNum) + number
	return int(result)
}