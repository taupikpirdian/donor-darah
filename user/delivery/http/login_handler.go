package http

import (
	"net/http"

	"github.com/bxcodec/go-clean-arch/cfg"
	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/helper"
	"github.com/bxcodec/go-clean-arch/user/delivery/http_response"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func isRequestValidLogin(m *domain.DtoRequestLogin) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (a *UserHandler) LoginController(c echo.Context) (err error) {
	loggerFile := cfg.NewLoger(a.cfg.PATH_LOGS)
	contentLog := cfg.ContentLogger{
		Url: "/api/v1/login",
	}

	var dto domain.DtoRequestLogin
	err = c.Bind(&dto)
	if err != nil {
		loggerFile.ErrorLogger.Println(err)
		responseError, _ := http_response.MapResponse(1, err.Error())
		return c.JSON(http.StatusUnprocessableEntity, responseError)
	}

	var ok bool
	if ok, err = isRequestValidLogin(&dto); !ok {
		loggerFile.ErrorLogger.Println(err)
		responseErrorRequest, _ := http_response.MapResponse(1, domain.ErrBadParamInput.Error())
		return c.JSON(http.StatusBadRequest, responseErrorRequest)
	}

	ctx := c.Request().Context()
	token, errUc := a.AUsecase.Login(ctx, &dto)
	if errUc != nil {
		loggerFile.ErrorLogger.Println(errUc)
		responseErrorUsecase, _ := http_response.MapResponse(1, domain.ErrBadBody.Error())
		return c.JSON(getStatusCode(errUc), responseErrorUsecase)
	}

	contentLog.Payload = helper.StructToString(dto)
	contentLog.Response = helper.StructToString(token.User)
	loggerFile.InfoLogger.Println(contentLog)

	responseSuccess, _ := http_response.MapResponseLogin(0, "success", token)
	return c.JSON(http.StatusOK, responseSuccess)
}
