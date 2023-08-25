package http_response

import (
	"time"

	"github.com/bxcodec/go-clean-arch/domain"
)

type CustomReponseListUser struct {
	Status *Status
	Data   []*userResponse
}

func MapperUser(data *domain.User) (*domain.UserData, error) {
	dataUser, errEntity := domain.NewUser3(data)
	if errEntity != nil {
		return nil, errEntity
	}

	return dataUser, nil
}

type userResponse struct {
	Id             string               `json:"id"`
	MemberCode     string               `json:"memberCode"`
	Name           string               `json:"name" validate:"required"`
	Email          string               `json:"email" validate:"required"`
	Phone          string               `json:"phone" validate:"required"`
	PlaceOfBirth   string               `json:"placeOfBirth"`
	DateOfBirth    string               `json:"dateOfBirth"`
	Gender         string               `json:"gender"`
	Address        string               `json:"address"`
	PostalCode     string               `json:"postalCode"`
	Role           string               `json:"role"`
	Job            jobResponse          `json:"job"`
	Unit           unitResponse         `json:"unit"`
	SubDistrict    districtDataResponse `json:"subDistrict"`
	Village        villageDataResponse  `json:"village"`
	AgendaResponse agendaResponse       `json:"agenda"`
}

type jobResponse struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type unitResponse struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type districtDataResponse struct {
	Id   int64  `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type villageDataResponse struct {
	Id   int64  `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type agendaResponse struct {
	TotalDonor           int64                          `json:"totalDonor"`
	LastDonor            string                         `json:"lastDonor"`
	NextDonor            string                         `json:"nextDonor"`
	HistoryDonorRegister []historyDonorRegisterResponse `json:"historyDonorRegister"`
}

type historyDonorRegisterResponse struct {
	Id         int64     `json:"id"`
	Code       string    `json:"code"`
	PlaceName  string    `json:"placeName"`
	Address    string    `json:"address"`
	Date       string    `json:"schedule_date"`
	TimeStart  string    `json:"timeStart"`
	TimeEnd    string    `json:"timeEnd"`
	Type       string    `json:"type"`
	DonorProof string    `json:"donorProof"`
	CreatedAt  time.Time `json:"createdAt"`
}

func MapResponseDonorRegister(data domain.HistoryDonorRegister, baseUrl string) historyDonorRegisterResponse {
	url := ""
	if data.DonorProof != "" {
		url = baseUrl + data.DonorProof
	}
	return historyDonorRegisterResponse{
		Id:         data.Id,
		Code:       data.Code,
		PlaceName:  data.PlaceName,
		Address:    data.Address,
		Date:       data.Date.Format("2006-01-02"),
		TimeStart:  data.TimeStart,
		TimeEnd:    data.TimeEnd,
		Type:       data.Type,
		DonorProof: url,
		CreatedAt:  data.CreatedAt,
	}
}

func MapResponseDonorRegisterList(data []domain.HistoryDonorRegister, baseUrl string) []historyDonorRegisterResponse {
	historyDonorRegisters := make([]historyDonorRegisterResponse, 0)
	for _, value := range data {
		dataResponse := MapResponseDonorRegister(value, baseUrl)
		historyDonorRegisters = append(historyDonorRegisters, dataResponse)
	}

	return historyDonorRegisters
}

func MapResponseUser(data *domain.User, baseUrl string) *userResponse {
	var nextDonor string
	var lastDonor string
	if data.Histories.NextDonor != (time.Time{}) {
		nextDonor = data.Histories.NextDonor.Format("2006-01-02")
	} else {
		nextDonor = "" // Return an empty string
	}
	if data.Histories.LastDonor != (time.Time{}) {
		lastDonor = data.Histories.LastDonor.Format("2006-01-02")
	} else {
		lastDonor = "" // Return an empty string
	}
	historyDonorRegister := MapResponseDonorRegisterList(data.Histories.HistoryDonorRegister, baseUrl)

	return &userResponse{
		Id:           data.Id,
		Name:         data.Name,
		Email:        data.Email,
		Phone:        data.Phone,
		PlaceOfBirth: data.PlaceOfBirth,
		DateOfBirth:  data.DateOfBirth,
		Gender:       data.Gender,
		Address:      data.Address,
		PostalCode:   data.PostalCode,
		Role:         data.Role,
		MemberCode:   data.MemberCode,
		Job: jobResponse{
			Id:   data.Job.Id,
			Name: data.Job.Name,
		},
		Unit: unitResponse{
			Id:   data.Unit.Id,
			Name: data.Unit.Name,
		},
		SubDistrict: districtDataResponse{
			Id:   data.SubDistrict.Id,
			Code: data.SubDistrict.Code,
			Name: data.SubDistrict.Name,
		},
		Village: villageDataResponse{
			Id:   data.Village.Id,
			Code: data.Village.Code,
			Name: data.Village.Name,
		},
		AgendaResponse: agendaResponse{
			TotalDonor:           data.Histories.TotalDonor,
			LastDonor:            lastDonor,
			NextDonor:            nextDonor,
			HistoryDonorRegister: historyDonorRegister,
		},
	}
}

func MapResponseUserList(code int, message string, data []*domain.User, baseUrl string) (*CustomReponseListUser, error) {
	datas := make([]*userResponse, 0)
	for _, value := range data {
		dataResponse := MapResponseUser(value, baseUrl)
		datas = append(datas, dataResponse)
	}
	httpResponse := &CustomReponseListUser{
		Status: &Status{
			Code:    code,
			Message: message,
		},
		Data: datas,
	}

	return httpResponse, nil
}
