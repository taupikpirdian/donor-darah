package http_response

import "github.com/bxcodec/go-clean-arch/domain"

type CustomReponseSingleData struct {
	Status *Status
	Data   *domain.NotificationData
}

func MapResponseNotificationSingle(code int, message string, data *domain.NotificationData) (*CustomReponseSingleData, error) {
	if data.Status == "1" {
		data.Status = "READ"
	} else {
		data.Status = "NOT READ"
	}
	httpResponse := &CustomReponseSingleData{
		Status: &Status{
			Code:    code,
			Message: message,
		},
		Data: data,
	}

	return httpResponse, nil
}
