package mysql

import (
	"context"
	"fmt"
)

func (m *mysqlDonorRepository) UploadBukti(ctx context.Context, id int64, path string) (err error) {
	query := `UPDATE donor_registers set donorProof=?, status=? WHERE ID = ?`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}
	res, err := stmt.ExecContext(ctx, path, "DONE", id)
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
