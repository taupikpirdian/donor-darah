package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/bxcodec/go-clean-arch/domain"
)

func (dus *donorUsecase) Reschedule(c context.Context, id int64, dto *domain.DonorSchedulleDTO) error {
	ctx, cancel := context.WithTimeout(c, dus.contextTimeout)
	defer cancel()

	/*
		find data
	*/
	data, errR := dus.donorRepo.FindSchedule(ctx, dto)
	fmt.Println(errR)
	if errR != nil {
		return errR
	}
	if data == nil {
		return errors.New("data not found")
	}

	/*
		build entity
	*/
	newSchedule, errSch := domain.NewDonorSchedule(*data)
	if errSch != nil {
		return errSch
	}

	errReschedule := dus.donorRepo.Reschedule(ctx, id, newSchedule)
	if errReschedule != nil {
		return errReschedule
	}

	return nil
}
