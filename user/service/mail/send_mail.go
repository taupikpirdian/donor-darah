package send_mail

import (
	"log"

	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

func (sm *serviceMail) SendEmail(mailer *gomail.Message) error {
	dialer := gomail.NewDialer(
		viper.GetString(`smtp.CONFIG_SMTP_HOST`),
		viper.GetInt(`smtp.CONFIG_SMTP_PORT`),
		viper.GetString(`smtp.CONFIG_AUTH_EMAIL`),
		viper.GetString(`smtp.CONFIG_AUTH_PASSWORD`),
	)

	errMail := dialer.DialAndSend(mailer)
	if errMail != nil {
		log.Fatal(errMail.Error())
		return errMail
	}

	return nil
}
