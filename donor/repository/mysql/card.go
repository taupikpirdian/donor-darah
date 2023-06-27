package mysql

import (
	"context"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/sirupsen/logrus"
)

func (m *mysqlDonorRepository) fetchSingleCard(ctx context.Context, query string, args ...interface{}) (result []*domain.Card, err error) {
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

	result = make([]*domain.Card, 0)
	for rows.Next() {
		t := &domain.Card{}
		err = rows.Scan(
			&t.Id,
			&t.Name,
			&t.MemberCode,
			&t.DateOfBirth,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}

func (m *mysqlDonorRepository) GetCard(ctx context.Context, userId int64) (res *domain.Card, err error) {
	query := `SELECT profiles.id, users.name, profiles.code, profiles.dateOfBirth
	FROM profiles
	JOIN users ON users.id = profiles.userId 
	where userId = ?`

	list, err := m.fetchSingleCard(ctx, query, userId)
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
