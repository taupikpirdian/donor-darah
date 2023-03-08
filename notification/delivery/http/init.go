package http

import (
	"net/http"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

// ResponseError represent the response error struct
type ResponseError struct {
	Message string `json:"message"`
}

// UserHandler  represent the httphandler for user
type NotificationHandler struct {
	AUsecase domain.NotificationUsecase
}

func isRequestValid(m *domain.NotificationData) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

// NewUserHandler will initialize the users/ resources endpoint
func NewNotificationHandler(e *echo.Echo, us domain.NotificationUsecase) {
	handler := &NotificationHandler{
		AUsecase: us,
	}
	e.GET("/api/v1/notification/list", handler.GetNotificationList)
	e.GET("/api/v1/notification/detail/:id", handler.GetNotificationDetail)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
