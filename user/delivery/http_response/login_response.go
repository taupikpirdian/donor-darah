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
	Id            string `json:"id"`
	Name          string `json:"name" validate:"required"`
	Email         string `json:"email" validate:"required"`
	Phone         string `json:"phone" validate:"required"`
	JobId         string `json:"jobId"`
	UnitId        string `json:"unitId"`
	PlaceOfBirth  string `json:"placeOfBirth"`
	DateOfBirth   string `json:"dateOfBirth"`
	Gender        string `json:"gender"`
	SubDistrictId string `json:"subDistrictId"`
	VillageId     string `json:"villageId"`
	Address       string `json:"address"`
	PostalCode    string `json:"postalCode"`
	Role          string `json:"role"`
	MemberCode    string `json:"memberCode"`
}

func MapResponseLogin(code int, message string, data *domain.Auth) (*CustomReponseAuth, error) {
	user := &User{
		Id:            data.User.Id,
		Name:          data.User.Name,
		Email:         data.User.Email,
		Phone:         data.User.Phone,
		JobId:         data.User.JobId.String,
		UnitId:        data.User.UnitId.String,
		PlaceOfBirth:  data.User.PlaceOfBirth,
		DateOfBirth:   helper.DateStringFormat(data.User.DateOfBirth),
		Gender:        data.User.Gender,
		SubDistrictId: data.User.SubDistrictId,
		VillageId:     data.User.VillageId,
		Address:       data.User.Address,
		PostalCode:    data.User.PostalCode,
		Role:          data.User.Role,
		MemberCode:    data.User.MemberCode.String,
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
