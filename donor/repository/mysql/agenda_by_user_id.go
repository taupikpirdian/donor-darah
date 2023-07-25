package mysql

import (
	"context"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/sirupsen/logrus"
)

func (m *mysqlDonorRepository) fetchAgendaByUserId(ctx context.Context, query string, args ...interface{}) (result []*domain.DonorRegisterDTO, err error) {
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
			&t.StatusApprove,
			&t.DonorProof,
			&t.Status,
			&t.DonorSchedulle.Id,
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

func (m *mysqlDonorRepository) AgendaByUserId(ctx context.Context, id int64) (res []*domain.DonorRegisterDTO, err error) {
	query := `SELECT donor_registers.id, donor_registers.code, donor_registers.isApprove, donor_registers.donorProof, donor_registers.status, donor_schedulle.id as donor_schedulle_id, donor_schedulle.placeName, donor_schedulle.address, donor_schedulle.date, donor_schedulle.timeStart, donor_schedulle.timeEnd, users.id as idUser, users.name, units.id as idUnit, units.name as unitName
	FROM donor_registers
	JOIN donor_schedulle ON donor_schedulle.id = donor_registers.donorSchedulleId
	JOIN users ON users.id = donor_registers.userId 
	JOIN units ON units.id = donor_schedulle.unitId 
	where donor_registers.userId = ?`

	list, err := m.fetchAgendaByUserId(ctx, query, id)
	if err != nil {
		return nil, err
	}

	return list, err
}