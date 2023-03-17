package usecase

import (
	"context"
	"mime/multipart"
	"testing"
	"time"

	"github.com/bxcodec/go-clean-arch/domain"
)

func Test_donorUsecase_UploadBukti(t *testing.T) {
	type fields struct {
		donorRepo      domain.DonorRepository
		contextTimeout time.Duration
	}
	type args struct {
		c    context.Context
		id   int64
		file *multipart.FileHeader
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
			if err := dus.UploadBukti(tt.args.c, tt.args.id, tt.args.file); (err != nil) != tt.wantErr {
				t.Errorf("donorUsecase.UploadBukti() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
