package usecase

import (
	"context"

	"github.com/bxcodec/go-clean-arch/domain"
)

func (dus *donorUsecase) ListRiwayat(c context.Context, userId int64) ([]*domain.DonorRegisterDTO, error) {
	ctx, cancel := context.WithTimeout(c, dus.contextTimeout)
	defer cancel()

	datas, errR := dus.donorRepo.ListRiwayat(ctx, userId)
	if errR != nil {
		return nil, errR
	}

	return datas, nil
}
