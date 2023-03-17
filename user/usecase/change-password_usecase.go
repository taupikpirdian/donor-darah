package usecase

import (
	"context"
	"fmt"

	"github.com/bxcodec/go-clean-arch/domain"
)

func (us *userUsecase) ChangePassword(c context.Context, user *domain.User, userId int64) error {

	ctx, cancel := context.WithTimeout(c, us.contextTimeout)
	defer cancel()

	dp, err := domain.NewUserp(&domain.User{Id: string(userId)})
	if err != nil {
		return err
	}
	dataUserDb, errUser := us.userRepo.FindUser(ctx, dp)
	if errUser != nil {
		return errUser
	}

	fmt.Println(dataUserDb)
	// err := us.userRepo.ChangePassword(ctx,dataUserDb)
	// if err != nil {
	// 	return err
	// }

	return nil

}
