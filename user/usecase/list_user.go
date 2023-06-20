package usecase

import (
	"context"

	"github.com/bxcodec/go-clean-arch/domain"
)

func (us *userUsecase) ListUser(c context.Context) ([]*domain.User, error) {
	ctx, cancel := context.WithTimeout(c, us.contextTimeout)
	defer cancel()

	datas, err := us.userRepo.GetListUser(ctx)
	if err != nil {
		return nil, err
	}

	return datas, nil
}
