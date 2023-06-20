package http_response

import "github.com/bxcodec/go-clean-arch/domain"

type CustomReponseListUser struct {
	Status *Status
	Data   []*domain.User
}

func MapperUser(data *domain.User) (*domain.UserData, error) {
	dataUser, errEntity := domain.NewUser3(data)
	if errEntity != nil {
		return nil, errEntity
	}

	return dataUser, nil
}

func MapResponseUserList(code int, message string, data []*domain.User) (*CustomReponseListUser, error) {
	httpResponse := &CustomReponseListUser{
		Status: &Status{
			Code:    code,
			Message: message,
		},
		Data: data,
	}

	return httpResponse, nil
}
