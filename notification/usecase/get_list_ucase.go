package usecase

import (
	"context"
	"fmt"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/helper"
)

func (not *notificationUsecase) GetListNotification(c context.Context, userId int64) ([]*domain.NotificationData, error) {
	ctx, cancel := context.WithTimeout(c, not.contextTimeout)
	defer cancel()
	data, err := not.notificationRepo.GetListNotification(ctx, userId)

	fmt.Println(helper.PrettyPrint(data))
	if err != nil {
		return nil, err
	}

	return data, nil
}
