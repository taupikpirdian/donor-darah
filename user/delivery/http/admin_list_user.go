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

func (a *UserHandler) AdminUserListController(c echo.Context) error {
	loggerFile := cfg.NewLoger(a.cfg.PATH_LOGS)
	contentLog := cfg.ContentLogger{
		Url: "/api/v1/user/list",
	}

	ctx := c.Request().Context()
	// data user by token
	userLogin := c.Get("user").(*jwt.Token)
	claims := userLogin.Claims.(*domain.JwtCustomClaims)
	role := claims.Role
	if role != "admin" {
		loggerFile.ErrorLogger.Println("only admin can access")
		responseError, _ := http_response.MapResponse(1, "only admin can access")
		return c.JSON(http.StatusUnauthorized, responseError)
	}

	datas, errUsecase := a.AUsecase.ListUser(ctx)
	if errUsecase != nil {
		loggerFile.ErrorLogger.Println(errUsecase)
		responseErrorUsecase, _ := http_response.MapResponse(1, errUsecase.Error())
		return c.JSON(getStatusCode(errUsecase), responseErrorUsecase)
	}

	contentLog.Payload = helper.StructToString("")
	contentLog.Response = helper.StructToString(datas)
	loggerFile.InfoLogger.Println(contentLog)
	responseSuccess, _ := http_response.MapResponseUserList(0, "success", datas, a.cfg.DOMAIN)
	return c.JSON(http.StatusOK, responseSuccess)
}
