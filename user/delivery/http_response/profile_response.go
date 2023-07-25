package http_response

import (
	"github.com/bxcodec/go-clean-arch/domain"
)

type ResponseProfileJson struct {
	Id         int64  `json:"id"`
	MemberCode string `json:"memberCode"`
	Name       string `json:"name"`
	UrlImage   string `json:"urlImage"`
	TotalDonor int    `json:"totalDonor"`
	LastDonor  string `json:"lastDonor"`
	NextDonor  string `json:"nextDonor"`
	User       User   `json:"user"`
}

type CustomReponseProfile struct {
	Status *Status
	Data   *ResponseProfileJson
}

func MapResponseProfile(code int, message string, data *domain.Profile) (*CustomReponseProfile, error) {
	var lastDonor string
	var nextDonor string
	if !data.LastDonor.IsZero() {
		lastDonor = data.LastDonor.Format("2006-01-02")
	}
	if !data.NextDonor.IsZero() {
		nextDonor = data.NextDonor.Format("2006-01-02")
	}

	user := &User{
		Id:            data.User.Id,
		Name:          data.User.Name,
		Email:         data.User.Email,
		Phone:         data.User.Phone,
		JobId:         data.User.JobId.String,
		UnitId:        data.User.UnitId.String,
		PlaceOfBirth:  data.User.PlaceOfBirth,
		DateOfBirth:   data.User.DateOfBirth,
		Gender:        data.User.Gender,
		SubDistrictId: data.User.SubDistrictId,
		VillageId:     data.User.VillageId,
		Address:       data.User.Address,
		PostalCode:    data.User.PostalCode,
		Role:          data.User.Role,
		MemberCode:    data.User.MemberCode.String,
	}

	httpResponse := &CustomReponseProfile{
		Status: &Status{
			Code:    code,
			Message: message,
		},
		Data: &ResponseProfileJson{
			Id:         data.Id,
			MemberCode: data.MemberCode.String,
			Name:       data.Name,
			UrlImage:   data.UrlImage,
			TotalDonor: int(data.TotalDonor),
			LastDonor:  lastDonor,
			NextDonor:  nextDonor,
			User:       *user,
		},
	}

	return httpResponse, nil
}
