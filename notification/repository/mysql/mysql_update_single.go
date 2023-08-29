package mysql

import (
	"context"
	"time"
)

func (m *mysqlNotificationRepository) UpdateSingleNotification(ctx context.Context, id int64) error {
	query := `UPDATE notifications SET status=?, updatedAt=? WHERE id=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	_, errExec := stmt.ExecContext(ctx, 1, time.Now(), id)
	if errExec != nil {
		return errExec
	}

	return nil
}
