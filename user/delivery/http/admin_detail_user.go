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

func (a *UserHandler) AdminUserDetailController(c echo.Context) (err error) {
	loggerFile := cfg.NewLoger(a.cfg.PATH_LOGS)
	contentLog := cfg.ContentLogger{
		Url: "/api/v1/user/detail/:id",
	}

	ctx := c.Request().Context()
	id := c.Param("id")

	// data user by token
	userLogin := c.Get("user").(*jwt.Token)
	claims := userLogin.Claims.(*domain.JwtCustomClaims)
	role := claims.Role
	if role != "admin" {
		loggerFile.ErrorLogger.Println("only admin can access")
		responseError, _ := http_response.MapResponse(1, "only admin can access")
		return c.JSON(http.StatusUnauthorized, responseError)
	}

	data, errUsecase := a.AUsecase.DetailUser(ctx, id)
	if errUsecase != nil {
		loggerFile.ErrorLogger.Println(errUsecase)
		responseErrorUsecase, _ := http_response.MapResponse(1, errUsecase.Error())
		return c.JSON(getStatusCode(err), responseErrorUsecase)
	}

	if data.UrlImage != "" {
		data.UrlImage = a.cfg.DOMAIN + data.UrlImage
	}

	contentLog.Payload = helper.StructToString(id)
	contentLog.Response = helper.StructToString(data)
	loggerFile.InfoLogger.Println(contentLog)

	responseSuccess, _ := http_response.MapResponseProfile(0, "success", data)
	return c.JSON(http.StatusOK, responseSuccess)
}
