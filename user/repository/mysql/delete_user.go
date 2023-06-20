package mysql

import (
	"context"
)

func (m *mysqlUserRepository) DeleteUser(ctx context.Context, id string) error {
	query := `DELETE FROM users WHERE id = ?;`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	_, errExec := stmt.ExecContext(ctx, id)
	if errExec != nil {
		return errExec
	}

	return nil
}

func (m *mysqlUserRepository) DeleteUserProfil(ctx context.Context, id string) error {
	query := `DELETE FROM profiles WHERE userId = ?;`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	_, errExec := stmt.ExecContext(ctx, id)
	if errExec != nil {
		return errExec
	}

	return nil
}
