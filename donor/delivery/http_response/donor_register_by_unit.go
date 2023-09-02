package http_response

import (
	"github.com/bxcodec/go-clean-arch/domain"
)

type RegisterUser struct {
	Id           int    `json:"id"`
	Code         string `json:"code"`
	Status       string `json:"status"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	PlaceOfBirth string `json:"pob"`
	DateOfBirth  string `json:"dob"`
	Gender       string `json:"gender"`
}

type CustomReponseDonorRegisterByUnit struct {
	Status *Status
	Data   []RegisterUser
}

func MapResponseDonorRegisterByUnit(code int, message string, datas []*domain.DonorRegisterDTO) (*CustomReponseDonorRegisterByUnit, error) {
	registers := make([]RegisterUser, 0)
	for _, value := range datas {
		register := RegisterUser{
			Id:           int(value.Id),
			Code:         value.Code,
			Status:       value.Status,
			Name:         value.User.Name,
			Email:        value.User.Email,
			Phone:        value.User.Phone,
			PlaceOfBirth: value.User.PlaceOfBirth,
			DateOfBirth:  value.User.DateOfBirth,
			Gender:       value.User.Gender,
		}
		registers = append(registers, register)
	}
	httpResponse := &CustomReponseDonorRegisterByUnit{
		Status: &Status{
			Code:    code,
			Message: message,
		},
		Data: registers,
	}

	return httpResponse, nil
}
