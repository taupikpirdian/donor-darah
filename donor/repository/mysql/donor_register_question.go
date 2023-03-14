package mysql

import (
	"context"

	"github.com/bxcodec/go-clean-arch/domain"
)

func (m *mysqlDonorRepository) DonorRegisterQuestioner(ctx context.Context, u *domain.DonorRegisterQuestioner, donorRegisterId int64) error {
	query := `INSERT  donor_register_questioners SET donorRegisterId=? , codeQuestion=? , title=? , answer=? , updatedAt=? , createdAt=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	_, errQ := stmt.ExecContext(ctx, donorRegisterId, u.GetCodeQuestion_DonorRegister(), u.GetTitle_DonorRegister(), u.GetAnswer_DonorRegister(), u.GetUpdateAt_DonorRegisterQ(), u.GetCreatedAt_DonorRegisterQ())
	if errQ != nil {
		return errQ
	}

	return nil
}
