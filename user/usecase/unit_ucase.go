package usecase

import (
	"context"

	"github.com/bxcodec/go-clean-arch/domain"
)

func (us *userUsecase) GetUnit(c context.Context) ([]*domain.UnitDTO, error) {
	ctx, cancel := context.WithTimeout(c, us.contextTimeout)
	defer cancel()

	datas, err := us.userRepo.GetUnit(ctx)
	if err != nil {
		return nil, err
	}

	return datas, nil
}
