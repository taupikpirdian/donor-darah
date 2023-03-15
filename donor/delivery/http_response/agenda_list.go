package http_response

import (
	"github.com/bxcodec/go-clean-arch/domain"
)

type ResponseAgenda struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	DateTime string `json:"dateTime"`
	Address  string `json:"address"`
}

type CustomReponseSingleListAgenda struct {
	Status *Status
	Data   []*ResponseAgenda
}

func MapResponseListAgenda(code int, message string, datas []*domain.DonorRegister) (*CustomReponseSingleListAgenda, error) {
	res := []*ResponseAgenda{}
	httpResponse := &CustomReponseSingleListAgenda{
		Status: &Status{
			Code:    code,
			Message: message,
		},
		Data: res,
	}

	return httpResponse, nil
}
