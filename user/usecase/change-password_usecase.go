package usecase

import (
	"context"

	"github.com/bxcodec/go-clean-arch/domain"
)

func (us *userUsecase) ChangePassword(c context.Context, user *domain.User, userId int64) error {

	ctx, cancel := context.WithTimeout(c, us.contextTimeout)
	defer cancel()

	dp, err := domain.NewUser2(userId, user.Password)
	if err != nil {
		return err
	}

	_, errUser := us.userRepo.FindUserById(ctx, dp)
	if errUser != nil {
		return errUser
	}

	err = us.userRepo.ChangePassword(ctx, dp)
	if err != nil {
		return err
	}

	return nil

}
