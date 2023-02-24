package http

import (
	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/labstack/echo"
)

// ResponseError represent the response error struct
type ResponseError struct {
	Message string `json:"message"`
}

// UserHandler  represent the httphandler for user
type NotificationHandler struct {
	AUsecase domain.NotificationUsecase
}

// NewUserHandler will initialize the users/ resources endpoint
func NewNotificationHandler(e *echo.Echo, us domain.NotificationUsecase) {
	handler := &NotificationHandler{
		AUsecase: us,
	}
	e.GET("/api/v1/notification/list", handler.GetNotificationList)
}
