package mysql

import (
	"context"
	"time"
)

func (m *mysqlDonorRepository) StoreStock(ctx context.Context, unitId int64, title string) (int64, error) {

	query := `INSERT blood_stock SET unitId=?, title=?, updatedAt=?, createdAt=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return 0, err
	}

	res, err := stmt.ExecContext(ctx, unitId, title, time.Now(), time.Now())
	if err != nil {
		return 0, err
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastID, nil
}
