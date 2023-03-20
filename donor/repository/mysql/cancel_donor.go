package mysql

import (
	"context"
	"fmt"
)

func (m *mysqlDonorRepository) CancelDonor(ctx context.Context, id int64) (err error) {
	query := `UPDATE donor_registers set isApprove=? WHERE ID = ?`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}
	res, err := stmt.ExecContext(ctx, "2", id)
	if err != nil {
		return
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return
	}
	if affect != 1 {
		err = fmt.Errorf("weird  Behavior. Total Affected: %d", affect)
		return
	}

	return
}
