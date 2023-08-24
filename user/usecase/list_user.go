package usecase

import (
	"context"
	"strconv"

	"github.com/bxcodec/go-clean-arch/domain"
)

func (us *userUsecase) ListUser(c context.Context) ([]*domain.User, error) {
	ctx, cancel := context.WithTimeout(c, us.contextTimeout)
	defer cancel()

	datas, err := us.userRepo.GetListUser(ctx)
	if err != nil {
		return nil, err
	}

	for _, value := range datas {
		userId, err := strconv.ParseInt(value.Id, 10, 64)
		if err != nil {
			return nil, err
		}
		profile, err := us.userRepo.GetProfileFull(ctx, userId)
		if err != nil {
			return nil, err
		}
		job, errJob := us.userRepo.GetJobById(ctx, profile.JobId.String)
		if errJob != nil {
			return nil, errJob
		}
		unit, errUnit := us.userRepo.GetUnitById(ctx, profile.UnitId.String)
		if errUnit != nil {
			return nil, errUnit
		}
		subDistrict, _ := us.userRepo.GetSubDistrictById(ctx, profile.SubDistrictId)
		villages, _ := us.userRepo.GetVillageById(ctx, profile.VillageId)

		value.SetUserList(
			profile,
			job,
			unit,
			subDistrict,
			villages,
		)
	}

	return datas, nil
}
