package usecase

import (
	"context"
	"fmt"

	"github.com/bxcodec/go-clean-arch/domain"
)

func (us *userUsecase) ChangePassword(c context.Context, user *domain.User, userId int64) error {

	ctx, cancel := context.WithTimeout(c, us.contextTimeout)
	defer cancel()

	dp, err := domain.NewUser2(userId, user.Password)
	if err != nil {
		return err
	}

	dataUserDb, errUser := us.userRepo.FindUserById(ctx, dp)
	fmt.Println(user)
	fmt.Println(dataUserDb)
	if errUser != nil {
		return errUser
	}

	// err := us.userRepo.ChangePassword(ctx, dataUserDb)
	// if err != nil {
	// 	return err
	// }

	return nil

}
