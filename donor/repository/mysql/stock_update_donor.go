package mysql

import (
	"context"
	"time"

	"github.com/bxcodec/go-clean-arch/donor/delivery/http_request"
)

func (m *mysqlDonorRepository) StockUpdateDonor(ctx context.Context, req *http_request.BodyBloodStock) error {
	query := `UPDATE blood_stock SET title=?, updatedAt=? WHERE id=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	_, errExec := stmt.ExecContext(ctx, req.Title, time.Now(), req.Id)
	if errExec != nil {
		return errExec
	}

	return nil
}
