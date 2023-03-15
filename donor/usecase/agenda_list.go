package usecase

import (
	"context"

	"github.com/bxcodec/go-clean-arch/domain"
)

func (dus *donorUsecase) ListAgenda(c context.Context, userId int64) ([]*domain.DonorRegister, error) {
	return []*domain.DonorRegister{}, nil
}
