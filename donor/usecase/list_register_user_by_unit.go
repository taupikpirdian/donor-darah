package usecase

import (
	"context"
	"time"

	"github.com/bxcodec/go-clean-arch/domain"
)

func (dus *donorUsecase) ListRegisterUserByUnit(c context.Context, unitId int, date string) ([]*domain.DonorRegisterDTO, error) {
	ctx, cancel := context.WithTimeout(c, dus.contextTimeout)
	defer cancel()

	datas, errR := dus.donorRepo.ListDonorRegisterByUnit(ctx, unitId, date)
	if errR != nil {
		return nil, errR
	}

	// jumlah donor and last donor
	for _, value := range datas {
		dataLastsRegis, errLr := dus.donorRepo.LastDonorByStatus(ctx, value.UserId, "DONE")
		if errLr != nil && errLr.Error() != "your requested Item is not found" {
			return nil, errLr
		}
		var lastDateDonor time.Time
		if dataLastsRegis != nil {
			lastDateDonor = dataLastsRegis.DonorSchedulle.Date
		}
		value.LastDonor = lastDateDonor

		dataDonorRegis, errDr := dus.donorRepo.ListRiwayatByStatus(ctx, value.UserId, "DONE")
		if errDr != nil {
			return nil, errDr
		}
		value.TotalDonor = len(dataDonorRegis)

		// Define the layout for the input string
		layout := "2006-01-02T15:04:05-07:00"
		// Parse the input string into a time.Time object
		t, err := time.Parse(layout, value.User.DateOfBirth)
		if err != nil {
			return nil, err
		}
		value.User.DateOfBirth = t.Format("2006-01-02")
	}

	return datas, nil
}
