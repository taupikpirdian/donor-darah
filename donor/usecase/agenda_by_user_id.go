package usecase

import (
	"context"
	"fmt"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/helper"
)

func (dus *donorUsecase) RegisterByUserId(c context.Context, id int64) ([]*domain.DonorRegisterDTO, error) {
	ctx, cancel := context.WithTimeout(c, dus.contextTimeout)
	defer cancel()

	data, errR := dus.donorRepo.AgendaByUserId(ctx, id)
	if errR != nil {
		return nil, errR
	}

	fmt.Println(helper.PrettyPrint(data))

	return data, nil
}
