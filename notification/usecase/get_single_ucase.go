package usecase

import (
	"context"

	"github.com/bxcodec/go-clean-arch/domain"
)

func (not *notificationUsecase) GetSingleNotification(c context.Context, id int64, userId int64) (*domain.NotificationData, error) {
	ctx, cancel := context.WithTimeout(c, not.contextTimeout)
	defer cancel()

	// update read
	errUpdate := not.notificationRepo.UpdateSingleNotification(ctx, id)
	if errUpdate != nil {
		return nil, errUpdate
	}

	data, err := not.notificationRepo.GetSingleNotification(ctx, id, userId)
	if err != nil {
		return nil, err
	}

	return data, nil
}
