package usecase

import (
	"context"

	"github.com/bxcodec/go-clean-arch/domain"
)

func (dus *donorUsecase) ListSchedulle(c context.Context, unitId int64) ([]*domain.DonorSchedulleDTO, error) {
	ctx, cancel := context.WithTimeout(c, dus.contextTimeout)
	defer cancel()

	datas, errR := dus.donorRepo.ListSchedulle(ctx, unitId)
	if errR != nil {
		return nil, errR
	}

	return datas, nil
}
