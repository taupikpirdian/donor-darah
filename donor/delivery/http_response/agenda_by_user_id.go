package http_response

import (
	"github.com/bxcodec/go-clean-arch/domain"
)

type RegisterByUserId struct {
	Id                     int64                  `json:"id"`
	Code                   string                 `json:"code"`
	StatusApprove          string                 `json:"is_approve"`
	Status                 string                 `json:"status"`
	DonorProof             string                 `json:"file_donor"`
	DonorSchedulleByUserId DonorSchedulleByUserId `json:"schedule"`
	UserByUserId           UserByUserId           `json:"user"`
}

type DonorSchedulleByUserId struct {
	Id        int64        `json:"id"`
	PlaceName string       `json:"place_name"`
	Address   string       `json:"address"`
	Date      string       `json:"date"`
	TimeStart string       `json:"time_start"`
	TimeEnd   string       `json:"time_end"`
	Unit      UnitByUserId `json:"unit"`
}

type UnitByUserId struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type UserByUserId struct {
	Id         string `json:"id"`
	MemberCode string `json:"member_code"`
	Name       string `json:"name"`
}

type CustomReponseSingleListAgendaByUserId struct {
	Status *Status
	Data   []*RegisterByUserId
}

func MapResponseListAgendaByUserId(code int, message string, datas []*domain.DonorRegisterDTO) (*CustomReponseSingleListAgendaByUserId, error) {
	res := []*RegisterByUserId{}
	for _, data := range datas {
		user := UserByUserId{
			Id:         data.User.Id,
			MemberCode: data.User.MemberCode.String,
			Name:       data.User.Name,
		}
		unit := UnitByUserId{
			Id:   data.Unit.Id,
			Name: data.Unit.Name,
		}
		schedule := DonorSchedulleByUserId{
			Id:        data.DonorSchedulle.Id,
			PlaceName: data.DonorSchedulle.PlaceName,
			Address:   data.DonorSchedulle.Address,
			Date:      data.DonorSchedulle.Date.Format("2006-01-02"),
			TimeStart: data.DonorSchedulle.TimeStart,
			TimeEnd:   data.DonorSchedulle.TimeEnd,
			Unit:      unit,
		}
		add := &RegisterByUserId{
			Id:                     data.Id,
			Code:                   data.Code,
			StatusApprove:          data.StatusApprove,
			Status:                 data.Status,
			DonorProof:             data.DonorProof,
			DonorSchedulleByUserId: schedule,
			UserByUserId:           user,
		}
		res = append(res, add)
	}

	httpResponse := &CustomReponseSingleListAgendaByUserId{
		Status: &Status{
			Code:    code,
			Message: message,
		},
		Data: res,
	}

	return httpResponse, nil
}
