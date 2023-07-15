package mysql

import (
	"context"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/sirupsen/logrus"
)

func (m *mysqlDonorRepository) fetchSingle(ctx context.Context, query string, args ...interface{}) (result []*domain.DonorRegisterDTO, err error) {
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

	result = make([]*domain.DonorRegisterDTO, 0)
	for rows.Next() {
		t := &domain.DonorRegisterDTO{}
		err = rows.Scan(
			&t.Id,
			&t.Code,
			&t.DonorSchedulle.PlaceName,
			&t.DonorSchedulle.Address,
			&t.DonorSchedulle.Date,
			&t.DonorSchedulle.TimeStart,
			&t.DonorSchedulle.TimeEnd,
			&t.User.Id,
			&t.User.Name,
			&t.Unit.Id,
			&t.Unit.Name,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}

func (m *mysqlDonorRepository) SingleAgenda(ctx context.Context, id int64) (res *domain.DonorRegisterDTO, err error) {
	query := `SELECT donor_registers.id, donor_registers.code, donor_schedulle.placeName, donor_schedulle.address, donor_schedulle.date, donor_schedulle.timeStart, donor_schedulle.timeEnd, users.id as idUser, users.name, units.id as idUnit, units.name as unitName
	FROM donor_registers
	JOIN donor_schedulle ON donor_schedulle.id = donor_registers.donorSchedulleId
	JOIN users ON users.id = donor_registers.userId 
	JOIN units ON units.id = donor_schedulle.unitId 
	where donor_registers.id = ?`

	list, err := m.fetchSingle(ctx, query, id)
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
