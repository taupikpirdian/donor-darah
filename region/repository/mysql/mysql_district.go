package mysql

import (
	"context"
	"database/sql"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/sirupsen/logrus"
)

type mysqlRegionRepository struct {
	Conn *sql.DB
}

// NewMysqlArticleRepository will create an object that represent the article.Repository interface
func NewMysqlRegionRepository(conn *sql.DB) domain.RegionRepository {
	return &mysqlRegionRepository{conn}
}

func (m *mysqlRegionRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []*domain.DistrictData, err error) {
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

	result = make([]*domain.DistrictData, 0)
	for rows.Next() {
		t := &domain.DistrictData{}
		err = rows.Scan(
			&t.Id,
			&t.Code,
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

func (m *mysqlRegionRepository) GetDistrict(ctx context.Context) ([]*domain.DistrictData, error) {
	query := `SELECT id,code,name FROM sub_districts ORDER BY name `

	res, err := m.fetch(ctx, query)
	if err != nil {
		return nil, err
	}

	return res, nil
}
