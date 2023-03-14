package mysql

import (
	"context"

	"github.com/bxcodec/go-clean-arch/domain"
)

func (m *mysqlDonorRepository) DonorRegister(ctx context.Context, u *domain.DonorRegister) (int64, error) {
	query := `INSERT  donor_registers SET code=? , userId=? , donorSchedulleId=? , isApprove=? , donorProof=? , updatedAt=? , createdAt=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return 0, err
	}

	res, err := stmt.ExecContext(ctx, u.GetCode_DonorRegister(), u.GetUserId_DonorRegister(), u.GetDonorSchedulleId_DonorRegister(), u.GetIsApprove_DonorRegister(), u.GetDonorProof_DonorRegister(), u.GetUpdateAt_DonorRegister(), u.GetCreatedAt_DonorRegister())
	if err != nil {
		return 0, err
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastID, nil
}
