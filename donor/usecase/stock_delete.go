package usecase

import (
	"context"
)

func (dus *donorUsecase) StockDelete(c context.Context, id int64) error {
	ctx, cancel := context.WithTimeout(c, dus.contextTimeout)
	defer cancel()

	err := dus.donorRepo.StockDelete(ctx, id)
	if err != nil {
		return err
	}

	errR := dus.donorRepo.StockDeleteDetail(ctx, id)
	if errR != nil {
		return errR
	}
	return nil
}
