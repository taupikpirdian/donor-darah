package mysql

import (
	"context"
	"time"
)

func (m *mysqlNotificationRepository) CreateNotification(ctx context.Context, title string, msg string, userId int64) error {
	query := `INSERT notifications SET title=? , message=? , status=? , userId=?, updatedAt=? , createdAt=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	_, errProses := stmt.ExecContext(ctx, title, msg, "UNREAD", userId, time.Now(), time.Now())
	if errProses != nil {
		return errProses
	}

	return nil
}
