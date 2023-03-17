package mysql

import (
	"context"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/sirupsen/logrus"
)

func (m *mysqlDonorRepository) fetchRiwayat(ctx context.Context, query string, args ...interface{}) (result []*domain.DonorRegisterDTO, err error) {
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
			&t.DonorSchedulle.TypeSchedulle,
			&t.DonorProof,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}

func (m *mysqlDonorRepository) ListRiwayat(ctx context.Context, userId int64) ([]*domain.DonorRegisterDTO, error) {
	query := `SELECT donor_registers.id, donor_registers.code, donor_schedulle.placeName, donor_schedulle.address, donor_schedulle.date, donor_schedulle.timeStart, donor_schedulle.timeEnd, donor_schedulle.type, donor_registers.donorProof 
	FROM donor_registers
	JOIN donor_schedulle ON donor_schedulle.id = donor_registers.donorSchedulleId 
	WHERE userId = ? 
	ORDER BY donor_registers.createdAt`

	res, err := m.fetchRiwayat(ctx, query, userId)
	if err != nil {
		return nil, err
	}

	return res, nil
}
