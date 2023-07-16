package http_response

import (
	"time"

	"github.com/bxcodec/go-clean-arch/domain"
)

type ResponseRiwayat struct {
	Id            int64     `json:"id"`
	Name          string    `json:"name"`
	Date          time.Time `json:"date"`
	TypeSchedulle string    `json:"type"`
	DonorProof    string    `json:"donorProof"`
}

type CustomReponseListRiwayat struct {
	Status *Status
	Data   []*ResponseRiwayat
}

func MapResponseListRiwayat(code int, message string, datas []*domain.DonorRegisterDTO, domain string) (*CustomReponseListRiwayat, error) {
	res := []*ResponseRiwayat{}
	for _, data := range datas {
		if data.DonorProof != "" {
			data.DonorProof = domain + data.DonorProof
		}
		add := &ResponseRiwayat{
			Id:            data.Id,
			Name:          data.DonorSchedulle.PlaceName,
			Date:          data.DonorSchedulle.Date,
			TypeSchedulle: data.DonorSchedulle.TypeSchedulle,
			DonorProof:    data.DonorProof,
		}
		res = append(res, add)
	}

	httpResponse := &CustomReponseListRiwayat{
		Status: &Status{
			Code:    code,
			Message: message,
		},
		Data: res,
	}

	return httpResponse, nil
}
