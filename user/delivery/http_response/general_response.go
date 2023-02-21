package http_response

import "encoding/json"

type Status struct {
	Code    int
	Message string
}

type CustomReponseSingle struct {
	Status *Status
	Data   *ResponseItemJson
}

type ResponseItemJson struct {
	Id    int    `json:"id"`
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required"`
	Phone string `json:"phone" validate:"required"`
}

func MapResponse(code int, message string) ([]byte, error) {
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
