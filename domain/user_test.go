package domain_test

import (
	"errors"
	"testing"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/domain/mocks"
	testdata "github.com/bxcodec/go-clean-arch/user/test_data"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewUser(t *testing.T) {
	type args struct {
		u *domain.User
	}

	userData := testdata.DataUser()

	/*
		negatif case: name null
	*/
	userData_nullName := testdata.DataUser()
	userData_nullName.Name = ""

	/*
		negatif case: password not same
	*/
	userData_passwordNotSame := testdata.DataUser()
	userData_passwordNotSame.Password = "123aaa"

	/*
		negatif case: error hashing
	*/
	mockBcrypt := new(mocks.UserRepository)
	mockBcrypt.On("GenerateFromPassword", mock.Anything, mock.Anything).
		Times(1).
		Return(nil, errors.New("ERROR"))
	// set data
	userDataEntity := &domain.UserData{}
	userDataEntity.SetName(userData.Name)
	userDataEntity.SetEmail(userData.Email)
	userDataEntity.SetPhone(userData.Phone)
	userDataEntity.SetJobId(userData.JobId)
	userDataEntity.SetUnitId(userData.UnitId)
	userDataEntity.SetPlaceOfBirth(userData.PlaceOfBirth)
	userDataEntity.SetDateOfBirth(userData.DateOfBirth)
	userDataEntity.SetGender(userData.Gender)
	userDataEntity.SetSubDistrictId(userData.SubDistrictId)
	userDataEntity.SetVillageId(userData.VillageId)
	userDataEntity.SetAddress(userData.Address)
	userDataEntity.SetPostalCode(userData.PostalCode)

	tests := []struct {
		name    string
		args    args
		want    *domain.UserData
		wantErr bool
	}{
		{
			name: "Error name null",
			args: args{
				u: userData_nullName,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Error password not same",
			args: args{
				u: userData_passwordNotSame,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Success",
			args: args{
				u: userData,
			},
			want:    userDataEntity,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := domain.NewUser(tt.args.u)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				assert.Equalf(t, tt.want.GetNameOnUser(), got.GetNameOnUser(), "NewUser(%v)")
				assert.Equalf(t, tt.want.GetEmailOnUser(), got.GetEmailOnUser(), "NewUser(%v)")
				assert.Equalf(t, tt.want.GetPhoneOnUser(), got.GetPhoneOnUser(), "NewUser(%v)")
				assert.Equalf(t, tt.want.GetJobIdOnProfile(), got.GetJobIdOnProfile(), "NewUser(%v)")
				assert.Equalf(t, tt.want.GetUnitIdOnProfile(), got.GetUnitIdOnProfile(), "NewUser(%v)")
				assert.Equalf(t, tt.want.GetPlaceOfBirthOnProfile(), got.GetPlaceOfBirthOnProfile(), "NewUser(%v)")
				assert.Equalf(t, tt.want.GetDateOfBirthOnProfile(), got.GetDateOfBirthOnProfile(), "NewUser(%v)")
				assert.Equalf(t, tt.want.GetGenderOnProfile(), got.GetGenderOnProfile(), "NewUser(%v)")
				assert.Equalf(t, tt.want.GetSubDistrictIdOnProfile(), got.GetSubDistrictIdOnProfile(), "NewUser(%v)")
				assert.Equalf(t, tt.want.GetVillageIdOnProfile(), got.GetVillageIdOnProfile(), "NewUser(%v)")
				assert.Equalf(t, tt.want.GetAddressOnProfile(), got.GetAddressOnProfile(), "NewUser(%v)")
				assert.Equalf(t, tt.want.GetPostalCodeOnProfile(), got.GetPostalCodeOnProfile(), "NewUser(%v)")
			}
		})
	}
}
