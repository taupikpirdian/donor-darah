package http

import (
	"net/http"

	"github.com/bxcodec/go-clean-arch/cfg"
	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/donor/delivery/http_response"
	"github.com/bxcodec/go-clean-arch/helper"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func (d *DonorHandler) Card(c echo.Context) (err error) {
	loggerFile := cfg.NewLoger(d.cfg.PATH_LOGS)
	contentLog := cfg.ContentLogger{
		Url: "/api/v1/donor/card",
	}

	// data user by token
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*domain.JwtCustomClaims)
	userId := claims.Id

	ctx := c.Request().Context()

	data, errUc := d.AUsecase.Card(ctx, userId)
	if errUc != nil {
		loggerFile.ErrorLogger.Println(errUc)
		responseError3, _ := http_response.MapResponseGeneral(1, domain.ErrBadBody.Error(), nil)
		return c.JSON(getStatusCode(err), responseError3)
	}

	contentLog.Payload = helper.StructToString(userId)
	contentLog.Response = helper.StructToString(data)
	loggerFile.InfoLogger.Println(contentLog)

	responseSuccess, _ := http_response.MapResponseGeneral(0, "success", data)
	return c.JSON(http.StatusOK, responseSuccess)
}
