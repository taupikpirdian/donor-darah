package mysql

import (
	"context"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/sirupsen/logrus"
)

func (m *mysqlUserRepository) fetchUser(ctx context.Context, query string, args ...interface{}) (result []*domain.User, err error) {
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
			&t.Name,
			&t.Email,
			&t.Phone,
			&t.Password,
			&t.Role,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}

func (m *mysqlUserRepository) FindUser(ctx context.Context, us *domain.UserData) (res *domain.User, err error) {

	query := `SELECT id,name,email,phone,password,role FROM users WHERE phone = ? `

	list, err := m.fetchUser(ctx, query, us.GetPhoneOnUser())
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

func (m *mysqlUserRepository) FindUserById(ctx context.Context, us *domain.UserData) (res *domain.User, err error) {

	query := `SELECT id,name,email,phone,password FROM users WHERE id = ? `

	list, err := m.fetchUser(ctx, query, us.GetIdOnUser())
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
