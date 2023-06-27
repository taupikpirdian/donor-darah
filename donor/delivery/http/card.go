package http

import (
	"net/http"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/donor/delivery/http_response"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func (d *DonorHandler) Card(c echo.Context) (err error) {
	// data user by token
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*domain.JwtCustomClaims)
	userId := claims.Id

	ctx := c.Request().Context()

	data, errUc := d.AUsecase.Card(ctx, userId)
	if errUc != nil {
		logrus.Error(errUc)
		responseError3, _ := http_response.MapResponseGeneral(1, domain.ErrBadBody.Error(), nil)
		return c.JSON(getStatusCode(err), responseError3)
	}

	responseSuccess, _ := http_response.MapResponseGeneral(0, "success", data)
	logrus.Info(responseSuccess)
	return c.JSON(http.StatusOK, responseSuccess)
}
