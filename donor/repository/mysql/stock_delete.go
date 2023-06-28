package mysql

import "context"

func (m *mysqlDonorRepository) StockDelete(ctx context.Context, id int64) error {

	query := `DELETE FROM blood_stock WHERE id=?`
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
