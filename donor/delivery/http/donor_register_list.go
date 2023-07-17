package http

import (
	"net/http"

	"github.com/bxcodec/go-clean-arch/cfg"
	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/donor/delivery/http_response"
	"github.com/bxcodec/go-clean-arch/helper"
	"github.com/labstack/echo/v4"
)

func (d *DonorHandler) DonorRegisterList(c echo.Context) (err error) {
	loggerFile := cfg.NewLoger(d.cfg.PATH_LOGS)
	contentLog := cfg.ContentLogger{
		Url: "/api/v1/donor/register",
	}

	ctx := c.Request().Context()
	datas, errUc := d.AUsecase.ListDonorRegister(ctx)
	if errUc != nil {
		loggerFile.ErrorLogger.Println(errUc)
		responseError3, _ := http_response.MapResponseGeneral(1, domain.ErrBadBody.Error(), nil)
		return c.JSON(getStatusCode(err), responseError3)
	}

	contentLog.Payload = helper.StructToString("")
	contentLog.Response = helper.StructToString(datas)
	loggerFile.InfoLogger.Println(contentLog)

	responseSuccess, _ := http_response.MapResponseListDonorRegister(0, "success", datas)
	return c.JSON(http.StatusOK, responseSuccess)
}
