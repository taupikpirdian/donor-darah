package mysql

import (
	"context"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/sirupsen/logrus"
)

func (m *mysqlDonorRepository) fetchSchedule(ctx context.Context, query string, args ...interface{}) (result []*domain.DonorSchedulleDTO, err error) {
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
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}

func (m *mysqlDonorRepository) FindSchedule(ctx context.Context, dto *domain.DonorSchedulleDTO) (res *domain.DonorSchedulleDTO, err error) {
	query := `SELECT id, placeName, address, date, timeStart, timeEnd
	FROM donor_schedulle
	where id = ?`

	list, err := m.fetchSchedule(ctx, query, dto.Id)
	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, domain.ErrNotFound
	}

	return
}
