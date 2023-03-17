package usecase

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/bxcodec/go-clean-arch/domain"
)

func Test_donorUsecase_ListRiwayat(t *testing.T) {
	type fields struct {
		donorRepo      domain.DonorRepository
		contextTimeout time.Duration
	}
	type args struct {
		c      context.Context
		userId int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*domain.DonorRegisterDTO
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
			got, err := dus.ListRiwayat(tt.args.c, tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("donorUsecase.ListRiwayat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("donorUsecase.ListRiwayat() = %v, want %v", got, tt.want)
			}
		})
	}
}
