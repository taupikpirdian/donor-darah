package http

import (
	"net/http"

	"github.com/bxcodec/go-clean-arch/cfg"
	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/helper"
	"github.com/bxcodec/go-clean-arch/user/delivery/http_response"
	"github.com/labstack/echo/v4"
)

func (a *UserHandler) JobController(c echo.Context) (err error) {
	loggerFile := cfg.NewLoger(a.cfg.PATH_LOGS)
	contentLog := cfg.ContentLogger{
		Url: "/api/v1/job",
	}

	ctx := c.Request().Context()
	data, err := a.AUsecase.GetJob(ctx)
	if err != nil {
		loggerFile.ErrorLogger.Println(err)
		responseErrorUsecase, _ := http_response.MapResponse(1, domain.ErrBadBody.Error())
		return c.JSON(getStatusCode(err), responseErrorUsecase)
	}

	contentLog.Payload = helper.StructToString("")
	contentLog.Response = helper.StructToString(data)
	loggerFile.InfoLogger.Println(contentLog)

	responseSuccess, _ := http_response.MapResponseJob(0, "success", data)
	return c.JSON(http.StatusOK, responseSuccess)
}
