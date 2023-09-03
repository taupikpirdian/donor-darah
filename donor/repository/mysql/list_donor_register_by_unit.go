package mysql

import (
	"context"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/sirupsen/logrus"
)

func (m *mysqlDonorRepository) fetchListDonorRegisterByUnit(ctx context.Context, query string, args ...interface{}) (result []*domain.DonorRegisterDTO, err error) {
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
			&t.Status,
			&t.StatusApprove,
			&t.UserId,
			&t.User.Name,
			&t.User.Email,
			&t.User.Phone,
			&t.User.PlaceOfBirth,
			&t.User.DateOfBirth,
			&t.User.Gender,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}

func (m *mysqlDonorRepository) ListDonorRegisterByUnit(ctx context.Context, unitId int, date string) ([]*domain.DonorRegisterDTO, error) {
	query := `SELECT donor_registers.id, donor_registers.code, donor_registers.status, donor_registers.isApprove, users.id as userId, users.name, users.email, users.phone, profiles.placeOfBirth, profiles.dateOfBirth, profiles.gender 
		FROM donor_registers
		JOIN donor_schedulle ON donor_schedulle.id = donor_registers.donorSchedulleId 
		JOIN users ON users.id = donor_registers.userId 
		JOIN profiles ON profiles.userId = users.id
		WHERE donor_schedulle.unitId = ? AND donor_schedulle.date = ?`

	res, err := m.fetchListDonorRegisterByUnit(ctx, query, unitId, date)
	if err != nil {
		return nil, err
	}

	return res, nil
}
