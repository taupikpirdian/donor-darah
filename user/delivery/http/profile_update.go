package http

import (
	"net/http"

	"github.com/bxcodec/go-clean-arch/cfg"
	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/helper"
	"github.com/bxcodec/go-clean-arch/user/delivery/http_request"
	"github.com/bxcodec/go-clean-arch/user/delivery/http_response"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func (a *UserHandler) ProfileUpdateController(c echo.Context) (err error) {
	loggerFile := cfg.NewLoger(a.cfg.PATH_LOGS)
	contentLog := cfg.ContentLogger{
		Url: "/api/v1/profile/update",
	}

	ctx := c.Request().Context()

	request, err := http_request.OrderFilterRequest(c)
	// data user by token
	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(*domain.JwtCustomClaims)
	userId := claims.Id

	errUsecase := a.AUsecase.UpdateProfile(ctx, userId, request)
	if errUsecase != nil {
		loggerFile.ErrorLogger.Println(errUsecase)
		responseErrorUsecase, _ := http_response.MapResponse(1, errUsecase.Error())
		return c.JSON(getStatusCode(errUsecase), responseErrorUsecase)
	}

	contentLog.Payload = helper.StructToString(userId)
	contentLog.Response = helper.StructToString("")
	loggerFile.InfoLogger.Println(contentLog)

	responseSuccess, _ := http_response.MapResponse(0, "success")
	return c.JSON(http.StatusOK, responseSuccess)
}
