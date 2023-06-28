package mysql

import (
	"context"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/sirupsen/logrus"
)

func (m *mysqlDonorRepository) fetchListStock(ctx context.Context, query string, args ...interface{}) (result []*domain.DonorStock, err error) {
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

	result = make([]*domain.DonorStock, 0)
	for rows.Next() {
		t := &domain.DonorStock{}
		err = rows.Scan(
			&t.Id,
			&t.Title,
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

func (m *mysqlDonorRepository) ListStock(ctx context.Context, unitId int64) ([]*domain.DonorStock, error) {
	query := `SELECT id, title, createdAt, updatedAt FROM blood_stock where unitId = ? ORDER BY title`

	res, err := m.fetchListStock(ctx, query, unitId)
	if err != nil {
		return nil, err
	}

	return res, nil
}
