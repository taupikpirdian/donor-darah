package http_response

import (
	"github.com/bxcodec/go-clean-arch/domain"
)

type CustomReponseBuktiDonor struct {
	Status *Status
	Data   []*ResponseRiwayat
}

func MapResponseBuktiDonor(code int, message string, datas []*domain.DonorRegisterDTO) (*CustomReponseBuktiDonor, error) {
	httpResponse := &CustomReponseBuktiDonor{
		Status: &Status{
			Code:    code,
			Message: message,
		},
		Data: nil,
	}

	return httpResponse, nil
}
