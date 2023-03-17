package usecase

import (
	"context"
	"testing"
	"time"

	"github.com/bxcodec/go-clean-arch/domain"
)

func Test_donorUsecase_DonorRegister(t *testing.T) {
	type fields struct {
		donorRepo      domain.DonorRepository
		contextTimeout time.Duration
	}
	type args struct {
		c      context.Context
		userId int64
		req    *domain.RequestRegisterDonor
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dus := &donorUsecase{
				donorRepo:      tt.fields.donorRepo,
				contextTimeout: tt.fields.contextTimeout,
			}
			if err := dus.DonorRegister(tt.args.c, tt.args.userId, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("donorUsecase.DonorRegister() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
