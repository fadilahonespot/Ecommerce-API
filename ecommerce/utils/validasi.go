package utils

import (
	"ecommerce/model"
	"fmt"
	"strings"
)

func LogginValidation(email string, pass string) error {
	if email == "" {
		return fmt.Errorf("email cannot be empty")
	}

	if pass == "" {
		return fmt.Errorf("password cannot be empty")
	}

	isExist := strings.Contains(email, "@")
	if isExist == false {
		return fmt.Errorf("Email address must contain symbol @")
	}

	isSpace := strings.Contains(email, " ")
	if isSpace == true {
		return fmt.Errorf("Email addresses cannot contain space")
	}
	return nil
}

func RegisterValidasi(user *model.User) error {
	if user.Alamat == "" || user.Email == "" || user.IDKota == 0 || user.Nama == "" || user.NoHP == 0 || user.Password == "" {
		return fmt.Errorf("column cannot be empty")
	}
	if user.ID != 0 {
		return fmt.Errorf("inputs are not permitted")
	}
	if strings.Contains(user.Email, "@") == false {
		return fmt.Errorf("invalid email format")
	}
	if len(strings.Split(user.Password, "")) <= 8 {
		return fmt.Errorf("password minimal 8 characters")
	}
	return nil
}

