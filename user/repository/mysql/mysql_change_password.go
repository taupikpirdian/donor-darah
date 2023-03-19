package mysql

import (
	"context"

	"github.com/bxcodec/go-clean-arch/domain"
)

func (m *mysqlUserRepository) ChangePassword(ctx context.Context, us *domain.UserData) error {

	// Query untuk mengubah kata sandi pengguna
	queryChangePassword := "UPDATE users SET password = ? WHERE id = ?"
	_, err := m.Conn.ExecContext(ctx, queryChangePassword, us.GetPasswordOnUser(), us.GetIdOnUser())
	if err != nil {
		return err
	}
	return nil
}
