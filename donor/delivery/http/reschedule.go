package http

import (
	"net/http"
	"strconv"

	"github.com/bxcodec/go-clean-arch/cfg"
	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/donor/delivery/http_response"
	"github.com/bxcodec/go-clean-arch/helper"
	"github.com/labstack/echo/v4"
)

func (d *DonorHandler) RescheduleDonor(c echo.Context) (err error) {
	loggerFile := cfg.NewLoger(d.cfg.PATH_LOGS)
	contentLog := cfg.ContentLogger{
		Url: "/api/v1/donor/reschedule/:donorRegisterId",
	}

	var schedule domain.DonorSchedulleDTO
	err = c.Bind(&schedule)
	if err != nil {
		loggerFile.ErrorLogger.Println(err)
		responseError, _ := http_response.MapResponseBuktiDonor(1, err.Error(), nil)
		return c.JSON(http.StatusUnprocessableEntity, responseError)
	}

	idP, errAToi := strconv.Atoi(c.Param("donorRegisterId"))

	if errAToi != nil {
		loggerFile.ErrorLogger.Println(errAToi)
		responseErrorConv, _ := http_response.MapResponseBuktiDonor(1, err.Error(), nil)
		return c.JSON(getStatusCode(errAToi), responseErrorConv)
	}

	id := int64(idP)
	ctx := c.Request().Context()

	errUc := d.AUsecase.Reschedule(ctx, id, &schedule)
	if errUc != nil {
		loggerFile.ErrorLogger.Println(errUc)
		responseError3, _ := http_response.MapResponseBuktiDonor(1, errUc.Error(), nil)
		return c.JSON(getStatusCode(err), responseError3)
	}

	contentLog.Payload = helper.StructToString(schedule)
	contentLog.Response = helper.StructToString("")
	loggerFile.InfoLogger.Println(contentLog)

	responseSuccess, _ := http_response.MapResponseBuktiDonor(0, "success", nil)
	return c.JSON(http.StatusOK, responseSuccess)
}
