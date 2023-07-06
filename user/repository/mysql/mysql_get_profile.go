package mysql

import (
	"context"
	"fmt"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/sirupsen/logrus"
)

func (m *mysqlUserRepository) fetchProfile(ctx context.Context, query string, args ...interface{}) (result []*domain.Profile, err error) {
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

	result = make([]*domain.Profile, 0)
	for rows.Next() {
		fmt.Println(rows)
		t := &domain.Profile{}
		err = rows.Scan(
			&t.Id,
			&t.Name,
			&t.MemberCode,
			&t.UrlImageFromDB,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}

func (m *mysqlUserRepository) GetProfile(ctx context.Context, userId int64) (res *domain.Profile, err error) {
	query := `SELECT users.id, users.name, profiles.code, profiles.urlImage
	FROM users
	LEFT JOIN profiles ON users.id = profiles.userId 
	where users.id = ?`

	list, err := m.fetchProfile(ctx, query, userId)

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
