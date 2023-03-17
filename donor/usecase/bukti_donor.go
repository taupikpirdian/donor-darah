package usecase

import (
	"context"
)

func (dus *donorUsecase) UploadBukti(c context.Context, id int64) error {
	ctx, cancel := context.WithTimeout(c, dus.contextTimeout)
	defer cancel()

	errR := dus.donorRepo.UploadBukti(ctx, id)
	if errR != nil {
		return errR
	}

	return nil
}
