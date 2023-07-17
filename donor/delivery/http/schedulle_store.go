package http

import (
	"net/http"

	"github.com/bxcodec/go-clean-arch/cfg"
	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/donor/delivery/http_request"
	"github.com/bxcodec/go-clean-arch/donor/delivery/http_response"
	"github.com/bxcodec/go-clean-arch/helper"
	"github.com/labstack/echo/v4"
)

func (d *DonorHandler) SchedulleStore(c echo.Context) (err error) {
	loggerFile := cfg.NewLoger(d.cfg.PATH_LOGS)
	contentLog := cfg.ContentLogger{
		Url: "/api/v1/donor/schedulle",
	}

	var req http_request.SchedulleStore
	err = c.Bind(&req)
	if err != nil {
		loggerFile.ErrorLogger.Println(err)
		responseError, _ := http_response.MapResponseGeneral(1, err.Error(), nil)
		return c.JSON(http.StatusUnprocessableEntity, responseError)
	}

	ctx := c.Request().Context()

	errUc := d.AUsecase.SchedulleStore(ctx, &req)
	if errUc != nil {
		loggerFile.ErrorLogger.Println(errUc)
		responseError3, _ := http_response.MapResponseGeneral(1, domain.ErrBadBody.Error(), nil)
		return c.JSON(getStatusCode(err), responseError3)
	}

	contentLog.Payload = helper.StructToString(&req)
	contentLog.Response = helper.StructToString("")
	loggerFile.InfoLogger.Println(contentLog)

	responseSuccess, _ := http_response.MapResponseGeneral(0, "success", nil)
	return c.JSON(http.StatusOK, responseSuccess)
}
