package usecase

import (
	"context"

	"github.com/bxcodec/go-clean-arch/donor/delivery/http_request"
)

func (dus *donorUsecase) StoreStock(c context.Context, unitId int64, req *http_request.BodyBloodStock) error {
	ctx, cancel := context.WithTimeout(c, dus.contextTimeout)
	defer cancel()

	/*
		store to blood_stock
		store to blood_stock_detail
	*/
	id, errR := dus.donorRepo.StoreStock(ctx, unitId, req.Title)
	if errR != nil {
		return errR
	}

	for _, data := range req.BodyBloodStockDetail {
		err := dus.donorRepo.StoreStockDetail(ctx, id, data)
		if err != nil {
			return err
		}
	}
	return nil
}
