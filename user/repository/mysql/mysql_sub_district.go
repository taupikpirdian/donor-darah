package mysql

import (
	"context"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/user/repository/model"
	"github.com/sirupsen/logrus"
)

func (m *mysqlUserRepository) fetchSubDistrictById(ctx context.Context, query string, args ...interface{}) (result []*model.SubDistrictModel, err error) {
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

	result = make([]*model.SubDistrictModel, 0)
	for rows.Next() {
		t := &model.SubDistrictModel{}
		err = rows.Scan(
			&t.Id,
			&t.Code,
			&t.RegencyId,
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

func (m *mysqlUserRepository) GetSubDistrictById(ctx context.Context, code string) (res *model.SubDistrictModel, err error) {
	query := `SELECT id,code,regency_id,name FROM sub_districts WHERE code = ? `

	list, err := m.fetchSubDistrictById(ctx, query, code)
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
