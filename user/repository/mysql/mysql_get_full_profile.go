package mysql

import (
	"context"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/sirupsen/logrus"
)

func (m *mysqlUserRepository) fetchProfileFull(ctx context.Context, query string, args ...interface{}) (result []*domain.User, err error) {
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

	result = make([]*domain.User, 0)
	for rows.Next() {
		t := &domain.User{}
		err = rows.Scan(
			&t.Id,
			&t.Role,
			&t.Name,
			&t.Email,
			&t.Phone,
			&t.MemberCode,
			&t.JobId,
			&t.UnitId,
			&t.PlaceOfBirth,
			&t.DateOfBirth,
			&t.Gender,
			&t.SubDistrictId,
			&t.VillageId,
			&t.Address,
			&t.PostalCode,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}

func (m *mysqlUserRepository) GetProfileFull(ctx context.Context, userId int64) (res *domain.User, err error) {
	query := `SELECT u.id, u.role, u.name, u.email, u.phone, p.code, p.jobId, p.unitId, p.placeOfBirth, p.dateOfBirth, p.gender, p.subDistrictId, p.villageId, p.address, p.postalCode
	FROM users u
	LEFT JOIN profiles p ON u.id = p.userId 
	where u.id = ?`

	list, err := m.fetchProfileFull(ctx, query, userId)

	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, nil
	}

	return
}