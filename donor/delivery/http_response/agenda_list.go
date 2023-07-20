package http_response

import (
	"time"

	"github.com/bxcodec/go-clean-arch/domain"
)

type ResponseAgenda struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	Date      time.Time `json:"date"`
	TimeStart string    `json:"timeStart"`
	TimeEnd   string    `json:"timeEnd"`
	Address   string    `json:"address"`
	Status    string    `json:"status"`
}

type CustomReponseSingleListAgenda struct {
	Status *Status
	Data   []*ResponseAgenda
}

func MapResponseListAgenda(code int, message string, datas []*domain.DonorRegisterDTO) (*CustomReponseSingleListAgenda, error) {
	res := []*ResponseAgenda{}
	for _, data := range datas {
		add := &ResponseAgenda{
			Id:        data.Id,
			Name:      data.DonorSchedulle.PlaceName,
			Date:      data.DonorSchedulle.Date,
			TimeStart: data.DonorSchedulle.TimeStart,
			TimeEnd:   data.DonorSchedulle.TimeEnd,
			Address:   data.DonorSchedulle.Address,
			Status:    data.Status,
		}
		res = append(res, add)
	}

	httpResponse := &CustomReponseSingleListAgenda{
		Status: &Status{
			Code:    code,
			Message: message,
		},
		Data: res,
	}

	return httpResponse, nil
}
