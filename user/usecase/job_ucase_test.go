package usecase_test

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/bxcodec/go-clean-arch/cfg"
	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/domain/mocks"
	testdata "github.com/bxcodec/go-clean-arch/user/test_data"
	"github.com/bxcodec/go-clean-arch/user/usecase"
	"github.com/stretchr/testify/mock"
)

func Test_userUsecase_GetJob(t *testing.T) {
	type fields struct {
		userRepo       domain.UserRepository
		serviceMail    domain.MailService
		contextTimeout time.Duration
		donorRepo      domain.DonorRepository
		cfg            cfg.Config
	}
	type args struct {
		c context.Context
	}

	/*
		data general
	*/
	data := testdata.MultipleJob()
	ctx := context.TODO()
	/*
		case negatif: repo error
	*/
	// mock
	repoUJobError := new(mocks.UserRepository)
	repoUJobError.On("GetJob", mock.Anything).
		Times(2).
		Return(nil, errors.New("error"))

	/*
		case positif
	*/
	repoUJob := new(mocks.UserRepository)
	repoUJob.On("GetJob", mock.Anything).
		Times(2).
		Return(data, nil)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*domain.Job
		wantErr bool
	}{
		{
			name: "Error Repo",
			fields: fields{
				userRepo:       repoUJobError,
				serviceMail:    nil,
				contextTimeout: 0,
				donorRepo:      nil,
				cfg:            cfg.Config{},
			},
			args: args{
				c: ctx,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Success",
			fields: fields{
				userRepo:       repoUJob,
				serviceMail:    nil,
				contextTimeout: 0,
				donorRepo:      nil,
				cfg:            cfg.Config{},
			},
			args: args{
				c: ctx,
			},
			want:    data,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := usecase.NewUserUsecase(tt.fields.userRepo, tt.fields.serviceMail, tt.fields.contextTimeout, tt.fields.donorRepo, tt.fields.cfg)
			got, err := us.GetJob(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("userUsecase.GetJob() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userUsecase.GetJob() = %v, want %v", got, tt.want)
			}
		})
	}
}
