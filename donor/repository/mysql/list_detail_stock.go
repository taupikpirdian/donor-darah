package mysql

import (
	"context"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/sirupsen/logrus"
)

func (m *mysqlDonorRepository) fetchListDetailStock(ctx context.Context, query string, args ...interface{}) (result []*domain.DonorDetailStock, err error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			logrus.Error(errRow)
		}
	}()

	result = make([]*domain.DonorDetailStock, 0)
	for rows.Next() {
		t := &domain.DonorDetailStock{}
		err = rows.Scan(
			&t.Id,
			&t.Title,
			&t.Stock,
			&t.CreatedAt,
			&t.UpdatedAt,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}

func (m *mysqlDonorRepository) ListDetailStock(ctx context.Context, stockId int64) ([]*domain.DonorDetailStock, error) {
	query := `SELECT id, title, stock, createdAt, updatedAt FROM blood_stock_detail where blood_stock_id = ? ORDER BY stock DESC`

	res, err := m.fetchListDetailStock(ctx, query, stockId)
	if err != nil {
		return nil, err
	}

	return res, nil
}
