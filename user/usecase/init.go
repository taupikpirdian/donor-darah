package usecase

import (
	"time"

	"github.com/bxcodec/go-clean-arch/cfg"
	"github.com/bxcodec/go-clean-arch/domain"
)

type userUsecase struct {
	userRepo       domain.UserRepository
	serviceMail    domain.MailService
	contextTimeout time.Duration
	donorRepo      domain.DonorRepository
	cfg            cfg.Config
}

// NewUserUsecase will create new an userUsecase object representation of domain.UserUsecase interface
func NewUserUsecase(a domain.UserRepository, b domain.MailService, timeout time.Duration, c domain.DonorRepository, cfg cfg.Config) domain.UserUsecase {
	return &userUsecase{
		userRepo:       a,
		serviceMail:    b,
		contextTimeout: timeout,
		donorRepo:      c,
		cfg:            cfg,
	}
}
