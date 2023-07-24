package http_response

import (
	"time"

	"github.com/bxcodec/go-clean-arch/domain"
)

type ResponseSchedulleList struct {
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	Address       string `json:"address"`
	TimeStart     string `json:"timeStart"`
	TimeEnd       string `json:"timeEnd"`
	TypeSchedulle string `json:"type"`
}

type ResponseGroupSchedulleList struct {
	Date      time.Time             `json:"date"`
	Schedulle ResponseSchedulleList `json:"schedule"`
}

type CustomReponseSchedulleList struct {
	Status *Status
	Data   []*ResponseGroupSchedulleMulti
}

type ResponseGroupSchedulleMulti struct {
	Date      time.Time                `json:"date"`
	Schedulle []*ResponseSchedulleList `json:"schedule"`
}

type MyStruct struct {
	GroupByField string
	ValueField   int
}

func MapResponseSchedulleList(code int, message string, datas []*domain.DonorSchedulleDTO) (*CustomReponseSchedulleList, error) {
	groups := make(map[time.Time][]ResponseGroupSchedulleList)
	for _, data := range datas {
		list := ResponseGroupSchedulleList{
			Date: data.Date,
			Schedulle: ResponseSchedulleList{
				Id:            data.Id,
				Name:          data.PlaceName,
				Address:       data.Address,
				TimeStart:     data.TimeStart,
				TimeEnd:       data.TimeEnd,
				TypeSchedulle: data.TypeSchedulle,
			},
		}
		groups[data.Date] = append(groups[data.Date], list)
	}

	res := []*ResponseGroupSchedulleMulti{}
	for key, values := range groups {
		schedulle := []*ResponseSchedulleList{}
		for _, value := range values {
			fAppNSchedulle := ResponseSchedulleList{
				Id:            value.Schedulle.Id,
				Name:          value.Schedulle.Name,
				Address:       value.Schedulle.Address,
				TimeStart:     value.Schedulle.TimeStart,
				TimeEnd:       value.Schedulle.TimeEnd,
				TypeSchedulle: value.Schedulle.TypeSchedulle,
			}
			schedulle = append(schedulle, &fAppNSchedulle)
		}
		f_app := ResponseGroupSchedulleMulti{
			Date:      key,
			Schedulle: schedulle,
		}

		res = append(res, &f_app)
	}

	httpResponse := &CustomReponseSchedulleList{
		Status: &Status{
			Code:    code,
			Message: message,
		},
		Data: res,
	}

	return httpResponse, nil
}
