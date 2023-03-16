package http_response

import "github.com/bxcodec/go-clean-arch/domain"

type CustomReponseUnit struct {
	Status *Status
	Data   []*domain.UnitDTO
}

func MapResponseUnit(code int, message string, data []*domain.UnitDTO) (*CustomReponseUnit, error) {
	httpResponse := &CustomReponseUnit{
		Status: &Status{
			Code:    code,
			Message: message,
		},
		Data: data,
	}

	return httpResponse, nil
}
