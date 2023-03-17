package mysql

import (
	"context"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/sirupsen/logrus"
)

func (m *mysqlDonorRepository) fetchDonorView(ctx context.Context, query string, args ...interface{}) (result []*domain.DonorRegisterDTO, err error) {
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

func (m *mysqlDonorRepository) UploadBuktiView(ctx context.Context, id int64) (res *domain.DonorRegisterDTO, err error) {
	query := `SELECT donor_registers.id, donor_registers.donorProof
	FROM donor_registers
	where donor_registers.id = ?`

	list, err := m.fetchDonorView(ctx, query, id)
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
