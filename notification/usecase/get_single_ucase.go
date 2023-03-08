package usecase

import (
	"context"

	"github.com/bxcodec/go-clean-arch/domain"
)

func (not *notificationUsecase) GetSingleNotification(c context.Context, id int64) (*domain.NotificationData, error) {
	ctx, cancel := context.WithTimeout(c, not.contextTimeout)
	defer cancel()

	data, err := not.notificationRepo.GetSingleNotification(ctx, id)
	if err != nil {
		return nil, err
	}

	return data, nil
}
