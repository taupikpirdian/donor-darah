package mysql

import (
	"context"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/sirupsen/logrus"
)

func (m *mysqlUserRepository) fetchUnit(ctx context.Context, query string, args ...interface{}) (result []*domain.UnitDTO, err error) {
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

	result = make([]*domain.UnitDTO, 0)
	for rows.Next() {
		t := &domain.UnitDTO{}
		err = rows.Scan(
			&t.Id,
			&t.Name,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}

func (m *mysqlUserRepository) GetUnit(ctx context.Context) ([]*domain.UnitDTO, error) {
	query := `SELECT id,name FROM units ORDER BY name `

	res, err := m.fetchUnit(ctx, query)
	if err != nil {
		return nil, err
	}

	return res, nil
}
