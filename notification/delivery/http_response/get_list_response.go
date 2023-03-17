package http_response

import (
	"github.com/bxcodec/go-clean-arch/domain"
)

type CustomReponseSingle struct {
	Status *Status
	Data   []*domain.NotificationData
}

func MapResponseNotificationList(code int, message string, data []*domain.NotificationData) (*CustomReponseSingle, error) {
	httpResponse := &CustomReponseSingle{
		Status: &Status{
			Code:    code,
			Message: message,
		},
		Data: data,
	}

	return httpResponse, nil
}
