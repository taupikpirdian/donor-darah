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

func (a *UserHandler) TokenRefreshController(c echo.Context) (err error) {
	loggerFile := cfg.NewLoger(a.cfg.PATH_LOGS)
	contentLog := cfg.ContentLogger{
		Url: "/api/v1/token/refresh",
	}

	ctx := c.Request().Context()
	// data user by token
	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(*domain.JwtCustomClaims)
	userId := claims.Id

	token, errUsecase := a.AUsecase.RefreshToken(ctx, userId)
	if errUsecase != nil {
		loggerFile.ErrorLogger.Println(errUsecase)
		responseErrorUsecase, _ := http_response.MapResponse(1, errUsecase.Error())
		return c.JSON(getStatusCode(errUsecase), responseErrorUsecase)
	}

	contentLog.Payload = helper.StructToString(userId)
	contentLog.Response = helper.StructToString("")
	loggerFile.InfoLogger.Println(contentLog)

	responseSuccess, _ := http_response.MapResponseSuccess(0, "success", token)
	return c.JSON(http.StatusOK, responseSuccess)
}
