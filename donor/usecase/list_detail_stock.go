package usecase

import (
	"context"

	"github.com/bxcodec/go-clean-arch/domain"
)

func (dus *donorUsecase) ListDetailStock(c context.Context, stockId int64) ([]*domain.DonorDetailStock, error) {
	ctx, cancel := context.WithTimeout(c, dus.contextTimeout)
	defer cancel()

	datas, errR := dus.donorRepo.ListDetailStock(ctx, stockId)
	if errR != nil {
		return nil, errR
	}

	return datas, nil
}
