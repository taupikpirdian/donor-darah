package http_response

import (
	"github.com/bxcodec/go-clean-arch/domain"
)

type CustomReponseSingle struct {
	Status *Status
	Data   []*domain.NotificationData
}

func MapResponseNotificationList(code int, message string, data []*domain.NotificationData) (*CustomReponseSingle, error) {
	for _, v := range data {
		if v.Status == "UNREAD" {
			v.IsRead = false
		} else {
			v.IsRead = true
		}
	}
	httpResponse := &CustomReponseSingle{
		Status: &Status{
			Code:    code,
			Message: message,
		},
		Data: data,
	}

	return httpResponse, nil
}
