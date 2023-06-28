package mysql

import (
	"context"
	"time"

	"github.com/bxcodec/go-clean-arch/donor/delivery/http_request"
)

func (m *mysqlDonorRepository) SchedulleStore(ctx context.Context, req *http_request.SchedulleStore) error {

	query := `INSERT donor_schedulle SET unitId=?, placeName=?, address=?, date=?, timeStart=?, timeEnd=?, type=?, updatedAt=?, createdAt=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	_, errExec := stmt.ExecContext(ctx, req.UnitId, req.PlaceName, req.Address, req.Date, req.TimeStart, req.TimeEnd, req.Type, time.Now(), time.Now())
	if errExec != nil {
		return errExec
	}

	return nil
}
