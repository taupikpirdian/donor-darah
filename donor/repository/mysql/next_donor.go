package mysql

import (
	"context"

	"github.com/bxcodec/go-clean-arch/domain"
)

func (m *mysqlDonorRepository) NextDonorByStatus(ctx context.Context, userId int64, status string) (res *domain.DonorRegisterDTO, err error) {
	query := `SELECT donor_registers.id, donor_registers.code, donor_schedulle.placeName, donor_schedulle.address, donor_schedulle.date, donor_schedulle.timeStart, donor_schedulle.timeEnd, donor_schedulle.type, donor_registers.donorProof 
	FROM donor_registers
	JOIN donor_schedulle ON donor_schedulle.id = donor_registers.donorSchedulleId 
	WHERE userId = ? 
	AND status = ? 
	ORDER BY donor_schedulle.date`

	list, err := m.fetchRiwayat(ctx, query, userId, status)

	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return nil, nil
	}

	return
}
