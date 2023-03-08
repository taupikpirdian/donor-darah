package usecase

import (
	"time"

	"github.com/bxcodec/go-clean-arch/domain"
)

type notificationUsecase struct {
	notificationRepo domain.NotificationRepository
	contextTimeout   time.Duration
}

// NewUserUsecase will create new an userUsecase object representation of domain.UserUsecase interface
func NewNotificationUsecase(a domain.NotificationRepository, timeout time.Duration) domain.NotificationUsecase {
	return &notificationUsecase{
		notificationRepo: a,
		contextTimeout:   timeout,
	}
}
