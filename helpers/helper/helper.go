package helper

import (
	"fmt"
	"strconv"

	"github.com/alfisar/jastip-import/domain"

	"gopkg.in/gomail.v2"
)

func SendEmailOTP(data domain.SMTP, email string, fullname string, OTP string) (err error) {
	body := "Hello " + fullname + " , thank u for the registration, so this the OTP for confirmation your registration, OTP : " + OTP

	_port, _ := strconv.Atoi(data.Port)

	mailer := data.Mailer
	mailer.SetHeader("From", data.From)
	mailer.SetHeader("To", email)
	mailer.SetHeader("Subject", "Registration Jastip.in")
	mailer.SetBody("text/html", body)

	dialer := gomail.NewDialer(
		data.Host,
		_port,
		data.User,
		data.Pass,
	)

	err = dialer.DialAndSend(mailer)

	if err != nil {
		fmt.Println("Sending email is Error : " + err.Error())
		return
	}

	fmt.Println("Sending email is success")
	return
}
