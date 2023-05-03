package domain

import "gopkg.in/gomail.v2"

type MailService interface {
	SendEmail(mailer *gomail.Message) error
}
