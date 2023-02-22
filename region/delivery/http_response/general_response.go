package http_response

import (
	"encoding/json"
	"time"

	"github.com/bxcodec/go-clean-arch/domain"
)

type Status struct {
	Code    int
	Message string
}

type CustomReponseSingle struct {
	Status *Status
	Data   []*ResponseItemDistrictJson
}

type ResponseItemDistrictJson struct {
	Id        int64     `json:"id"`
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	UpdatedAt time.Time `json:"updatedAt"`
	CreatedAt time.Time `json:"createdAt"`
}

func MapResponseDistrict(code int, message string, dataDistrict []*domain.DistrictData) ([]byte, error) {
	httpResponse := &CustomReponseSingle{
		Status: &Status{
			Code:    code,
			Message: message,
		},
		Data: nil,
	}

	respJson, err := json.Marshal(httpResponse)
	if err != nil {
		return nil, err
	}

	return respJson, nil
}
