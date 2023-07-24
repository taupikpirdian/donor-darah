package usecase

import (
	"context"
	"time"

	"github.com/bxcodec/go-clean-arch/domain"
)

func (us *userUsecase) GetProfile(c context.Context, userId int64) (*domain.Profile, error) {
	ctx, cancel := context.WithTimeout(c, us.contextTimeout)
	defer cancel()

	/*
		get data profile
		get and count data donor_registers, status == DONE
		get first data order By asc, status == OPEN
	*/

	data, err := us.userRepo.GetProfile(ctx, userId)
	if err != nil {
		return nil, err
	}

	profile, err := us.userRepo.GetProfileFull(ctx, userId)
	if err != nil {
		return nil, err
	}

	dataDonorRegis, errDr := us.donorRepo.ListRiwayatByStatus(ctx, userId, "DONE")
	if errDr != nil {
		return nil, errDr
	}

	dataNextRegis, errNr := us.donorRepo.NextDonorByStatus(ctx, userId, "OPEN")
	if errNr != nil {
		return nil, errNr
	}

	var nextDateDonor time.Time
	if dataNextRegis != nil {
		nextDateDonor = dataNextRegis.DonorSchedulle.Date
	}
	dataLastsRegis, errLr := us.donorRepo.LastDonorByStatus(ctx, userId, "DONE")
	if errNr != nil {
		return nil, errLr
	}
	var lastDateDonor time.Time
	if dataLastsRegis != nil {
		lastDateDonor = dataLastsRegis.DonorSchedulle.Date
	}
	result := domain.NewProfileV2(data, profile, len(dataDonorRegis), nextDateDonor, lastDateDonor)
	return result, nil
}
