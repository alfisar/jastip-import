package domain

import "gopkg.in/gomail.v2"

type SMTP struct {
	Host   string
	Port   string
	User   string
	Pass   string
	From   string
	Mailer *gomail.Message
}
