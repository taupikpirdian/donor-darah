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

func isRequestValid(m *domain.User) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Register will store the user by given request body
func (a *UserHandler) Register(c echo.Context) (err error) {
	loggerFile := cfg.NewLoger(a.cfg.PATH_LOGS)
	contentLog := cfg.ContentLogger{
		Url: "/api/v1/register",
	}

	var user domain.User
	err = c.Bind(&user)

	if err != nil {
		loggerFile.ErrorLogger.Println(err)
		responseError, _ := http_response.MapResponse(1, err.Error())
		return c.JSON(http.StatusUnprocessableEntity, responseError)
	}

	var ok bool
	if ok, err = isRequestValid(&user); !ok {
		loggerFile.ErrorLogger.Println(err)
		responseErrorRequest, _ := http_response.MapResponse(1, domain.ErrBadParamInput.Error())
		return c.JSON(http.StatusBadRequest, responseErrorRequest)
	}

	ctx := c.Request().Context()
	err = a.AUsecase.Register(ctx, &user)
	if err != nil {
		loggerFile.ErrorLogger.Println(err)
		responseErrorUsecase, _ := http_response.MapResponse(1, err.Error())
		return c.JSON(getStatusCode(err), responseErrorUsecase)
	}

	contentLog.Payload = helper.StructToString(user)
	contentLog.Response = helper.StructToString("")
	loggerFile.InfoLogger.Println(contentLog)

	responseSuccess, _ := http_response.MapResponse(0, "success")
	return c.JSON(http.StatusCreated, responseSuccess)
}
