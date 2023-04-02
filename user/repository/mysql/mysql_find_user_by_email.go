package mysql

import (
	"context"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/user/delivery/http_response"
)

func (m *mysqlUserRepository) FindUserByEmail(ctx context.Context, email string) (res *domain.UserData, err error) {
	query := `SELECT id,name,email,phone,password FROM users WHERE email = ? `
	list, err := m.fetchUser(ctx, query, email)
	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		mapper, errMapper := http_response.MapperUser(list[0])
		if errMapper != nil {
			return res, errMapper
		}
		res = mapper
	} else {
		return res, domain.ErrNotFound
	}

	return
}
