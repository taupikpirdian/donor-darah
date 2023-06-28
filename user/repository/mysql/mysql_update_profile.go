package mysql

import (
	"context"

	"github.com/bxcodec/go-clean-arch/user/delivery/http_request"
)

func (m *mysqlUserRepository) UpdateProfile(ctx context.Context, userId int64, req *http_request.BodyUpdateProfile) error {
	query := "UPDATE profiles SET jobId = ?, unitId = ?, placeOfBirth = ?, placeOfBirth = ?, dateOfBirth = ?, gender = ?, subDistrictId = ?, villageId = ?, postalCode = ?, urlImage = ? WHERE userId = ?"
	_, err := m.Conn.ExecContext(ctx, query, req.JobId, req.UnitId, req.PlaceOfBirth, req.PlaceOfBirth, req.DateOfBirth, req.Gender, req.SubDistrictId, req.VillageId, req.PostalCode, req.Path, userId)
	if err != nil {
		return err
	}
	return nil
}
