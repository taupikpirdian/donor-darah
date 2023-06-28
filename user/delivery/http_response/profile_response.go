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
	httpResponse := &CustomReponseProfile{
		Status: &Status{
			Code:    code,
			Message: message,
		},
		Data: &ResponseProfileJson{
			Id:         data.Id,
			MemberCode: data.MemberCode,
			Name:       data.Name,
			UrlImage:   data.UrlImage,
			TotalDonor: int(data.TotalDonor),
			LastDonor:  lastDonor,
			NextDonor:  nextDonor,
		},
	}

	return httpResponse, nil
}
