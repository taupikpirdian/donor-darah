package mysql

import (
	"context"

	"github.com/bxcodec/go-clean-arch/domain"
)

func (m *mysqlNotificationRepository) GetSingleNotification(ctx context.Context, id int64) (res *domain.NotificationData, err error) {
	query := `SELECT id,userId,title,message,createdAt,updatedAt FROM notifications WHERE id = ? `

	list, err := m.fetch(ctx, query, id)
	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, domain.ErrNotFound
	}

	return
}
