package mysql

import (
	"context"
	"fmt"

	"github.com/bxcodec/go-clean-arch/domain"
)

func (m *mysqlDonorRepository) Reschedule(ctx context.Context, id int64, dataNew *domain.DonorSchedulle) (err error) {
	query := `UPDATE donor_registers set donorSchedulleId = ? WHERE ID = ?`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, dataNew.GetId_DonorSchedule(), id)
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
