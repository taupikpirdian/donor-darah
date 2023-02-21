package domain

import (
	"context"
	"errors"
	"time"
)

// User is representing the User data struct
type User struct {
	Name                 string `json:"name" validate:"required"`
	Email                string `json:"email" validate:"required"`
	Phone                string `json:"phone" validate:"required"`
	Password             string `json:"password" validate:"required"`
	PasswordConfirmation string `json:"passwordConfirmation" validate:"required"`
	JobId                int64  `json:"jobId"`
	UnitId               int64  `json:"unitId"`
	PlaceOfBirth         string `json:"placeOfBirth"`
	DateOfBirth          string `json:"dateOfBirth"`
	Gender               string `json:"gender"`
	SubDistrictId        int64  `json:"subDistrictId"`
	VillageId            int64  `json:"villageId"`
	Address              string `json:"address"`
	PostalCode           string `json:"postalCode"`
}

type userData struct {
	id          int64
	name        string
	email       string
	phone       string
	password    string
	updatedAt   time.Time
	createdAt   time.Time
	profileData profileData
}

type profileData struct {
	id            int64
	userId        int64
	jobId         int64
	unitId        int64
	placeOfBirth  string
	dateOfBirth   time.Time
	gender        string
	subDistrictId int64
	villageId     int64
	address       string
	postalCode    string
	updatedAt     time.Time
	createdAt     time.Time
}

// UserUsecase represent the user's usecases
type UserUsecase interface {
	Register(ctx context.Context, us *User) (User, error)
}

// UserRepository represent the user's repository contract
type UserRepository interface {
	Register(ctx context.Context, us *User) (User, error)
}

func NewUser(u *User) (*userData, error) {
	if u.Name == "" {
		return nil, errors.New("NAME NOT SET")
	}

	if u.Password != u.PasswordConfirmation {
		return nil, errors.New("PASSWORD CONFIRMATION NOT SAME")
	}

	return &userData{
		name:      u.Name,
		email:     u.Email,
		phone:     u.Phone,
		password:  u.Password,
		updatedAt: time.Now(),
		createdAt: time.Now(),
		profileData: profileData{
			jobId:         u.JobId,
			unitId:        u.UnitId,
			placeOfBirth:  u.PlaceOfBirth,
			dateOfBirth:   time.Now(),
			gender:        u.Gender,
			subDistrictId: u.SubDistrictId,
			villageId:     u.VillageId,
			address:       u.Address,
			postalCode:    u.PostalCode,
		},
	}, nil
}

func (cu *userData) GetIdOnUser() int {
	return int(cu.id)
}

func (cu *userData) GetPlaceOfBirthOnProfile() string {
	return cu.profileData.placeOfBirth
}

func (cu *userData) GetDateOfBirthOnProfile() time.Time {
	return cu.profileData.dateOfBirth
}

func (cu *userData) GetGenderOnProfile() string {
	return cu.profileData.gender
}

func (cu *userData) GetJobIdOnProfile() int64 {
	return cu.profileData.jobId
}

func (cu *userData) GetPmiIdOnProfile() int64 {
	return cu.profileData.unitId
}

func (cu *userData) GetPhoneOnUser() string {
	return cu.phone
}

func (cu *userData) GetEmailOnUser() string {
	return cu.email
}

func (cu *userData) GetAddressOnProfile() string {
	return cu.profileData.address
}

func (cu *userData) GetSubDistrictIdOnProfile() int64 {
	return cu.profileData.subDistrictId
}

func (cu *userData) GetVillageIdOnProfile() int64 {
	return cu.profileData.subDistrictId
}

func (cu *userData) GetPostalCodeOnProfile() string {
	return cu.profileData.postalCode
}
