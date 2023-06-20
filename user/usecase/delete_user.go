package usecase

import (
	"context"
)

func (us *userUsecase) DeleteUser(c context.Context, id string) error {
	ctx, cancel := context.WithTimeout(c, us.contextTimeout)
	defer cancel()

	err := us.userRepo.DeleteUser(ctx, id)
	if err != nil {
		return err
	}

	errProfil := us.userRepo.DeleteUserProfil(ctx, id)
	if errProfil != nil {
		return errProfil
	}

	return nil
}
