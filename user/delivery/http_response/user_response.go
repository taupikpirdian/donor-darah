package http_response

import "github.com/bxcodec/go-clean-arch/domain"

func MapperUser(data *domain.User) (*domain.UserData, error) {
	dataUser, errEntity := domain.NewUser3(data)
	if errEntity != nil {
		return nil, errEntity
	}

	return dataUser, nil
}
