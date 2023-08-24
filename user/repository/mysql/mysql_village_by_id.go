package mysql

import (
	"context"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/user/repository/model"
	"github.com/sirupsen/logrus"
)

func (m *mysqlUserRepository) fetchVillageById(ctx context.Context, query string, args ...interface{}) (result []*model.VillageModel, err error) {
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

	result = make([]*model.VillageModel, 0)
	for rows.Next() {
		t := &model.VillageModel{}
		err = rows.Scan(
			&t.Id,
			&t.Code,
			&t.SubDistrictId,
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

func (m *mysqlUserRepository) GetVillageById(ctx context.Context, code string) (res *model.VillageModel, err error) {
	query := `SELECT id,code,sub_district_id,name FROM villages WHERE code = ? `

	list, err := m.fetchVillageById(ctx, query, code)
	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, domain.ErrNotFoundCity
	}

	return
}
