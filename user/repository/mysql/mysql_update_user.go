package mysql

import (
	"context"

	"github.com/bxcodec/go-clean-arch/user/delivery/http_request"
)

func (m *mysqlUserRepository) UpdateUser(ctx context.Context, userId int64, req *http_request.BodyUpdateProfile) error {
	query := "UPDATE users SET name = ?, email = ?, phone = ? WHERE id = ?"
	_, err := m.Conn.ExecContext(ctx, query, req.Name, req.Email, req.Phone, userId)
	if err != nil {
		return err
	}
	return nil
}
