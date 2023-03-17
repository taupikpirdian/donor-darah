package http_response

import "github.com/bxcodec/go-clean-arch/domain"

type CustomReponseAuth struct {
	Status *Status
	Data   *domain.Auth
}

func MapResponseLogin(code int, message string, data *domain.Auth) (*CustomReponseAuth, error) {
	httpResponse := &CustomReponseAuth{
		Status: &Status{
			Code:    code,
			Message: message,
		},
		Data: data,
	}

	return httpResponse, nil
}
