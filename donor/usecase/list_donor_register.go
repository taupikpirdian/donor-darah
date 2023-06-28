package usecase

import (
	"context"

	"github.com/bxcodec/go-clean-arch/domain"
)

func (dus *donorUsecase) ListDonorRegister(c context.Context) ([]*domain.DonorRegisterDTO, error) {
	ctx, cancel := context.WithTimeout(c, dus.contextTimeout)
	defer cancel()

	datas, errR := dus.donorRepo.ListDonorRegister(ctx)
	if errR != nil {
		return nil, errR
	}

	// add answer
	for _, data := range datas {
		answer, errA := dus.donorRepo.ListAnswer(ctx, data.Id)
		if errA != nil {
			return nil, errA
		}
		data.DonorRegisterQuestionerDTO = answer
	}
	return datas, nil
}
