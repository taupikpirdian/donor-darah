package http_response

import (
	"time"

	"github.com/bxcodec/go-clean-arch/domain"
)

type ResponseAgendaSingle struct {
	Id        int64        `json:"id"`
	Code      string       `json:"code"`
	Name      string       `json:"name"`
	Date      time.Time    `json:"date"`
	TimeStart string       `json:"timeStart"`
	TimeEnd   string       `json:"timeEnd"`
	Unit      ResponseUnit `json:"unit"`
	User      ResponseUser `json:"user"`
}

type ResponseUnit struct {
	Name string `json:"name"`
}

type ResponseUser struct {
	Name string `json:"name"`
}

type CustomReponseSingleAgenda struct {
	Status *Status
	Data   *ResponseAgendaSingle
}

func MapResponseSingleAgenda(code int, message string, data *domain.DonorRegisterDTO) (*CustomReponseSingleAgenda, error) {
	res := &ResponseAgendaSingle{
		Id:        data.Id,
		Code:      data.Code,
		Name:      data.DonorSchedulle.PlaceName,
		Date:      data.DonorSchedulle.Date,
		TimeStart: data.DonorSchedulle.TimeStart,
		TimeEnd:   data.DonorSchedulle.TimeEnd,
		User: ResponseUser{
			Name: data.User.Name,
		},
		Unit: ResponseUnit{
			Name: data.Unit.Name,
		},
	}

	httpResponse := &CustomReponseSingleAgenda{
		Status: &Status{
			Code:    code,
			Message: message,
		},
		Data: res,
	}

	return httpResponse, nil
}
