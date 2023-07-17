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

func (d *DonorHandler) DonorRegister(c echo.Context) (err error) {
	loggerFile := cfg.NewLoger(d.cfg.PATH_LOGS)
	contentLog := cfg.ContentLogger{
		Url: "/api/v1/donor/questionnaire",
	}

	var req domain.RequestRegisterDonor
	err = c.Bind(&req)
	if err != nil {
		loggerFile.ErrorLogger.Println(err)
		responseError, _ := http_response.MapResponseDonorRegister(1, err.Error())
		return c.JSON(http.StatusUnprocessableEntity, responseError)
	}

	var ok bool
	if ok, err = isRequestValid_Register(&req); !ok {
		loggerFile.ErrorLogger.Println(err)
		responseError2, _ := http_response.MapResponseDonorRegister(1, err.Error())
		return c.JSON(http.StatusBadRequest, responseError2)
	}

	ctx := c.Request().Context()
	// data user by token
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*domain.JwtCustomClaims)
	userId := claims.Id

	errUc := d.AUsecase.DonorRegister(ctx, userId, &req)
	if errUc != nil {
		loggerFile.ErrorLogger.Println(errUc)
		responseError3, _ := http_response.MapResponseDonorRegister(1, domain.ErrBadBody.Error())
		return c.JSON(getStatusCode(err), responseError3)
	}

	contentLog.Payload = helper.StructToString(&req)
	contentLog.Response = helper.StructToString("")
	loggerFile.InfoLogger.Println(contentLog)

	responseSuccess, _ := http_response.MapResponseDonorRegister(0, "success")
	return c.JSON(http.StatusOK, responseSuccess)
}
