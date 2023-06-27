package http_response

type CustomReponseGeneral struct {
	Status *Status
	Data   interface{}
}

func MapResponseGeneral(code int, message string, datas interface{}) (*CustomReponseGeneral, error) {
	httpResponse := &CustomReponseGeneral{
		Status: &Status{
			Code:    code,
			Message: message,
		},
		Data: datas,
	}

	return httpResponse, nil
}
