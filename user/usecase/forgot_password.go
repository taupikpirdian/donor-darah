package usecase

import (
	"context"

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
	bodyContent := "Hello, Anda telah melakukan penggantian password, berikut password baru Anda: <b>" + newPassword + "</b>"
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", viper.GetString(`smtp.CONFIG_SENDER_NAME`))
	mailer.SetHeader("To", fUser.GetEmailOnUser())
	mailer.SetHeader("Subject", "Forgot Password")
	mailer.SetBody("text/html", bodyContent)

	errMail := us.serviceMail.SendEmail(mailer)
	if errMail != nil {
		return errMail
	}

	return nil
}
