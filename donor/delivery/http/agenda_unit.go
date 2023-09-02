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

func (d *DonorHandler) ListRegisterUserByUnit(c echo.Context) (err error) {
	loggerFile := cfg.NewLoger(d.cfg.PATH_LOGS)
	unitId, _ := strconv.Atoi(c.Param("unitId"))
	date := c.Request().URL.Query().Get("date")
	contentLog := cfg.ContentLogger{
		Url: "/api/v1/admin/donor/register/unit/",
	}

	ctx := c.Request().Context()
	datas, errUc := d.AUsecase.ListRegisterUserByUnit(ctx, unitId, date)
	if errUc != nil {
		loggerFile.ErrorLogger.Println(errUc)
		responseError3, _ := http_response.MapResponseListAgenda(1, domain.ErrBadBody.Error(), nil)
		return c.JSON(getStatusCode(err), responseError3)
	}

	contentLog.Response = helper.StructToString(datas)
	loggerFile.InfoLogger.Println(contentLog)

	responseSuccess, _ := http_response.MapResponseDonorRegisterByUnit(0, "success", datas)
	return c.JSON(http.StatusOK, responseSuccess)
}
