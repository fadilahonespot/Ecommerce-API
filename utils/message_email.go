package utils

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

func WelcomeEmail(subject string, body string, email string, link string, asset string) {
	var resFile2 string
	if email != "" && link != "" {
		txt, err := ReadTemplateHtml(body)
		if err != nil {
			fmt.Printf("[Utils.WelcomeEmail] error message %v \n", err)
			return
		}
		strFile := string(txt)
		resFile1 := strings.Replace(strFile, "%%link", link, -1)
		resFile2 = strings.Replace(resFile1, "%%email", email, -1)
	}
	emailProcess(email, subject, asset, resFile2)
}

func ResetPassEmail(email string, subjek string, body string, link string) {
	txt, err := ReadTemplateHtml(body)
	if err != nil {
		fmt.Printf("[Utils.ResetPassEmail] error ioutil read file %v \n", err)
		return
	}
	repFile := strings.Replace(string(txt), "%%link", link, -1)
	emailProcess(email, subjek, "", repFile)
}

func NewsLetterEmail(email string, subject string, body string, article string) {
	txt, err := ReadTemplateHtml(body)
	if err != nil {
		fmt.Printf("[Utils.NewsLetterEmail] error ioutil read file %v \n", err)
		return
	}
	repFile := strings.Replace(string(txt), "%%artikel", article, -1)
	emailProcess(email, subject, "", repFile)
}

func ResetPassEmailSucces(email string, subject string, body string, pass string) {
	txt, err := ReadTemplateHtml(body)
	if err != nil {
		fmt.Printf("[Utils.ResetPassEmail] error ioutil read file %v \n", err)
		return
	}
	repFile := strings.Replace(string(txt), "%%password", pass, -1)
	repFile2 := strings.Replace(repFile, "%%email", email, -1)
	emailProcess(email, subject, "", repFile2)
}

func emailProcess(email string, subject string, asset string, body string) {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", viper.GetString("email.CONFIG_EMAIL"))
	mailer.SetHeader("To", email)
	// mailer.SetAddressHeader("Cc", "tralalala@gmail.com", "Tra Lala La")
	mailer.SetHeader("Subject", subject)
	if asset != "" {
		mailer.Embed(asset)
	}
	// mailer.Attach("utils/indonesia.jpg")
	if body != "" {
		mailer.SetBody("text/html", body)
	}

	dialer := gomail.NewDialer(
		viper.GetString("email.CONFIG_SMTP_HOST"),
		viper.GetInt("email.CONFIG_SMTP_PORT"),
		viper.GetString("email.CONFIG_EMAIL"),
		viper.GetString("email.CONFIG_PASSWORD"),
	)

	if viper.GetString("email.switch") == "on" {
		err := dialer.DialAndSend(mailer)
		if err != nil {
			fmt.Printf("[welcomeEmail] error configure %v \n", err)
			panic(err)
		}
	}
}

func ReadTemplateHtml(part string) ([]byte, error) {
	temp, err := ioutil.ReadFile(part)
	if err != nil {
		fmt.Printf("[Utils.ReadTemplateSuccesEmailVerification] Error Message : %v \n", err)
		return nil, fmt.Errorf("failed read data html")
	}
	return temp, nil
}
