package usecase

import (
	"context"

	"github.com/bxcodec/go-clean-arch/domain"
)

func (dus *donorUsecase) ListStock(c context.Context, unitId int64) ([]*domain.DonorStock, error) {
	ctx, cancel := context.WithTimeout(c, dus.contextTimeout)
	defer cancel()

	datas, errR := dus.donorRepo.ListStock(ctx, unitId)
	if errR != nil {
		return nil, errR
	}

	return datas, nil
}
