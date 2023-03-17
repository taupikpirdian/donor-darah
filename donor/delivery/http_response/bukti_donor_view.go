package http_response

import (
	"github.com/bxcodec/go-clean-arch/domain"
)

type ResponseRiwayatView struct {
	Id         int64  `json:"id"`
	DonorProof string `json:"url"`
}

type CustomReponseBuktiDonorView struct {
	Status *Status
	Data   *ResponseRiwayatView
}

func MapResponseBuktiDonorView(code int, message string, data *domain.DonorRegisterDTO) (*CustomReponseBuktiDonorView, error) {
	res := ResponseRiwayatView{
		Id:         data.Id,
		DonorProof: data.DonorProof,
	}
	httpResponse := &CustomReponseBuktiDonorView{
		Status: &Status{
			Code:    code,
			Message: message,
		},
		Data: &res,
	}

	return httpResponse, nil
}
