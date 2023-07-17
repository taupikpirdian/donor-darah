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

func (d *DonorHandler) ListStockDonor(c echo.Context) (err error) {
	loggerFile := cfg.NewLoger(d.cfg.PATH_LOGS)
	contentLog := cfg.ContentLogger{
		Url: "/api/v1/donor/stock/:unitId",
	}

	ctx := c.Request().Context()
	idP, err := strconv.Atoi(c.Param("unitId"))
	if err != nil {
		loggerFile.ErrorLogger.Println(err)
		responseErrorConv, _ := http_response.MapResponseSingleAgenda(1, err.Error(), nil)
		return c.JSON(getStatusCode(err), responseErrorConv)
	}
	id := int64(idP)

	datas, errUc := d.AUsecase.ListStock(ctx, id)
	if errUc != nil {
		loggerFile.ErrorLogger.Println(errUc)
		responseError3, _ := http_response.MapResponseGeneral(1, domain.ErrBadBody.Error(), nil)
		return c.JSON(getStatusCode(err), responseError3)
	}

	contentLog.Payload = helper.StructToString(id)
	contentLog.Response = helper.StructToString(datas)
	loggerFile.InfoLogger.Println(contentLog)

	responseSuccess, _ := http_response.MapResponseGeneral(0, "success", datas)
	return c.JSON(http.StatusOK, responseSuccess)
}
