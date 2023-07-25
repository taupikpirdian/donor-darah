package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/bxcodec/go-clean-arch/cfg"
	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/domain/mocks"
	testdata "github.com/bxcodec/go-clean-arch/user/test_data"
	"github.com/bxcodec/go-clean-arch/user/usecase"
	"github.com/stretchr/testify/mock"
)

func Test_userUsecase_ForgotPassword(t *testing.T) {
	type fields struct {
		userRepo       domain.UserRepository
		serviceMail    domain.MailService
		contextTimeout time.Duration
		donorRepo      domain.DonorRepository
		cfg            cfg.Config
	}
	type args struct {
		c    context.Context
		user *domain.User
	}

	/*
		data testing
	*/
	ctx := context.TODO()

	userData := testdata.DataUser()
	userDataEntity := &domain.UserData{}
	userDataEntity.SetName(userData.Name)
	userDataEntity.SetEmail(userData.Email)
	userDataEntity.SetPhone(userData.Phone)
	userDataEntity.SetJobId(userData.JobId.String)
	userDataEntity.SetUnitId(userData.UnitId.String)
	userDataEntity.SetPlaceOfBirth(userData.PlaceOfBirth)
	userDataEntity.SetDateOfBirth(userData.DateOfBirth)
	userDataEntity.SetGender(userData.Gender)
	userDataEntity.SetSubDistrictId(userData.SubDistrictId)
	userDataEntity.SetVillageId(userData.VillageId)
	userDataEntity.SetAddress(userData.Address)
	userDataEntity.SetPostalCode(userData.PostalCode)

	/*
		negatif case: error Find User By Email
	*/
	repoUser_ErrorFindUser := new(mocks.UserRepository)
	repoUser_ErrorFindUser.On("FindUserByEmail", mock.Anything, mock.Anything).
		Times(1).
		Return(nil, errors.New("error"))

	/*
		negatif case: error change password
	*/
	repoUser_ErrorChangePassword := new(mocks.UserRepository)
	repoUser_ErrorChangePassword.On("FindUserByEmail", mock.Anything, mock.Anything).
		Times(1).
		Return(userDataEntity, nil)
	repoUser_ErrorChangePassword.On("ChangePassword", mock.Anything, mock.Anything).
		Times(1).
		Return(errors.New("error"))

	/*
		negatif case: error send email
	*/
	repoUser_ErrorSendEmail := new(mocks.UserRepository)
	repoUser_ErrorSendEmail.On("FindUserByEmail", mock.Anything, mock.Anything).
		Times(1).
		Return(userDataEntity, nil)
	repoUser_ErrorSendEmail.On("ChangePassword", mock.Anything, mock.Anything).
		Times(1).
		Return(nil)

	serviceMail_ErrorSendEmail := new(mocks.MailService)
	serviceMail_ErrorSendEmail.On("SendEmail", mock.Anything).
		Times(1).
		Return(errors.New("error"))

	/*
		success
	*/
	repoUser_Success := new(mocks.UserRepository)
	repoUser_Success.On("FindUserByEmail", mock.Anything, mock.Anything).
		Times(1).
		Return(userDataEntity, nil)
	repoUser_Success.On("ChangePassword", mock.Anything, mock.Anything).
		Times(1).
		Return(nil)

	serviceMail := new(mocks.MailService)
	serviceMail.On("SendEmail", mock.Anything).
		Times(1).
		Return(nil)

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Negatif: Error Find User By Email",
			fields: fields{
				userRepo:       repoUser_ErrorFindUser,
				serviceMail:    nil,
				contextTimeout: 0,
				donorRepo:      nil,
				cfg:            cfg.Config{},
			},
			args: args{
				c:    ctx,
				user: userData,
			},
			wantErr: true,
		},
		{
			name: "Negatif: Error Change Password",
			fields: fields{
				userRepo:       repoUser_ErrorChangePassword,
				serviceMail:    nil,
				contextTimeout: 0,
				cfg:            cfg.Config{},
			},
			args: args{
				c:    ctx,
				user: userData,
			},
			wantErr: true,
		},
		{
			name: "Success",
			fields: fields{
				userRepo:       repoUser_Success,
				serviceMail:    serviceMail,
				contextTimeout: 0,
				donorRepo:      nil,
				cfg:            cfg.Config{},
			},
			args: args{
				c:    ctx,
				user: userData,
			},
			wantErr: false,
		},
		{
			name: "Negatif: Error Send Email",
			fields: fields{
				userRepo:       repoUser_ErrorSendEmail,
				serviceMail:    serviceMail_ErrorSendEmail,
				contextTimeout: 0,
				donorRepo:      nil,
				cfg:            cfg.Config{},
			},
			args: args{
				c:    ctx,
				user: userData,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := usecase.NewUserUsecase(tt.fields.userRepo, tt.fields.serviceMail, tt.fields.contextTimeout, tt.fields.donorRepo, tt.fields.cfg)
			if err := us.ForgotPassword(tt.args.c, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("userUsecase.ForgotPassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
