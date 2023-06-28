package mysql

import (
	"context"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/sirupsen/logrus"
)

func (m *mysqlDonorRepository) fetchListAnswer(ctx context.Context, query string, args ...interface{}) (result []*domain.DonorRegisterQuestionerDTO, err error) {
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

	result = make([]*domain.DonorRegisterQuestionerDTO, 0)
	for rows.Next() {
		t := &domain.DonorRegisterQuestionerDTO{}
		err = rows.Scan(
			&t.Id,
			&t.CodeQuestion,
			&t.Answer,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}

func (m *mysqlDonorRepository) ListAnswer(ctx context.Context, registerId int64) ([]*domain.DonorRegisterQuestionerDTO, error) {
	query := `SELECT id, codeQuestion, answer FROM donor_register_questioners where donorRegisterId = ? ORDER BY codeQuestion`

	res, err := m.fetchListAnswer(ctx, query, registerId)
	if err != nil {
		return nil, err
	}

	return res, nil
}
