package mysql

import (
	"context"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/sirupsen/logrus"
)

func (m *mysqlDonorRepository) fetchSchedulle(ctx context.Context, query string, args ...interface{}) (result []*domain.DonorSchedulleDTO, err error) {
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

	result = make([]*domain.DonorSchedulleDTO, 0)
	for rows.Next() {
		t := &domain.DonorSchedulleDTO{}
		err = rows.Scan(
			&t.Id,
			&t.PlaceName,
			&t.Address,
			&t.Date,
			&t.TimeStart,
			&t.TimeEnd,
			&t.TypeSchedulle,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}

func (m *mysqlDonorRepository) ListSchedulle(ctx context.Context, unitId int64) ([]*domain.DonorSchedulleDTO, error) {
	query := `SELECT donor_schedulle.id, donor_schedulle.placeName, donor_schedulle.address, donor_schedulle.date, donor_schedulle.timeStart, donor_schedulle.timeEnd, donor_schedulle.type 
	FROM donor_schedulle
	WHERE unitId = ? 
	ORDER BY donor_schedulle.date DESC, donor_schedulle.timeStart ASC`

	res, err := m.fetchSchedulle(ctx, query, unitId)
	if err != nil {
		return nil, err
	}

	return res, nil
}
