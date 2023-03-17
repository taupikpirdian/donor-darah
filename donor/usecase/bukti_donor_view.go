package usecase

import (
	"context"

	"github.com/bxcodec/go-clean-arch/domain"
)

func (dus *donorUsecase) UploadBuktiView(c context.Context, id int64) (*domain.DonorRegisterDTO, error) {
	ctx, cancel := context.WithTimeout(c, dus.contextTimeout)
	defer cancel()

	data, errR := dus.donorRepo.UploadBuktiView(ctx, id)
	if errR != nil {
		return nil, errR
	}

	return data, nil
}
