package mysql

import (
	"context"

	"github.com/bxcodec/go-clean-arch/domain"
)

func (m *mysqlDonorRepository) ListRiwayatByStatus(ctx context.Context, userId int64, status string) ([]*domain.DonorRegisterDTO, error) {
	query := `SELECT donor_registers.id, donor_registers.code, donor_schedulle.placeName, donor_schedulle.address, donor_schedulle.date, donor_schedulle.timeStart, donor_schedulle.timeEnd, donor_schedulle.type, donor_registers.donorProof, donor_registers.createdAt 
	FROM donor_registers
	JOIN donor_schedulle ON donor_schedulle.id = donor_registers.donorSchedulleId 
	WHERE userId = ? 
	AND status = ? 
	ORDER BY donor_registers.createdAt`

	res, err := m.fetchRiwayat(ctx, query, userId, status)
	if err != nil {
		return nil, err
	}

	return res, nil
}
