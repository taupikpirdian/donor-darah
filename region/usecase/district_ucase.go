package usecase

import (
	"context"
	"time"

	"github.com/bxcodec/go-clean-arch/domain"
)

type regionUsecase struct {
	regionRepo     domain.RegionRepository
	contextTimeout time.Duration
}

func NewRegionUsecase(a domain.RegionRepository, timeout time.Duration) domain.RegionUsecase {
	return &regionUsecase{
		regionRepo:     a,
		contextTimeout: timeout,
	}
}

func (reg *regionUsecase) GetDistrict(c context.Context) ([]*domain.DistrictData, error) {
	ctx, cancel := context.WithTimeout(c, reg.contextTimeout)
	defer cancel()

	data, err := reg.regionRepo.GetDistrict(ctx)
	if err != nil {
		return nil, err
	}

	return data, nil
}
