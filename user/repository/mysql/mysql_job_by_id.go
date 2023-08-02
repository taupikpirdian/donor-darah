package mysql

import (
	"context"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/user/repository/model"
	"github.com/sirupsen/logrus"
)

func (m *mysqlUserRepository) fetchJobById(ctx context.Context, query string, args ...interface{}) (result []*model.JobModel, err error) {
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

	result = make([]*model.JobModel, 0)
	for rows.Next() {
		t := &model.JobModel{}
		err = rows.Scan(
			&t.Id,
			&t.Name,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}

func (m *mysqlUserRepository) GetJobById(ctx context.Context, id string) (res *model.JobModel, err error) {
	query := `SELECT id,name FROM jobs WHERE id = ? `

	list, err := m.fetchJobById(ctx, query, id)
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
