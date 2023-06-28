package mysql

import (
	"context"
	"time"

	"github.com/bxcodec/go-clean-arch/donor/delivery/http_request"
)

func (m *mysqlDonorRepository) StoreStockDetail(ctx context.Context, stockId int64, req *http_request.BodyBloodStockDetail) error {

	query := `INSERT blood_stock_detail SET blood_stock_id=?, title=?, stock=?, updatedAt=?, createdAt=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	_, errExec := stmt.ExecContext(ctx, stockId, req.Title, req.Stock, time.Now(), time.Now())
	if errExec != nil {
		return errExec
	}

	return nil
}
