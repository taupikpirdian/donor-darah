package http

import (
	"net/http"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/user/delivery/http_response"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func (a *UserHandler) ProfileController(c echo.Context) (err error) {
	ctx := c.Request().Context()

	// data user by token
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*domain.JwtCustomClaims)
	userId := claims.Id

	data, err := a.AUsecase.GetProfile(ctx, userId)
	if err != nil {
		logrus.Error(err)
		responseErrorUsecase, _ := http_response.MapResponseProfile(1, domain.ErrBadBody.Error(), nil)
		return c.JSON(getStatusCode(err), responseErrorUsecase)
	}

	responseSuccess, _ := http_response.MapResponseProfile(0, "success", data)
	return c.JSON(http.StatusOK, responseSuccess)
}
