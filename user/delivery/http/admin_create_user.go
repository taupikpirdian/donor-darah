package http

import (
	"errors"
	"net/http"

	"github.com/bxcodec/go-clean-arch/cfg"
	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/helper"
	"github.com/bxcodec/go-clean-arch/user/delivery/http_response"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func (a *UserHandler) AdminUserCreateController(c echo.Context) (err error) {
	loggerFile := cfg.NewLoger(a.cfg.PATH_LOGS)
	contentLog := cfg.ContentLogger{
		Url: "/api/v1/user/create",
	}

	ctx := c.Request().Context()
	var user *domain.User

	err = c.Bind(&user)
	if err != nil {
		loggerFile.ErrorLogger.Println(err)
		responseError, _ := http_response.MapResponse(1, err.Error())
		return c.JSON(http.StatusUnprocessableEntity, responseError)
	}

	// data user by token
	userLogin := c.Get("user").(*jwt.Token)
	claims := userLogin.Claims.(*domain.JwtCustomClaims)
	role := claims.Role
	if role != "admin" {
		loggerFile.ErrorLogger.Println(errors.New("Role Not Admin"))
		responseError, _ := http_response.MapResponse(1, "only admin can access")
		return c.JSON(http.StatusUnauthorized, responseError)
	}

	errUsecase := a.AUsecase.CreatetUser(ctx, user)
	if errUsecase != nil {
		loggerFile.ErrorLogger.Println(errUsecase)
		responseErrorUsecase, _ := http_response.MapResponse(1, errUsecase.Error())
		return c.JSON(getStatusCode(err), responseErrorUsecase)
	}

	contentLog.Payload = helper.StructToString(user)
	contentLog.Response = helper.StructToString("")
	loggerFile.InfoLogger.Println(contentLog)
	responseSuccess, _ := http_response.MapResponse(0, "success")
	return c.JSON(http.StatusCreated, responseSuccess)
}
