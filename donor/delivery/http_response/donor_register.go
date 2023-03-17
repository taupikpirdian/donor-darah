package http_response

import (
	"github.com/bxcodec/go-clean-arch/domain"
)

type CustomReponseSingle struct {
	Status *Status
	Data   *domain.DonorRegister
}

func MapResponseDonorRegister(code int, message string) (*CustomReponseSingle, error) {
	httpResponse := &CustomReponseSingle{
		Status: &Status{
			Code:    code,
			Message: message,
		},
		Data: nil,
	}

	return httpResponse, nil
}
