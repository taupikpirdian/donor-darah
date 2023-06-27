package usecase

import (
	"time"

	"github.com/bxcodec/go-clean-arch/cfg"
	"github.com/bxcodec/go-clean-arch/domain"
)

type donorUsecase struct {
	donorRepo      domain.DonorRepository
	contextTimeout time.Duration
	cfg            cfg.Config
}

// NewUserUsecase will create new an userUsecase object representation of domain.UserUsecase interface
func NewDonorUsecase(a domain.DonorRepository, timeout time.Duration, cfg cfg.Config) domain.DonorUsecase {
	return &donorUsecase{
		donorRepo:      a,
		contextTimeout: timeout,
		cfg:            cfg,
	}
}
