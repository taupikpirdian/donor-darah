package mysql

import (
	"context"
	"database/sql"

	"github.com/bxcodec/go-clean-arch/domain"
)

func (m *mysqlUserRepository) Register(ctx context.Context, u *domain.UserData) error {
	var role string
	if u.GetRoleOnUser() == "" {
		role = "member"
	} else {
		role = u.GetRoleOnUser()
	}

	query := `INSERT  users SET name=? , email=? , phone=? , password=? , updatedAt=? , createdAt=?, role=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx, u.GetNameOnUser(), u.GetEmailOnUser(), u.GetPhoneOnUser(), u.GetPasswordOnUser(), u.GetUpdateAtOnUser(), u.GetCreatedAtOnUser(), role)
	if err != nil {
		return err
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return err
	}

	u.SetIdNewUser(u, lastID)

	return nil
}

func (m *mysqlUserRepository) StoreProfile(ctx context.Context, u *domain.UserData) error {
	query := `INSERT  profiles SET code=? , userId=? , jobId=? , unitId=? , placeOfBirth=? , dateOfBirth=? , gender=? , subDistrictId=? , villageId=? , address=? , postalCode=? , updatedAt=? , createdAt=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	_, errProses := stmt.ExecContext(ctx, NewNullString(u.GetCodeOnProfile()), u.GetIdOnUser(), NewNullString(u.GetJobIdOnProfile()), NewNullString(u.GetUnitIdOnProfile()), u.GetPlaceOfBirthOnProfile(), u.GetDateOfBirthOnProfile(), u.GetGenderOnProfile(), u.GetSubDistrictIdOnProfile(), u.GetVillageIdOnProfile(), u.GetAddressOnProfile(), u.GetPostalCodeOnProfile(), u.GetUpdateAtOnUser(), u.GetCreatedAtOnUser())
	if errProses != nil {
		return errProses
	}

	return nil
}

func NewNullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}
