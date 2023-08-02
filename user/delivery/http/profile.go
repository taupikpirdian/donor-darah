package http

import (
	"net/http"

	"github.com/bxcodec/go-clean-arch/cfg"
	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/helper"
	"github.com/bxcodec/go-clean-arch/user/delivery/http_response"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func (a *UserHandler) ProfileController(c echo.Context) (err error) {
	loggerFile := cfg.NewLoger(a.cfg.PATH_LOGS)
	contentLog := cfg.ContentLogger{
		Url: "/api/v1/profile",
	}

	ctx := c.Request().Context()

	// data user by token
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*domain.JwtCustomClaims)
	userId := claims.Id

	data, err := a.AUsecase.GetProfile(ctx, userId)
	if err != nil {
		loggerFile.ErrorLogger.Println(err)
		responseErrorUsecase, _ := http_response.MapResponseProfileError(1, err.Error())
		return c.JSON(getStatusCode(err), responseErrorUsecase)
	}

	if data.UrlImage != "" {
		data.UrlImage = a.cfg.DOMAIN + data.UrlImage
	}

	contentLog.Payload = helper.StructToString(userId)
	contentLog.Response = helper.StructToString(data)
	loggerFile.InfoLogger.Println(contentLog)

	responseSuccess, _ := http_response.MapResponseProfile(0, "success", data)
	return c.JSON(http.StatusOK, responseSuccess)
}
