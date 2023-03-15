package http_response

import (
	"github.com/bxcodec/go-clean-arch/domain"
)

type CustomReponseSingleListAgenda struct {
	Status *Status
	Data   []*domain.DonorRegister
}

func MapResponseListAgenda(code int, message string, datas []*domain.DonorRegister) (*CustomReponseSingleListAgenda, error) {
	httpResponse := &CustomReponseSingleListAgenda{
		Status: &Status{
			Code:    code,
			Message: message,
		},
		Data: datas,
	}

	return httpResponse, nil
}
