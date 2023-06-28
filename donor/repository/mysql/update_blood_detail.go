package mysql

import (
	"context"
	"time"

	"github.com/bxcodec/go-clean-arch/donor/delivery/http_request"
)

func (m *mysqlDonorRepository) UpdateStockDetail(ctx context.Context, id int64, req *http_request.BodyBloodStockDetail) error {

	query := `UPDATE blood_stock_detail SET title=?, stock=?, updatedAt=? WHERE id=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	_, errExec := stmt.ExecContext(ctx, req.Title, req.Stock, time.Now(), id)
	if errExec != nil {
		return errExec
	}

	return nil
}
