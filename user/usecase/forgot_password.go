package usecase

import (
	"context"
	"log"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

func (us *userUsecase) ForgotPassword(c context.Context, user *domain.User) error {
	ctx, cancel := context.WithTimeout(c, us.contextTimeout)
	defer cancel()

	fUser, errUser := us.userRepo.FindUserByEmail(ctx, user.Email)
	if errUser != nil {
		return errUser
	}

	// set new password
	newPassword := domain.GenerateCodeStringLen(8)
	fUser.SetPasswordNew(fUser, newPassword)

	err := us.userRepo.ChangePassword(ctx, fUser)
	if err != nil {
		return err
	}

	// send email after change password
	bodyContent := "Hello, Anda telah melakukan penggantian password, berkut passwordnya: <b>" + newPassword + "</b>"
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", viper.GetString(`smtp.CONFIG_SENDER_NAME`))
	mailer.SetHeader("To", fUser.GetEmailOnUser())
	mailer.SetHeader("Subject", "Forgot Password")
	mailer.SetBody("text/html", bodyContent)

	dialer := gomail.NewDialer(
		viper.GetString(`smtp.CONFIG_SMTP_HOST`),
		viper.GetInt(`smtp.CONFIG_SMTP_PORT`),
		viper.GetString(`smtp.CONFIG_AUTH_EMAIL`),
		viper.GetString(`smtp.CONFIG_AUTH_PASSWORD`),
	)

	errMail := dialer.DialAndSend(mailer)
	if errMail != nil {
		log.Fatal(err.Error())
		return errMail
	}

	return nil
}
