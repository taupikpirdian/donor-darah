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
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}

func (m *mysqlNotificationRepository) GetListNotification(ctx context.Context) ([]*domain.NotificationData, error) {
	query := `SELECT id,name FROM notifications ORDER BY name `

	res, err := m.fetch(ctx, query)
	if err != nil {
		return nil, err
	}

	return res, nil
}
