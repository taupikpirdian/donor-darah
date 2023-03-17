package usecase

import (
	"time"

	"github.com/bxcodec/go-clean-arch/domain"
)

type donorUsecase struct {
	donorRepo      domain.DonorRepository
	contextTimeout time.Duration
}

// NewUserUsecase will create new an userUsecase object representation of domain.UserUsecase interface
func NewDonorUsecase(a domain.DonorRepository, timeout time.Duration) domain.DonorUsecase {
	return &donorUsecase{
		donorRepo:      a,
		contextTimeout: timeout,
	}
}
