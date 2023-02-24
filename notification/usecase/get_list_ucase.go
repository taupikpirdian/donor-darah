package usecase

import (
	"context"

	"github.com/bxcodec/go-clean-arch/domain"
)

func (not *notificationUsecase) GetListNotification(c context.Context) ([]*domain.NotificationData, error) {
	ctx, cancel := context.WithTimeout(c, not.contextTimeout)
	defer cancel()

	data, err := not.notificationRepo.GetListNotification(ctx)
	if err != nil {
		return nil, err
	}

	return data, nil
}
