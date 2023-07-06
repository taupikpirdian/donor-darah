package usecase

import (
	"time"

	"github.com/bxcodec/go-clean-arch/cfg"
	"github.com/bxcodec/go-clean-arch/domain"
)

type regionUsecase struct {
	regionRepo     domain.RegionRepository
	contextTimeout time.Duration
	cfg            cfg.Config
}

func NewRegionUsecase(a domain.RegionRepository, timeout time.Duration, cfg cfg.Config) domain.RegionUsecase {
	return &regionUsecase{
		regionRepo:     a,
		contextTimeout: timeout,
		cfg:            cfg,
	}
}
