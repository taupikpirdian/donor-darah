package http

import (
	"fmt"
	"net/http"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/notification/delivery/http_response"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

// Register will store the user by given request body
func (a *NotificationHandler) GetNotificationList(c echo.Context) (err error) {
	var notif domain.NotificationData
	err = c.Bind(&notif)
	if err != nil {
		responseError, _ := http_response.MapResponseNotificationList(1, err.Error(), nil)
		return c.JSON(http.StatusUnprocessableEntity, responseError)
	}

	var ok bool
	if ok, err = isRequestValid(&notif); !ok {
		responseError2, _ := http_response.MapResponseNotificationList(1, err.Error(), nil)
		return c.JSON(http.StatusBadRequest, responseError2)
	}

	// data user by token
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*domain.JwtCustomClaims)
	fmt.Println(claims)
	// id := claims.Id

	ctx := c.Request().Context()
	data, errUc := a.AUsecase.GetListNotification(ctx)
	if errUc != nil {
		responseError3, _ := http_response.MapResponseNotificationList(1, domain.ErrBadBody.Error(), nil)
		return c.JSON(getStatusCode(err), responseError3)
	}

	responseSuccess, _ := http_response.MapResponseNotificationList(0, "success", data)
	return c.JSON(http.StatusOK, responseSuccess)
}
