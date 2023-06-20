package mysql

import (
	"context"

	"github.com/bxcodec/go-clean-arch/domain"
)

func (m *mysqlUserRepository) GetListUser(ctx context.Context) ([]*domain.User, error) {
	query := `SELECT id,name,email,phone,password,role,createdAt FROM users ORDER BY createdAt `

	res, err := m.fetchUser(ctx, query)
	if err != nil {
		return nil, err
	}

	return res, nil
}
