package usecase

import (
	"context"

	"github.com/bxcodec/go-clean-arch/domain"
)

func (dus *donorUsecase) Card(c context.Context, userId int64) (*domain.Card, error) {
	ctx, cancel := context.WithTimeout(c, dus.contextTimeout)
	defer cancel()

	data, errR := dus.donorRepo.GetCard(ctx, userId)
	if errR != nil {
		return nil, errR
	}

	return data, nil
}
