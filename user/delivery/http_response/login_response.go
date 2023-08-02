package http_response

import (
	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/helper"
)

type CustomReponseAuth struct {
	Status *Status
	Data   *Auth
}

type Auth struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

type User struct {
	Id            string      `json:"id"`
	Name          string      `json:"name" validate:"required"`
	Email         string      `json:"email" validate:"required"`
	Phone         string      `json:"phone" validate:"required"`
	JobId         string      `json:"jobId"`
	UnitId        string      `json:"unitId"`
	PlaceOfBirth  string      `json:"placeOfBirth"`
	DateOfBirth   string      `json:"dateOfBirth"`
	Gender        string      `json:"gender"`
	SubDistrictId string      `json:"subDistrictId"`
	VillageId     string      `json:"villageId"`
	Address       string      `json:"address"`
	PostalCode    string      `json:"postalCode"`
	Role          string      `json:"role"`
	MemberCode    string      `json:"memberCode"`
	Job           Job         `json:"job"`
	Unit          Unit        `json:"unit"`
	SubDistrict   SubDistrict `json:"subDistrict"`
	Village       Village     `json:"village"`
}

type Job struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type Unit struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
type SubDistrict struct {
	Id   int64  `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type Village struct {
	Id            int64  `json:"id"`
	SubDistrictId int64  `json:"subDistrictId"`
	Code          string `json:"code"`
	Name          string `json:"name"`
}

func MapResponseLogin(code int, message string, data *domain.Auth) (*CustomReponseAuth, error) {
	user := &User{
		Id:            data.User.Id,
		Name:          data.User.Name,
		Email:         data.User.Email,
		Phone:         data.User.Phone,
		JobId:         data.User.JobId,
		UnitId:        data.User.UnitId,
		PlaceOfBirth:  data.User.PlaceOfBirth,
		DateOfBirth:   helper.DateStringFormat(data.User.DateOfBirth),
		Gender:        data.User.Gender,
		SubDistrictId: data.User.SubDistrictId,
		VillageId:     data.User.VillageId,
		Address:       data.User.Address,
		PostalCode:    data.User.PostalCode,
		Role:          data.User.Role,
		MemberCode:    data.User.MemberCode,
	}
	token := &Auth{
		Token: data.Token,
		User:  *user,
	}

	httpResponse := &CustomReponseAuth{
		Status: &Status{
			Code:    code,
			Message: message,
		},
		Data: token,
	}

	return httpResponse, nil
}
