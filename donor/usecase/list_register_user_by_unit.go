package usecase

import (
	"context"

	"github.com/bxcodec/go-clean-arch/domain"
)

func (dus *donorUsecase) ListRegisterUserByUnit(c context.Context, unitId int, date string) ([]*domain.DonorRegisterDTO, error) {
	ctx, cancel := context.WithTimeout(c, dus.contextTimeout)
	defer cancel()

	datas, errR := dus.donorRepo.ListDonorRegisterByUnit(ctx, unitId, date)
	if errR != nil {
		return nil, errR
	}

	return datas, nil
}
