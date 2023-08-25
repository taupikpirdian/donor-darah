package mysql

import (
	"context"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/sirupsen/logrus"
)

func (m *mysqlNotificationRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []*domain.NotificationData, err error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			logrus.Error(errRow)
		}
	}()

	result = make([]*domain.NotificationData, 0)
	for rows.Next() {
		t := &domain.NotificationData{}
		err = rows.Scan(
			&t.Id,
			&t.UserId,
			&t.Title,
			&t.Message,
			&t.Status,
			&t.CreatedAt,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}

func (m *mysqlNotificationRepository) GetListNotification(ctx context.Context, userId int64) ([]*domain.NotificationData, error) {
	query := `SELECT id,userId,title,message,status,createdAt FROM notifications Where userId = ? ORDER BY createdAt `

	res, err := m.fetch(ctx, query, userId)
	if err != nil {
		return nil, err
	}

	return res, nil
}
