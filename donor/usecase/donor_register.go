package usecase

import (
	"context"

	"github.com/bxcodec/go-clean-arch/domain"
)

func (dus *donorUsecase) DonorRegister(c context.Context, userId int64, req *domain.RequestRegisterDonor) error {
	ctx, cancel := context.WithTimeout(c, dus.contextTimeout)
	defer cancel()

	/*
		build data on entity
	*/
	donorRegister, errEntity := domain.NewDonorRegister(userId, *req)
	if errEntity != nil {
		return errEntity
	}

	/*
		store to table donor_registers
	*/
	lastId, errR := dus.donorRepo.DonorRegister(ctx, donorRegister)
	if errR != nil {
		return errR
	}

	/*
		store to table donor_register_questioners
	*/
	for _, q := range donorRegister.GetQuestion_DonorRegister() {
		errQ := dus.donorRepo.DonorRegisterQuestioner(ctx, q, lastId)
		if errQ != nil {
			return errQ
		}
	}

	return nil
}
