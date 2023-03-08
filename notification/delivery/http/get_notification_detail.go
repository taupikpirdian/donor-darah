package http

import (
	"net/http"
	"strconv"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/notification/delivery/http_response"
	"github.com/labstack/echo"
)

func (a *NotificationHandler) GetNotificationDetail(c echo.Context) (err error) {
	idP, err := strconv.Atoi(c.Param("id"))
	var notif domain.NotificationData
	err = c.Bind(&notif)
	if err != nil {
		responseError, _ := http_response.MapResponseNotificationSingle(1, err.Error(), nil)
		return c.JSON(http.StatusUnprocessableEntity, responseError)
	}

	var ok bool
	if ok, err = isRequestValid(&notif); !ok {
		responseError2, _ := http_response.MapResponseNotificationSingle(1, err.Error(), nil)
		return c.JSON(http.StatusBadRequest, responseError2)
	}

	ctx := c.Request().Context()
	id := int64(idP)

	data, errUc := a.AUsecase.GetSingleNotification(ctx, id)
	if errUc != nil {
		responseError3, _ := http_response.MapResponseNotificationSingle(1, domain.ErrBadBody.Error(), nil)
		return c.JSON(getStatusCode(err), responseError3)
	}

	responseSuccess, _ := http_response.MapResponseNotificationSingle(0, "success", data)
	return c.JSON(http.StatusOK, responseSuccess)
}
