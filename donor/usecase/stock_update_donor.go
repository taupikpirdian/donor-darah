package usecase

import (
	"context"

	"github.com/bxcodec/go-clean-arch/donor/delivery/http_request"
)

func (dus *donorUsecase) StockUpdateDonor(c context.Context, req *http_request.BodyBloodStock) error {
	ctx, cancel := context.WithTimeout(c, dus.contextTimeout)
	defer cancel()

	errR := dus.donorRepo.StockUpdateDonor(ctx, req)
	if errR != nil {
		return errR
	}

	for _, data := range req.BodyBloodStockDetail {
		if data.Id != 0 {
			err := dus.donorRepo.UpdateStockDetail(ctx, data.Id, data)
			if err != nil {
				return err
			}
		} else {
			err := dus.donorRepo.StoreStockDetail(ctx, req.Id, data)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
