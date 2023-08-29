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
		if v.Status == "0" {
			v.Status = "UNREAD"
		} else {
			v.Status = "READ"
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
