package usecase

import (
	"context"
	"errors"
)

func (dus *donorUsecase) SchedulleDelete(c context.Context, id int64) error {
	ctx, cancel := context.WithTimeout(c, dus.contextTimeout)
	defer cancel()

	// jika sudah ada yang mendaftar di donor_registers, tidak bisa d hapus
	datas, err := dus.donorRepo.FindDonorRegister(ctx, id)
	if err != nil {
		return err
	}

	if len(datas) > 0 {
		return errors.New("Data tidak bisa dihapus, jadwal sudah memiliki pendaftar donor")
	}

	errR := dus.donorRepo.SchedulleDelete(ctx, id)
	if errR != nil {
		return errR
	}
	return nil
}
