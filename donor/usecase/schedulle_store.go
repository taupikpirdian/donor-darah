package usecase

import (
	"context"

	"github.com/bxcodec/go-clean-arch/donor/delivery/http_request"
)

func (dus *donorUsecase) SchedulleStore(c context.Context, req *http_request.SchedulleStore) error {
	ctx, cancel := context.WithTimeout(c, dus.contextTimeout)
	defer cancel()

	errR := dus.donorRepo.SchedulleStore(ctx, req)
	if errR != nil {
		return errR
	}
	return nil
}
