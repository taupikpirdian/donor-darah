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

func MapResponseProfileError(code int, message string) (*CustomReponseProfile, error) {
	httpResponse := &CustomReponseProfile{
		Status: &Status{
			Code:    code,
			Message: message,
		},
		Data: nil,
	}

	return httpResponse, nil
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

	job := Job{
		Id:   data.User.Job.Id,
		Name: data.User.Job.Name,
	}
	unit := Unit{
		Id:   data.User.Unit.Id,
		Name: data.User.Unit.Name,
	}
	sub_district := SubDistrict{
		Id:   data.User.SubDistrict.Id,
		Code: data.User.SubDistrict.Code,
		Name: data.User.SubDistrict.Name,
	}
	village := Village{
		Id:            data.User.Village.Id,
		Code:          data.User.Village.Code,
		Name:          data.User.Village.Name,
		SubDistrictId: data.User.Village.SubDistrictId,
	}

	user := &User{
		Id:            data.User.Id,
		Name:          data.User.Name,
		Email:         data.User.Email,
		Phone:         data.User.Phone,
		JobId:         data.User.JobId,
		UnitId:        data.User.UnitId,
		PlaceOfBirth:  data.User.PlaceOfBirth,
		DateOfBirth:   data.User.DateOfBirth,
		Gender:        data.User.Gender,
		SubDistrictId: data.User.SubDistrictId,
		VillageId:     data.User.VillageId,
		Address:       data.User.Address,
		PostalCode:    data.User.PostalCode,
		Role:          data.User.Role,
		MemberCode:    data.User.MemberCode,
		Job:           job,
		Unit:          unit,
		SubDistrict:   sub_district,
		Village:       village,
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
