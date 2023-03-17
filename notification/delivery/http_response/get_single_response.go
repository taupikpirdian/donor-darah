package http_response

import "github.com/bxcodec/go-clean-arch/domain"

type CustomReponseSingleData struct {
	Status *Status
	Data   *domain.NotificationData
}

func MapResponseNotificationSingle(code int, message string, data *domain.NotificationData) (*CustomReponseSingleData, error) {
	httpResponse := &CustomReponseSingleData{
		Status: &Status{
			Code:    code,
			Message: message,
		},
		Data: data,
	}

	return httpResponse, nil
}
