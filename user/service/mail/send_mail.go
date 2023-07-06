package send_mail

import (
	"log"

	"gopkg.in/gomail.v2"
)

func (sm *serviceMail) SendEmail(mailer *gomail.Message) error {
	dialer := gomail.NewDialer(
		sm.cfg.CONFIG_SMTP_HOST,
		sm.cfg.CONFIG_SMTP_PORT,
		sm.cfg.CONFIG_AUTH_EMAIL,
		sm.cfg.CONFIG_AUTH_PASSWORD,
	)

	errMail := dialer.DialAndSend(mailer)

	if errMail != nil {
		log.Fatal(errMail.Error())
		return errMail
	}

	return nil
}
