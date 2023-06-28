package mysql

import (
	"context"
)

func (m *mysqlDonorRepository) SchedulleDelete(ctx context.Context, id int64) error {

	query := `DELETE FROM donor_schedulle WHERE id=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	_, errExec := stmt.ExecContext(ctx, id)
	if errExec != nil {
		return errExec
	}

	return nil
}
