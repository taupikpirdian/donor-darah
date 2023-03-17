package mysql

import (
	"context"

	"github.com/bxcodec/go-clean-arch/domain"
)

func (m *mysqlUserRepository) ChangePassword(ctx context.Context, us *domain.UserData) error {
	// Query untuk mencari pengguna berdasarkan ID
	// queryFindUserByID := "SELECT * FROM users WHERE id = ?"
	// list, err := m.fetchUser(ctx, queryFindUserByID, us.GetIdOnUser())

	// if len(list) > 0 {
	// 	res = list[0]
	// } else {
	// 	return res, domain.ErrNotFound
	// }

	// // Hash kata sandi baru
	// hashedPassword, err := bcrypt.GenerateFromPassword([]byte(us.GetPasswordOnUser()), bcrypt.DefaultCost)
	// if err != nil {
	// 	return nil, err
	// }

	// // Query untuk mengubah kata sandi pengguna
	// queryChangePassword := "UPDATE users SET password = ? WHERE id = ?"
	// _, err = m.ChangePassword(ctx,queryChangePassword,string(hashedPassword),res.Id)
	// if err != nil {
	// 	return nil,err
	// }

	return nil
}

func (m *mysqlUserRepository) UpdatePassword(ctx context.Context, query string, hPassword string, us *domain.UserData) (res *domain.User, err error) {
	// hashedPassword, err := bcrypt.GenerateFromPassword([]byte(us.GetPasswordOnUser()), bcrypt.DefaultCost)
	// if err != nil {
	// 	return nil, err
	// }

	// // Query untuk mengubah kata sandi pengguna
	// queryChangePassword := "UPDATE users SET password = ? WHERE id = ?"
	// _, err = m.UpdatePassword(ctx,query,string(hashedPassword)))
	// if err != nil {
	// 	return nil, err
	// }
	return
}
