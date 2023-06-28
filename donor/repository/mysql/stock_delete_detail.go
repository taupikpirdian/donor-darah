package mysql

import "context"

func (m *mysqlDonorRepository) StockDeleteDetail(ctx context.Context, stockId int64) error {

	query := `DELETE FROM blood_stock_detail WHERE blood_stock_id=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	_, errExec := stmt.ExecContext(ctx, stockId)
	if errExec != nil {
		return errExec
	}

	return nil
}
