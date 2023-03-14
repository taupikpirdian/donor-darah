package usecase

import (
	"context"

	"github.com/bxcodec/go-clean-arch/domain"
)

func (not *donorUsecase) DonorRegister(c context.Context, userId int64, dto *domain.DonorRegisterDTO) error {
	return nil
}
